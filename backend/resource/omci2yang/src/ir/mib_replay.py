#!/usr/bin/env python3
"""Replay an OMCI message log into a MIB snapshot.

A trace is a time-ordered stream of operations, but what determines the
configuration is the *state they leave behind*: which MEs exist and what their
attributes ended up as. Feeding a model the raw stream would make it re-derive
that state every time, so the stream is collapsed here first.

Three rules matter, and each corresponds to something real in the traces:

* **Only requests count.** Every request appears twice, once with `AR=1` and
  once as its `AK=1` response echoing the same TCID. Counting both double-counts
  every operation.
* **Only Create/Set/Delete change state.** Get, Test, Reboot, MIB upload and
  alarm traffic are observation, not configuration. `MIBReset` is the exception:
  it wipes the snapshot, so anything before it is dead configuration and must be
  dropped -- a log that spans two provisioning attempts otherwise merges them.
* **Table attributes accumulate.** Setting `Received frame VLAN tagging
  operation data` writes *one row*; a VLAN rule set is built by repeated Sets.
  Treating it as a scalar keeps only the last rule and silently loses the rest.

Usage:
    python3 mib_replay.py <onu.omci.json>
    python3 mib_replay.py --json <onu.omci.json> > snapshot.json
"""

from __future__ import annotations

import argparse
import json
import os
import sys
from dataclasses import dataclass, field

CONFIG_OPS = ("Create", "Set", "Delete")

# The corpus holds three different message encodings, because the analyzer's
# output format changed over time and some records were never decoded at all.
# Handling only the first silently discards about a quarter of the corpus, which
# looks like "these ONUs have no configuration" rather than like a bug.
FORMAT_DECODED = "decoded"        # {"header": ..., "contents": ...}
FORMAT_WRAPPED = "wrapped"        # {"Timestamp": ..., "Omci": "<json string>"}
FORMAT_RAW_HEX = "raw-hex"        # [timestamp, "TX"|"RX", "<hex>"]

# Attributes written one row at a time rather than replaced wholesale.
TABLE_ATTRS = {
    "Received frame VLAN tagging operation data",
    "DSCP to p-bit mapping",
    "IPv4 multicast address table",
    "IPv6 multicast address table",
    "Dynamic access control list table",
    "Static access control list table",
    "Multicast service package table",
    "Allowed preview groups table",
    "Lost packets table",
}

# Fields identifying a row of `Received frame VLAN tagging operation data`; a
# later Set with the same filter replaces that rule instead of adding one.
_EVTOCD_ROW_KEY = (
    "FilterOuterPriority", "FilterOuterVID", "FilterOuterTPID",
    "FilterInnerPriority", "FilterInnerVID", "FilterInnerTPID",
    "FilterEthertype",
)


@dataclass
class MeInstance:
    me_class: int
    me_class_name: str
    instance: int
    attrs: dict = field(default_factory=dict)
    created: bool = False          # seen in a Create, not just a Set
    set_count: int = 0

    @property
    def key(self) -> str:
        return f"{self.me_class_name}:{self.instance}"

    def to_dict(self) -> dict:
        return {
            "me_class": self.me_class,
            "me_class_name": self.me_class_name,
            "instance": self.instance,
            "instance_hex": f"0x{self.instance:x}",
            "created": self.created,
            "set_count": self.set_count,
            "attrs": self.attrs,
        }


@dataclass
class MibSnapshot:
    instances: dict[tuple[int, int], MeInstance] = field(default_factory=dict)
    mib_resets: int = 0
    applied: int = 0
    skipped_non_config: int = 0
    responses_ignored: int = 0
    deleted: int = 0
    dropped_by_reset: int = 0
    source_format: str = ""
    undecoded: int = 0
    warnings: list[str] = field(default_factory=list)

    def by_class_name(self, name: str) -> list[MeInstance]:
        return sorted((m for m in self.instances.values()
                       if m.me_class_name == name),
                      key=lambda m: m.instance)

    def get(self, name: str, instance: int) -> MeInstance | None:
        for me in self.instances.values():
            if me.me_class_name == name and me.instance == instance:
                return me
        return None

    def class_histogram(self) -> dict[str, int]:
        out: dict[str, int] = {}
        for me in self.instances.values():
            out[me.me_class_name] = out.get(me.me_class_name, 0) + 1
        return dict(sorted(out.items(), key=lambda kv: (-kv[1], kv[0])))

    def to_dict(self) -> dict:
        return {
            "stats": {
                "instances": len(self.instances),
                "applied_operations": self.applied,
                "mib_resets": self.mib_resets,
                "dropped_by_reset": self.dropped_by_reset,
                "deleted": self.deleted,
                "skipped_non_config": self.skipped_non_config,
                "responses_ignored": self.responses_ignored,
                "source_format": self.source_format,
                "undecoded": self.undecoded,
                "warnings": len(self.warnings),
            },
            "class_histogram": self.class_histogram(),
            "instances": [m.to_dict() for m in
                          sorted(self.instances.values(),
                                 key=lambda m: (m.me_class, m.instance))],
            "warnings": self.warnings[:50],
        }


def _merge_table(existing, incoming) -> list:
    """Merge table rows, replacing a row whose filter matches."""
    rows = list(existing) if isinstance(existing, list) else []
    new_rows = incoming if isinstance(incoming, list) else [incoming]

    for row in new_rows:
        if not isinstance(row, dict):
            if row not in rows:
                rows.append(row)
            continue
        key = tuple(row.get(k) for k in _EVTOCD_ROW_KEY)
        if any(v is None for v in key):
            rows.append(row)
            continue
        for i, old in enumerate(rows):
            if isinstance(old, dict) and \
                    tuple(old.get(k) for k in _EVTOCD_ROW_KEY) == key:
                rows[i] = row
                break
        else:
            rows.append(row)
    return rows


def normalize_messages(messages) -> tuple[list, str, int]:
    """Bring any of the three corpus encodings to `{header, contents}`.

    Returns the decoded messages, the detected format, and how many entries
    could not be decoded (raw hex, which would need a full OMCI decoder).
    """
    if not isinstance(messages, list) or not messages:
        return [], "", 0

    first = messages[0]

    if isinstance(first, dict) and "header" in first:
        return messages, FORMAT_DECODED, 0

    if isinstance(first, dict) and "Omci" in first:
        out, undecoded = [], 0
        for entry in messages:
            if not isinstance(entry, dict):
                undecoded += 1
                continue
            payload = entry.get("Omci")
            if isinstance(payload, dict):
                out.append(payload)
                continue
            if not isinstance(payload, str):
                undecoded += 1
                continue
            try:
                decoded = json.loads(payload)
            except json.JSONDecodeError:
                undecoded += 1
                continue
            if isinstance(decoded, dict) and "header" in decoded:
                if "Timestamp" in entry:
                    decoded.setdefault("timestamp", entry["Timestamp"])
                out.append(decoded)
            else:
                undecoded += 1
        return out, FORMAT_WRAPPED, undecoded

    if isinstance(first, list):
        # [timestamp, direction, hex] -- never decoded by the analyzer.
        # Decoding OMCI from hex is a separate job; report and move on.
        return [], FORMAT_RAW_HEX, len(messages)

    return [], "", len(messages)


def replay(messages: list, stop_at_last_reset: bool = True) -> MibSnapshot:
    """Collapse an OMCI message list into final ME state.

    `stop_at_last_reset` keeps only what follows the final MIB reset, which is
    what the ONU is actually running.
    """
    snap = MibSnapshot()
    if not isinstance(messages, list):
        snap.warnings.append("input is not a list of messages")
        return snap

    messages, fmt, undecoded = normalize_messages(messages)
    snap.source_format = fmt
    snap.undecoded = undecoded
    if undecoded:
        snap.warnings.append(
            f"{undecoded} message(s) could not be decoded ({fmt})")

    requests = []
    for msg in messages:
        if not isinstance(msg, dict):
            continue
        header = msg.get("header")
        if not isinstance(header, dict):
            continue
        if header.get("AR") != 1:
            snap.responses_ignored += 1
            continue
        requests.append(msg)

    if stop_at_last_reset:
        last_reset = -1
        for i, msg in enumerate(requests):
            if msg["header"].get("MsgTypeName") == "MIBReset":
                last_reset = i
        if last_reset >= 0:
            snap.mib_resets = sum(
                1 for m in requests
                if m["header"].get("MsgTypeName") == "MIBReset")
            snap.dropped_by_reset = last_reset + 1
            requests = requests[last_reset + 1:]

    for msg in requests:
        header = msg["header"]
        op = header.get("MsgTypeName")
        if op not in CONFIG_OPS:
            snap.skipped_non_config += 1
            continue

        me_class = header.get("MeClass")
        me_name = header.get("MeClassName") or f"class-{me_class}"
        instance = header.get("MeInst")
        if me_class is None or instance is None:
            snap.warnings.append(f"{op} without ME class/instance")
            continue
        ident = (me_class, instance)

        if op == "Delete":
            if snap.instances.pop(ident, None) is None:
                snap.warnings.append(
                    f"Delete of absent {me_name}:0x{instance:x}")
            snap.deleted += 1
            snap.applied += 1
            continue

        me = snap.instances.get(ident)
        if me is None:
            me = MeInstance(me_class=me_class, me_class_name=me_name,
                            instance=instance)
            snap.instances[ident] = me
        if op == "Create":
            # A second Create for a live instance means the log covers more
            # than one provisioning pass; the newer one wins.
            me.attrs = {}
            me.created = True
        else:
            me.set_count += 1

        for attr in (msg.get("contents") or {}).get("Attrs") or []:
            if not isinstance(attr, dict):
                continue
            name = attr.get("Name")
            if name is None:
                continue
            value = attr.get("Value")
            if name in TABLE_ATTRS:
                me.attrs[name] = _merge_table(me.attrs.get(name), value)
            else:
                me.attrs[name] = value
        snap.applied += 1

    return snap


_IDENTITY_ATTRS = {
    "vendor id": "vendor_id",
    "equipment id": "equipment_id",
    "actual equipment id": "equipment_id",
    "vendor product code": "product_code",
    "version": "version",
}

_IDENTITY_CLASSES = ("OnuG", "Onu2G", "OnuData")


def harvest_identity(messages: list) -> dict:
    """Pull the ONU's identity out of the *responses*.

    Vendor is required to pick the right YANGMAP overrides, but it never
    appears in Create/Set traffic -- the OLT learns it by issuing a Get, so it
    only exists in the `AK=1` replies that the state replay deliberately
    ignores. Harvesting it separately is what keeps the replay clean while
    still recovering the vendor.

    Serial numbers are skipped: they are redacted on export and identify
    customer hardware.
    """
    out: dict[str, object] = {}
    normalized, _fmt, _undecoded = normalize_messages(messages)
    for msg in normalized:
        header = msg.get("header") or {}
        if header.get("MeClassName") not in _IDENTITY_CLASSES:
            continue
        for attr in (msg.get("contents") or {}).get("Attrs") or []:
            if not isinstance(attr, dict):
                continue
            field_name = _IDENTITY_ATTRS.get(str(attr.get("Name", "")).lower())
            if field_name is None:
                continue
            value = attr.get("Value")
            if value in (None, ""):
                continue
            if isinstance(value, str) and value.startswith("REDACTED-"):
                continue
            out.setdefault(field_name, value)
    return out


def load_messages(path: str) -> list:
    with open(path, encoding="utf-8") as fh:
        return json.load(fh)


def main() -> int:
    ap = argparse.ArgumentParser(
        description=__doc__,
        formatter_class=argparse.RawDescriptionHelpFormatter)
    ap.add_argument("files", nargs="+")
    ap.add_argument("--json", action="store_true",
                    help="emit the snapshot as JSON")
    ap.add_argument("--keep-pre-reset", action="store_true",
                    help="keep operations before the last MIB reset")
    args = ap.parse_args()

    for path in args.files:
        snap = replay(load_messages(path),
                      stop_at_last_reset=not args.keep_pre_reset)
        if args.json:
            print(json.dumps(snap.to_dict(), indent=1))
            continue
        print(f"{os.path.basename(path)}")
        stats = snap.to_dict()["stats"]
        print(f"  instances={stats['instances']} "
              f"applied={stats['applied_operations']} "
              f"resets={stats['mib_resets']} "
              f"dropped_pre_reset={stats['dropped_by_reset']} "
              f"deleted={stats['deleted']}")
        for name, count in snap.class_histogram().items():
            print(f"    {count:3}  {name}")
    return 0


if __name__ == "__main__":
    sys.exit(main())
