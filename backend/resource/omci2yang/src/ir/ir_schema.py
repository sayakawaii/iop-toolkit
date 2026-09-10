#!/usr/bin/env python3
"""The business IR: the single interface between OMCI and Nokia configuration.

A MIB snapshot is still OMCI-shaped -- numbered ME classes wired together by
integer pointers. The IR resolves those pointers into named objects and, most
importantly, reassembles the end-to-end **service flows**, because that is the
unit Nokia configuration is written in. A YANG config does not say "create ME
266 instance 0x11a"; it describes a subscriber service carrying certain VLANs
over a GEM port on a T-CONT.

Everything both the compiler and any model sees goes through this structure, so
two properties are deliberate:

* **Pointers become references.** `GEMInterworkingTerminationPoint`'s
  connectivity pointer is followed to the GEM CTP, whose T-CONT pointer is
  followed to the alloc-id. The IR stores the resolved values, so downstream
  code never re-walks ME pointers.
* **Nothing is silently dropped.** MEs that are not part of the modelled service
  chain are kept in `other_mes`. A dropped ME would look like a config the
  compiler simply cannot express, which is exactly the signal needed to spot
  gaps.

Vendor is carried explicitly. `vonu-mgmt` has per-vendor YANGMAP overrides (491
for ALCL alone), so the same YANG config lands as different MEs depending on the
ONU. An IR without a vendor cannot be compiled correctly.

Usage:
    python3 ir_schema.py <onu.omci.json>
    python3 ir_schema.py --json <onu.omci.json> > ir.json
"""

from __future__ import annotations

import argparse
import json
import os
import re
import sys
from dataclasses import asdict, dataclass, field

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
from mib_replay import (  # noqa: E402
    MibSnapshot,
    harvest_identity,
    load_messages,
    replay,
)

SCHEMA_VERSION = "1.0"

# The enum tables below are transcribed from G.988 as reproduced in
# iop-toolkit's `backend/docs/g988/me/`, and cross-checked against the
# `AssociationType` class lists the analyzer itself uses in
# `omciDiagram/def/*.go`.  They are worth stating carefully: an off-by-one here
# silently mislabels a whole service chain rather than failing.

# MAC bridge port TP type -> what its TP pointer refers to.
# Cross-check: meMBPCD.go maps the code to ME class
# [-, 11, 14, 130, 134, 266, 281, 98, 117, 286, -, 329, 162].
TP_TYPE = {
    1: "pptp-ethernet-uni",             # class 11
    2: "interworking-vcc-tp",           # class 14
    3: "pbit-mapper",                   # class 130
    4: "ip-host-config-data",           # class 134
    5: "gem-interworking-tp",           # class 266
    6: "multicast-gem-interworking-tp",  # class 281
    7: "pptp-xdsl-uni",                 # class 98
    8: "pptp-vdsl-uni",                 # class 117
    9: "eth-flow-tp",                   # class 286
    11: "veip",                         # class 329
    12: "pptp-moca-uni",                # class 162
}

# G.988 9.3.13 association type -> what the EVTOCD is attached to.
# Cross-check: meEVTOCD.go maps the code to ME class
# [47, 138, 11, 134, 98, 266, 281, 162, -, 286, 329, 333].
EVTOCD_ASSOCIATION = {
    0: "mac-bridge-port",               # class 47
    1: "pbit-mapper",                   # class 130
    2: "pptp-ethernet-uni",             # class 11
    3: "ip-host-config-data",           # class 134
    4: "pptp-xdsl-uni",                 # class 98
    5: "gem-interworking-tp",           # class 266
    6: "multicast-gem-interworking-tp",  # class 281
    7: "pptp-moca-uni",                 # class 162
    8: "reserved",
    9: "eth-flow-tp",                   # class 286
    10: "veip",                         # class 329
    11: "mpls-pw-tp",                   # class 333
    12: "efm-bonding-group",
}

# G.988 Table 9.3.11-1: VlanTagFilterData forward operation.
VTFD_FORWARD_OP = {
    0x00: "forward-all",
    0x02: "forward-matching-tci-discard-others",
    0x04: "forward-untagged-discard-tagged",
    0x06: "forward-matching-priority-discard-others",
    0x08: "discard-matching-vid-forward-others",
    0x0A: "discard-matching-tci-forward-others",
    0x0C: "discard-untagged-forward-tagged",
    0x0E: "discard-all",
    0x10: "forward-matching-vid-discard-others",
    0x18: "forward-matching-vid-and-untagged",
    0x1A: "forward-matching-tci-and-untagged",
}

GEM_DIRECTION = {1: "upstream", 2: "downstream", 3: "bidirectional"}

# G.988 9.3.3 interworking option.  Note that 5 is the p-bit mapper and 6 is
# downstream broadcast -- the corpus is overwhelmingly 5, which is consistent
# with p-bit mapper based provisioning.
INTERWORKING_OPTION = {
    0: "circuit-emulated-tdm",
    1: "mac-bridged-lan",
    2: "reserved",
    3: "reserved",
    4: "video-return-path",
    5: "pbit-mapper",
    6: "downstream-broadcast",
    7: "mpls-pw-tdm",
}

# A "null" pointer / unset VID in OMCI.
NULL_POINTER = 0xFFFF
NULL_VID = 4096
NULL_PRIORITY = 15

_OUI_RE = re.compile(r"\b([A-Z]{4})[0-9A-Fa-f]{8}\b")

# The analyzer's ME class names, taken from the corpus rather than from G.988
# prose: the Ethernet UNI is reported as `PPTPEthUNI`, not
# `PhysicalPathTerminationPointEthernetUNI`, and using the long name loses
# every Ethernet UNI in the corpus without any error.
UNI_CLASSES = {
    "PPTPEthUNI": "pptp-ethernet-uni",
    "VirtualEthernetInterfacePoint": "veip",
    "IPHostConfigData": "ip-host-config-data",
    "PhysicalPathTerminationPointPOTSUNI": "pptp-pots-uni",
}

MODELLED_CLASSES = {
    "T-CONT", "GEMPortNwCTP", "GEMInterworkingTerminationPoint",
    "IEEE8021pMapperServiceProfile", "MACBridgeServiceProfile",
    "MACBridgePortConfigurationData", "VlanTagFilterData",
    "ExtVlanTagOperationConfigData", "PriorityQueue", "TCPUDPConfigData",
} | set(UNI_CLASSES)


# ----------------------------------------------------------------------
# IR objects
# ----------------------------------------------------------------------


@dataclass
class OnuIdentity:
    onu_name: str = ""
    vendor: str = ""
    vendor_oui: str = ""
    onu_type: str = ""
    pon_type: str = ""
    rg_mode: str = ""
    ls_release: str = ""
    hw_version: str = ""
    sw_version: str = ""
    customer: str = ""


@dataclass
class Tcont:
    instance: int
    alloc_id: int | None = None
    policy: int | None = None


@dataclass
class GemPort:
    instance: int
    port_id: int | None = None
    direction: str = ""
    tcont_instance: int | None = None
    tcont_alloc_id: int | None = None
    downstream_priority_queue: int | None = None
    upstream_traffic_management: int | None = None
    upstream_traffic_descriptor: int | None = None
    downstream_traffic_descriptor: int | None = None
    encryption_key_ring: int | None = None
    # from the interworking TP that points at this CTP
    iw_tp_instance: int | None = None
    interworking_option: str = ""
    service_profile_pointer: int | None = None
    service_profile_kind: str = ""


@dataclass
class PMapperEntry:
    pbit: int
    gem_iw_tp_instance: int
    gem_port_id: int | None = None


@dataclass
class PMapper:
    instance: int
    entries: list[PMapperEntry] = field(default_factory=list)
    unmarked_frame_option: int | None = None
    default_pbit_assumption: int | None = None
    tp_pointer: int | None = None


@dataclass
class Bridge:
    instance: int
    learning: bool | None = None
    port_bridging: bool | None = None


@dataclass
class BridgePort:
    instance: int
    bridge_instance: int | None = None
    port_num: int | None = None
    tp_type_code: int | None = None
    tp_type: str = ""
    tp_pointer: int | None = None
    priority: int | None = None


@dataclass
class VlanFilter:
    instance: int
    vids: list[int] = field(default_factory=list)
    priorities: list[int] = field(default_factory=list)
    forward_operation_code: int | None = None
    forward_operation: str = ""
    attached_to: str = ""


@dataclass
class VlanRule:
    """One row of the EVTOCD table, in readable form."""

    filter_outer_priority: int | None = None
    filter_outer_vid: int | None = None
    filter_inner_priority: int | None = None
    filter_inner_vid: int | None = None
    filter_ethertype: int | None = None
    tags_to_remove: int | None = None
    treatment_outer_priority: int | None = None
    treatment_outer_vid: int | None = None
    treatment_inner_priority: int | None = None
    treatment_inner_vid: int | None = None
    description: str = ""


@dataclass
class VlanOperation:
    instance: int
    association_code: int | None = None
    association: str = ""
    associated_pointer: int | None = None
    input_tpid: int | None = None
    output_tpid: int | None = None
    downstream_mode: int | None = None
    rules: list[VlanRule] = field(default_factory=list)


@dataclass
class PriorityQueueIR:
    """An upstream or downstream queue.

    Every GEM port points at one for the downstream direction, and the upstream
    ones carry the scheduling weight, so this is what a QoS policy is derived
    from.
    """

    instance: int
    related_port: int | None = None
    weight: int | None = None
    traffic_scheduler: int | None = None
    direction: str = ""              # upstream queues have the high bit set


@dataclass
class IpHost:
    """The ONU's own IP interface (G.988 9.4.2, ME 134).

    `ietf-interface-iphost.json` shows this becomes an `ipForward` interface
    with `ietf-ip` / `nokia-ip-aug` content, and that every field is written
    from a decomposed `IP options` bitmask -- so all of it inverts cleanly.
    """

    instance: int
    ip_options: int | None = None
    dhcp: bool | None = None             # options bit 0
    respond_to_pings: bool | None = None  # options bit 1
    respond_to_traceroute: bool | None = None  # options bit 2
    ip_stack_enabled: bool | None = None  # options bit 3
    ip_address: str = ""
    netmask: str = ""
    gateway: str = ""
    primary_dns: str = ""
    secondary_dns: str = ""
    onu_identifier: str = ""
    dscp_mark: int | None = None         # from the paired TCPUDPConfigData
    from_defaults: bool = False          # referenced, but never Set by the OLT


@dataclass
class Uni:
    instance: int
    kind: str                            # pptp-ethernet-uni | veip
    admin_state: int | None = None
    interdomain_name: str = ""


@dataclass
class ServiceFlow:
    """One end-to-end subscriber service, reassembled from the ME graph.

    `reaches_ani` distinguishes a real subscriber service from a UNI that is
    bridged only locally. Many HGUs put each Ethernet UNI on its own MAC bridge
    with no ANI-side port at all, because the residential gateway routes that
    traffic itself. Such a port has no GEM port and no T-CONT by design, and
    reporting it as a service with zero GEMs makes a correct parse look like a
    failure -- it accounted for the largest apparent "shape" in the corpus.
    """

    uni_kind: str = ""
    uni_instance: int | None = None
    bridge_instance: int | None = None
    bridge_port_instance: int | None = None
    pmapper_instance: int | None = None
    gem_port_ids: list[int] = field(default_factory=list)
    tcont_alloc_ids: list[int] = field(default_factory=list)
    vlans: list[int] = field(default_factory=list)
    vlan_operation_instance: int | None = None
    vlan_rule_count: int = 0
    pbits: list[int] = field(default_factory=list)
    reaches_ani: bool = False
    description: str = ""


@dataclass
class IR:
    schema_version: str = SCHEMA_VERSION
    onu: OnuIdentity = field(default_factory=OnuIdentity)
    tconts: list[Tcont] = field(default_factory=list)
    gem_ports: list[GemPort] = field(default_factory=list)
    pmappers: list[PMapper] = field(default_factory=list)
    bridges: list[Bridge] = field(default_factory=list)
    bridge_ports: list[BridgePort] = field(default_factory=list)
    vlan_filters: list[VlanFilter] = field(default_factory=list)
    vlan_operations: list[VlanOperation] = field(default_factory=list)
    unis: list[Uni] = field(default_factory=list)
    ip_hosts: list[IpHost] = field(default_factory=list)
    priority_queues: list[PriorityQueueIR] = field(default_factory=list)
    service_flows: list[ServiceFlow] = field(default_factory=list)
    other_mes: list[dict] = field(default_factory=list)
    source: dict = field(default_factory=dict)

    def to_dict(self) -> dict:
        return asdict(self)

    def summary(self) -> str:
        return (f"{len(self.service_flows)} flows, "
                f"{len(self.tconts)} tconts, {len(self.gem_ports)} gems, "
                f"{len(self.vlan_operations)} vlan-ops, "
                f"{len(self.unis)} unis, "
                f"{len(self.other_mes)} unmodelled MEs")


# ----------------------------------------------------------------------
# Derivation
# ----------------------------------------------------------------------


def _int(value) -> int | None:
    if isinstance(value, bool):
        return int(value)
    if isinstance(value, int):
        return value
    if isinstance(value, str):
        try:
            return int(value, 0)
        except ValueError:
            return None
    return None


def _pointer(value) -> int | None:
    out = _int(value)
    if out is None or out == NULL_POINTER:
        return None
    return out


def _ipv4(value) -> str:
    """Render an OMCI IPv4 attribute as a dotted quad.

    The corpus carries these as 32-bit integers. Zero is returned as unset
    rather than `0.0.0.0`, because the forward map writes 0 precisely when the
    address is *not* configured statically (`origin == 'static' ? addr : 0x00`).
    """
    if isinstance(value, str) and "." in value:
        return value.strip()
    raw = _int(value)
    if raw is None or raw == 0:
        return ""
    return ".".join(str((raw >> shift) & 0xFF) for shift in (24, 16, 8, 0))


def describe_rule(rule: VlanRule) -> str:
    """Say what a VLAN rule does, for review by an engineer.

    Two G.988 conventions decide the reading, and both are easy to get wrong:
    a *filter* priority of 15 means "match anything" (with VID 4096 as the
    don't-care VID), while a *treatment* priority of 15 means "add no tag".
    Reading the treatment VID without checking the priority reports a spurious
    "push VLAN 0" on every rule that actually pushes nothing.
    """
    matched = []
    if rule.filter_outer_priority != NULL_PRIORITY:
        vid = ("any" if rule.filter_outer_vid == NULL_VID
               else rule.filter_outer_vid)
        matched.append(f"outer vid {vid} pri {rule.filter_outer_priority}")
    elif rule.filter_outer_vid not in (None, NULL_VID):
        matched.append(f"outer vid {rule.filter_outer_vid}")
    if rule.filter_inner_priority != NULL_PRIORITY:
        vid = ("any" if rule.filter_inner_vid == NULL_VID
               else rule.filter_inner_vid)
        matched.append(f"inner vid {vid} pri {rule.filter_inner_priority}")
    elif rule.filter_inner_vid not in (None, NULL_VID):
        matched.append(f"inner vid {rule.filter_inner_vid}")
    match = " + ".join(matched) if matched else "any frame"

    actions = []
    if rule.tags_to_remove:
        actions.append("pop all tags" if rule.tags_to_remove == 3
                       else f"pop {rule.tags_to_remove}")
    if rule.treatment_outer_priority != NULL_PRIORITY:
        actions.append(f"push outer vid {rule.treatment_outer_vid} "
                       f"pri {rule.treatment_outer_priority}")
    if rule.treatment_inner_priority != NULL_PRIORITY:
        actions.append(f"push inner vid {rule.treatment_inner_vid} "
                       f"pri {rule.treatment_inner_priority}")
    action = ", ".join(actions) if actions else "no tag change"
    return f"match {match} -> {action}"


def rule_vlans(rule: VlanRule) -> list[int]:
    """VLAN IDs a rule genuinely references.

    Only tags actually added count (treatment priority != 15), plus filters on
    a concrete VID.  This is what feeds the service flow's VLAN list, so
    including don't-care values would attribute VLAN 4096 to every flow.
    """
    out = []
    if rule.treatment_outer_priority != NULL_PRIORITY and \
            rule.treatment_outer_vid not in (None, NULL_VID):
        out.append(rule.treatment_outer_vid)
    if rule.treatment_inner_priority != NULL_PRIORITY and \
            rule.treatment_inner_vid not in (None, NULL_VID):
        out.append(rule.treatment_inner_vid)
    for vid in (rule.filter_outer_vid, rule.filter_inner_vid):
        if vid not in (None, 0, NULL_VID):
            out.append(vid)
    return out


def _vlan_rules(raw) -> list[VlanRule]:
    rules = []
    if not isinstance(raw, list):
        return rules
    for row in raw:
        if not isinstance(row, dict):
            continue
        rule = VlanRule(
            filter_outer_priority=_int(row.get("FilterOuterPriority")),
            filter_outer_vid=_int(row.get("FilterOuterVID")),
            filter_inner_priority=_int(row.get("FilterInnerPriority")),
            filter_inner_vid=_int(row.get("FilterInnerVID")),
            filter_ethertype=_int(row.get("FilterEthertype")),
            tags_to_remove=_int(row.get("TreatmentTagsToRemove")),
            treatment_outer_priority=_int(row.get("TreatmentOuterPriority")),
            treatment_outer_vid=_int(row.get("TreatmentOuterVID")),
            treatment_inner_priority=_int(row.get("TreatmentInnerPriority")),
            treatment_inner_vid=_int(row.get("TreatmentInnerVID")),
        )
        rule.description = describe_rule(rule)
        rules.append(rule)
    return rules


def _vlan_filter_list(raw, number_of_entries: int | None = None
                      ) -> tuple[list[int], list[int]]:
    """Decode a VlanTagFilterData filter list into VIDs and priorities.

    The attribute is a fixed 12-slot array of 16-bit TCI values, of which only
    the first `Number of entries` are meaningful; the remaining slots hold
    leftover data.  Honouring that count is what keeps stale slots from being
    reported as real VLANs.
    """
    vids: list[int] = []
    prios: list[int] = []

    items = raw if isinstance(raw, list) else [raw]
    if number_of_entries is not None and number_of_entries >= 0:
        items = items[:number_of_entries]

    for item in items:
        if isinstance(item, dict):
            vid = _int(item.get("vlanID") or item.get("VlanID")
                       or item.get("VID"))
            pri = _int(item.get("Priority") or item.get("priority"))
            if vid:
                vids.append(vid)
                prios.append(pri if pri is not None else 0)
            continue
        tci = _int(item)
        if tci:
            vid = tci & 0x0FFF
            if vid:
                vids.append(vid)
                prios.append((tci >> 13) & 0x7)
    return vids, prios


def guess_vendor_oui(onu_name: str, snapshot: MibSnapshot,
                     identity: dict | None = None) -> str:
    """Recover the vendor OUI, in order of reliability.

    The vendor decides which YANGMAP overrides apply, so it is worth chasing
    through several sources: the ONU's own reported `Vendor id` first, then the
    OUI embedded in corpus ONU names (`..._ALCLB457B980_...`), then anything
    vendor-ish left in the snapshot.
    """
    if identity:
        vendor = identity.get("vendor_id")
        if isinstance(vendor, str) and vendor.strip():
            return vendor.strip()[:4].upper()
    m = _OUI_RE.search(onu_name or "")
    if m:
        return m.group(1)
    for me in snapshot.instances.values():
        if me.me_class_name not in ("OnuG", "Onu2G", "OnuData"):
            continue
        for key, value in me.attrs.items():
            if "vendor" in key.lower() and isinstance(value, str) and value:
                return value.strip()[:4].upper()
    return ""


def derive_ir(snapshot: MibSnapshot, identity: OnuIdentity | None = None,
              source: dict | None = None,
              reported: dict | None = None) -> IR:
    ir = IR(onu=identity or OnuIdentity(), source=source or {})
    if not ir.onu.vendor_oui:
        ir.onu.vendor_oui = guess_vendor_oui(ir.onu.onu_name, snapshot, reported)
    if reported:
        if not ir.onu.onu_type and reported.get("equipment_id"):
            ir.onu.onu_type = str(reported["equipment_id"]).strip()
        ir.source.setdefault("reported_identity", reported)

    # --- T-CONT -------------------------------------------------------
    tcont_by_instance: dict[int, Tcont] = {}
    for me in snapshot.by_class_name("T-CONT"):
        tcont = Tcont(instance=me.instance,
                      alloc_id=_int(me.attrs.get("Alloc ID")),
                      policy=_int(me.attrs.get("Policy")))
        tcont_by_instance[me.instance] = tcont
        ir.tconts.append(tcont)

    # --- GEM CTP ------------------------------------------------------
    gem_by_instance: dict[int, GemPort] = {}
    for me in snapshot.by_class_name("GEMPortNwCTP"):
        tcont_ptr = _pointer(me.attrs.get("T-CONT pointer"))
        gem = GemPort(
            instance=me.instance,
            port_id=_int(me.attrs.get("Port ID")),
            direction=GEM_DIRECTION.get(_int(me.attrs.get("Direction")), ""),
            tcont_instance=tcont_ptr,
            tcont_alloc_id=(tcont_by_instance[tcont_ptr].alloc_id
                            if tcont_ptr in tcont_by_instance else None),
            downstream_priority_queue=_pointer(
                me.attrs.get("Priority queue pointer for downstream")),
            upstream_traffic_management=_pointer(
                me.attrs.get("Traffic management pointer for upstream")),
            upstream_traffic_descriptor=_pointer(
                me.attrs.get("Traffic descriptor profile pointer for upstream")),
            downstream_traffic_descriptor=_pointer(
                me.attrs.get("Traffic descriptor profile pointer for downstream")),
            encryption_key_ring=_int(me.attrs.get("Encryption key ring")),
        )
        gem_by_instance[me.instance] = gem
        ir.gem_ports.append(gem)

    # --- GEM interworking TP: attaches a GEM CTP to a service ---------
    iw_to_gem: dict[int, GemPort] = {}
    for me in snapshot.by_class_name("GEMInterworkingTerminationPoint"):
        ctp_ptr = _pointer(
            me.attrs.get("GEM port network CTP connectivity pointer"))
        gem = gem_by_instance.get(ctp_ptr) if ctp_ptr is not None else None
        if gem is None:
            continue
        gem.iw_tp_instance = me.instance
        gem.interworking_option = INTERWORKING_OPTION.get(
            _int(me.attrs.get("Interworking option")), "")
        gem.service_profile_pointer = _pointer(
            me.attrs.get("Service profile pointer"))
        iw_to_gem[me.instance] = gem

    # --- P-bit mapper -------------------------------------------------
    pmapper_by_instance: dict[int, PMapper] = {}
    for me in snapshot.by_class_name("IEEE8021pMapperServiceProfile"):
        pmapper = PMapper(
            instance=me.instance,
            unmarked_frame_option=_int(me.attrs.get("Unmarked frame option")),
            default_pbit_assumption=_int(me.attrs.get("Default P-bit assumption")),
            tp_pointer=_pointer(me.attrs.get("TP pointer")),
        )
        for pbit in range(8):
            ptr = _pointer(
                me.attrs.get(f"Interwork TP pointer for P-bit priority {pbit}"))
            if ptr is None:
                continue
            gem = iw_to_gem.get(ptr)
            pmapper.entries.append(PMapperEntry(
                pbit=pbit, gem_iw_tp_instance=ptr,
                gem_port_id=gem.port_id if gem else None))
            if gem is not None:
                gem.service_profile_kind = "pbit-mapper"
        pmapper_by_instance[me.instance] = pmapper
        ir.pmappers.append(pmapper)

    # --- Bridges and their ports --------------------------------------
    for me in snapshot.by_class_name("MACBridgeServiceProfile"):
        learning = _int(me.attrs.get("Learning ind"))
        bridging = _int(me.attrs.get("Port bridging ind"))
        ir.bridges.append(Bridge(
            instance=me.instance,
            learning=None if learning is None else bool(learning),
            port_bridging=None if bridging is None else bool(bridging)))

    port_by_instance: dict[int, BridgePort] = {}
    for me in snapshot.by_class_name("MACBridgePortConfigurationData"):
        code = _int(me.attrs.get("TP type"))
        port = BridgePort(
            instance=me.instance,
            bridge_instance=_pointer(me.attrs.get("Bridge id pointer")),
            port_num=_int(me.attrs.get("Port num")),
            tp_type_code=code,
            tp_type=TP_TYPE.get(code, f"unknown-{code}" if code else ""),
            tp_pointer=_pointer(me.attrs.get("TP pointer")),
            priority=_int(me.attrs.get("Port priority")),
        )
        port_by_instance[me.instance] = port
        ir.bridge_ports.append(port)

    # --- VLAN tag filters (instance shares the bridge port's id) ------
    for me in snapshot.by_class_name("VlanTagFilterData"):
        vids, prios = _vlan_filter_list(
            me.attrs.get("Vlan filter list"),
            _int(me.attrs.get("Number of entries")))
        code = _int(me.attrs.get("Forward operation"))
        port = port_by_instance.get(me.instance)
        ir.vlan_filters.append(VlanFilter(
            instance=me.instance,
            vids=vids,
            priorities=prios,
            forward_operation_code=code,
            forward_operation=VTFD_FORWARD_OP.get(code, ""),
            attached_to=(f"bridge-port:{port.instance}" if port else ""),
        ))

    # --- VLAN tagging operations --------------------------------------
    vlan_op_by_pointer: dict[tuple[str, int], VlanOperation] = {}
    for me in snapshot.by_class_name("ExtVlanTagOperationConfigData"):
        code = _int(me.attrs.get("Association type"))
        op = VlanOperation(
            instance=me.instance,
            association_code=code,
            association=EVTOCD_ASSOCIATION.get(code, ""),
            associated_pointer=_int(me.attrs.get("Associated me pointer")),
            input_tpid=_int(me.attrs.get("Input TPID")),
            output_tpid=_int(me.attrs.get("Output TPID")),
            downstream_mode=_int(me.attrs.get("Downstream mode")),
            rules=_vlan_rules(
                me.attrs.get("Received frame VLAN tagging operation data")),
        )
        ir.vlan_operations.append(op)
        if op.association and op.associated_pointer is not None:
            vlan_op_by_pointer[(op.association, op.associated_pointer)] = op

    # --- UNIs ---------------------------------------------------------
    for name, kind in UNI_CLASSES.items():
        for me in snapshot.by_class_name(name):
            ir.unis.append(Uni(
                instance=me.instance,
                kind=kind,
                admin_state=_int(me.attrs.get("Administrative state")),
                interdomain_name=str(me.attrs.get("Interdomain name") or ""),
            ))

    # --- IP hosts -----------------------------------------------------
    # A TCPUDPConfigData points back at the IP host it belongs to and carries
    # the DSCP mark, so index it by that pointer first.
    dscp_by_host: dict[int, int] = {}
    for me in snapshot.by_class_name("TCPUDPConfigData"):
        host = _pointer(me.attrs.get("IP host pointer"))
        mark = _int(me.attrs.get("TOS/diffserv field"))
        if host is not None and mark is not None:
            # G.988 names this "TOS/diffserv field", where the DSCP sits in the
            # top six bits -- and traces do carry full TOS bytes (184 = EF<<2).
            # But Nokia's own forward map assigns the YANG `dscp-mark` straight
            # into the attribute (`compute: dscpmarkgot`, default 24), so ONUs
            # it configured hold a bare DSCP instead. Both populations exist,
            # so the width decides: anything too wide for a 6-bit DSCP is a TOS
            # byte and gets shifted. Guessing one convention for both yields
            # values outside `inet:dscp`, which the OLT rejects.
            dscp_by_host.setdefault(host, mark >> 2 if mark > 0x3F else mark)

    for me in snapshot.by_class_name("IPHostConfigData"):
        options = _int(me.attrs.get("IP options"))
        host = IpHost(
            instance=me.instance,
            ip_options=options,
            ip_address=_ipv4(me.attrs.get("IP address")),
            netmask=_ipv4(me.attrs.get("Mask")),
            gateway=_ipv4(me.attrs.get("Gateway")),
            primary_dns=_ipv4(me.attrs.get("Primary DNS")),
            secondary_dns=_ipv4(me.attrs.get("Secondary DNS")),
            onu_identifier=str(me.attrs.get("Onu identifier") or "").strip(),
            dscp_mark=dscp_by_host.get(me.instance),
        )
        if options is not None:
            # G.988 9.4.2 IP options bitmap, the same decomposition the YANGMAP
            # builds up in its `enableDhcp | respPin | RespTrace | enableIPStack`
            # chain, read backwards.
            host.dhcp = bool(options & 0x01)
            host.respond_to_pings = bool(options & 0x02)
            host.respond_to_traceroute = bool(options & 0x04)
            host.ip_stack_enabled = bool(options & 0x08)
        ir.ip_hosts.append(host)

    # An IP host is pre-provisioned by the ONU, so a bridge port routinely
    # points at instance 0 that the OLT never Set. Record it anyway, flagged as
    # running on ONU defaults: it is referenced and therefore real, and
    # dropping it would misreport the service as uncompilable.
    known_hosts = {h.instance for h in ir.ip_hosts}
    referenced = {port.tp_pointer for port in ir.bridge_ports
                  if port.tp_type == "ip-host-config-data"
                  and port.tp_pointer is not None}
    for instance in sorted(referenced - known_hosts):
        ir.ip_hosts.append(IpHost(instance=instance,
                                  dscp_mark=dscp_by_host.get(instance),
                                  from_defaults=True))

    # --- Priority queues ----------------------------------------------
    for me in snapshot.by_class_name("PriorityQueue"):
        # G.988 9.2.10: bit 15 of the ME id distinguishes upstream (1) from
        # downstream (0) queues.
        ir.priority_queues.append(PriorityQueueIR(
            instance=me.instance,
            related_port=_int(me.attrs.get("Related port")),
            weight=_int(me.attrs.get("Weight")),
            traffic_scheduler=_pointer(me.attrs.get("Traffic scheduler pointer")),
            direction="upstream" if me.instance & 0x8000 else "downstream",
        ))

    # --- Everything not modelled, kept so gaps are visible ------------
    for me in sorted(snapshot.instances.values(),
                     key=lambda m: (m.me_class, m.instance)):
        if me.me_class_name in MODELLED_CLASSES:
            continue
        ir.other_mes.append({
            "me_class_name": me.me_class_name,
            "me_class": me.me_class,
            "instance": me.instance,
            "attrs": me.attrs,
        })

    ir.service_flows = build_service_flows(ir, vlan_op_by_pointer,
                                           pmapper_by_instance, iw_to_gem)
    return ir


def build_service_flows(ir: IR,
                        vlan_op_by_pointer: dict[tuple[str, int], VlanOperation],
                        pmappers: dict[int, PMapper],
                        iw_to_gem: dict[int, GemPort]) -> list[ServiceFlow]:
    """Reassemble subscriber services from the resolved ME graph.

    Each bridge port facing a UNI (an Ethernet UNI, a VEIP, or an IP host) is
    one service. Its GEM ports and T-CONTs are reached through the sibling ports
    of the same bridge, since a bridge is what joins the user side to the
    network side.
    """
    filters_by_instance = {f.instance: f for f in ir.vlan_filters}
    ports_by_bridge: dict[int, list[BridgePort]] = {}
    for port in ir.bridge_ports:
        if port.bridge_instance is not None:
            ports_by_bridge.setdefault(port.bridge_instance, []).append(port)

    uni_kinds = {"pptp-ethernet-uni", "veip", "ip-host-config-data"}
    flows: list[ServiceFlow] = []

    for port in ir.bridge_ports:
        if port.tp_type not in uni_kinds:
            continue
        flow = ServiceFlow(
            uni_kind=port.tp_type,
            uni_instance=port.tp_pointer,
            bridge_instance=port.bridge_instance,
            bridge_port_instance=port.instance,
        )

        gem_ids: list[int] = []
        allocs: list[int] = []
        pbits: list[int] = []
        for sibling in ports_by_bridge.get(port.bridge_instance, []):
            if sibling.instance == port.instance:
                continue
            if sibling.tp_type == "pbit-mapper":
                pmapper = pmappers.get(sibling.tp_pointer)
                if pmapper is None:
                    continue
                flow.pmapper_instance = pmapper.instance
                for entry in pmapper.entries:
                    pbits.append(entry.pbit)
                    gem = iw_to_gem.get(entry.gem_iw_tp_instance)
                    if gem is None:
                        continue
                    if gem.port_id is not None:
                        gem_ids.append(gem.port_id)
                    if gem.tcont_alloc_id is not None:
                        allocs.append(gem.tcont_alloc_id)
            elif sibling.tp_type in ("gem-interworking-tp",
                                     "multicast-gem-interworking-tp"):
                gem = iw_to_gem.get(sibling.tp_pointer)
                if gem is None:
                    continue
                if gem.port_id is not None:
                    gem_ids.append(gem.port_id)
                if gem.tcont_alloc_id is not None:
                    allocs.append(gem.tcont_alloc_id)

        flow.gem_port_ids = sorted(set(gem_ids))
        flow.tcont_alloc_ids = sorted(set(allocs))
        flow.pbits = sorted(set(pbits))

        # VLAN handling can be attached to the bridge port, to the UNI itself,
        # or to the p-bit mapper; check each in that order.
        candidates = [("mac-bridge-port", port.instance)]
        if port.tp_pointer is not None:
            candidates.append((port.tp_type, port.tp_pointer))
        if flow.pmapper_instance is not None:
            candidates.append(("pbit-mapper", flow.pmapper_instance))
        for candidate in candidates:
            op = vlan_op_by_pointer.get(candidate)
            if op is None:
                continue
            flow.vlan_operation_instance = op.instance
            flow.vlan_rule_count = len(op.rules)
            for rule in op.rules:
                flow.vlans.extend(rule_vlans(rule))
            break

        vfilter = filters_by_instance.get(port.instance)
        if vfilter:
            flow.vlans.extend(vfilter.vids)

        flow.vlans = sorted(set(flow.vlans))
        flow.reaches_ani = bool(flow.gem_port_ids or flow.tcont_alloc_ids)

        if flow.reaches_ani:
            flow.description = (
                f"{flow.uni_kind} -> bridge 0x{(flow.bridge_instance or 0):x}"
                f"{' via p-bit mapper' if flow.pmapper_instance else ''}"
                f", vlans {flow.vlans or 'none'}"
                f", gem {flow.gem_port_ids}"
                f", alloc {flow.tcont_alloc_ids}")
        else:
            flow.description = (
                f"{flow.uni_kind} on bridge 0x{(flow.bridge_instance or 0):x} "
                "with no ANI-side port: locally bridged only, no PON service")
        flows.append(flow)

    return flows


def ir_from_omci_file(path: str, identity: OnuIdentity | None = None) -> IR:
    messages = load_messages(path)
    snapshot = replay(messages)
    ident = identity or OnuIdentity(
        onu_name=os.path.basename(path).split(".omci")[0])
    return derive_ir(snapshot, ident,
                     source={"omci_json": path,
                             "messages": len(messages),
                             "me_instances": len(snapshot.instances)},
                     reported=harvest_identity(messages))


def main() -> int:
    ap = argparse.ArgumentParser(
        description=__doc__,
        formatter_class=argparse.RawDescriptionHelpFormatter)
    ap.add_argument("files", nargs="+")
    ap.add_argument("--json", action="store_true")
    args = ap.parse_args()

    for path in args.files:
        ir = ir_from_omci_file(path)
        if args.json:
            print(json.dumps(ir.to_dict(), indent=1))
            continue
        print(f"{os.path.basename(path)}: {ir.summary()}")
        print(f"  vendor OUI: {ir.onu.vendor_oui or '(unknown)'}")
        for tcont in ir.tconts:
            print(f"  T-CONT 0x{tcont.instance:x} alloc-id={tcont.alloc_id}")
        for gem in ir.gem_ports:
            print(f"  GEM 0x{gem.instance:x} port-id={gem.port_id} "
                  f"dir={gem.direction} alloc={gem.tcont_alloc_id} "
                  f"iw={gem.interworking_option}")
        for pmapper in ir.pmappers:
            mapping = ", ".join(f"p{e.pbit}->gem{e.gem_port_id}"
                                for e in pmapper.entries)
            print(f"  PMapper 0x{pmapper.instance:x}: {mapping}")
        for op in ir.vlan_operations:
            print(f"  EVTOCD 0x{op.instance:x} on {op.association} "
                  f"0x{(op.associated_pointer or 0):x}, {len(op.rules)} rules")
            for rule in op.rules:
                print(f"      {rule.description}")
        for flow in ir.service_flows:
            print(f"  FLOW {flow.description}")
    return 0


if __name__ == "__main__":
    sys.exit(main())
