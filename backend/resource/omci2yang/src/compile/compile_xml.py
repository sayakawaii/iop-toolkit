#!/usr/bin/env python3
"""Compile an IR into a Nokia LightSpan `edit-config` document.

This is the deterministic core of the project. Given an IR derived from OMCI
plus the handful of facts OMCI cannot carry, it emits the mounted-YANG
configuration that would produce that OMCI -- with no model in the loop. What a
model is for is the residue this cannot resolve, which is why the compiler
reports its own gaps rather than guessing.

The correspondence it implements, read backwards out of the OMCI:

| OMCI                                    | Nokia mounted YANG                        |
|-----------------------------------------|-------------------------------------------|
| `T-CONT.Alloc ID`                       | `xpongemtcont/tconts/tcont/alloc-id`      |
| `GEMPortNwCTP.Port ID`                  | `xpongemtcont/gemports/gemport/gemport-id`|
| p-bit mapper entry (p-bit -> GEM)       | `gemport/traffic-class` (one per p-bit)   |
| `PPTPEthUNI` instance                   | an `ethernetCsmacd` interface + component |
| `VirtualEthernetInterfacePoint`         | an `onu-v-vrefpoint` interface            |
| `IpHostConfigData`                      | an `onu-v-enet` plus an `ipForward`       |
| `EVTOCD` rule (filter/treatment)        | `inline-frame-processing` ingress/egress  |
| `VlanTagFilterData` VIDs                | the match VLAN of the sub-interface       |

Two design points are worth stating because they are what makes the output
trustworthy:

**Names are generated, not recovered.** OMCI contains no `onu-name`, no
`tcont1_...`, no profile names -- those are OLT-side identifiers. They are
derived from the operator-supplied ONU name using the same conventions as the
plugfest reference config, so output is deterministic and reviewable, but they
are *chosen*, not observed.

**Gaps are reported, not filled.** Anything the IR carries that this does not
turn into configuration lands in `CompileResult.gaps`. That list is the input
to the next stage of the project: a gap is either a template still to write or
a genuine ambiguity for a classifier, and silently dropping it would hide both.

Usage:
    python3 compile_xml.py <ir.json> --onu-name cp1_ont1
    python3 compile_xml.py <onu.omci.json> --onu-name cp1_ont1 --from-omci
    python3 compile_xml.py <ir.json> --onu-name cp1_ont1 -o out.xml
"""

from __future__ import annotations

import argparse
import json
import os
import sys
from dataclasses import dataclass, field

HERE = os.path.dirname(os.path.abspath(__file__))
sys.path.insert(0, HERE)
sys.path.insert(0, os.path.join(HERE, "..", "validate"))
sys.path.insert(0, os.path.join(HERE, "..", "ir"))

from ir_schema import NULL_PRIORITY, NULL_VID  # noqa: E402
from schema_builder import ConfigBuilder, SchemaBuildError  # noqa: E402
from schema_index import default_yang_dir, load_index  # noqa: E402
from validate_xml import XmlValidator  # noqa: E402

# The C-VLAN tag type identity, as used by the reference config.
C_VLAN = "bbf-dot1qt:c-vlan"
MAX_TRAFFIC_CLASS = 8

# Where the VEIP's virtual port sits among the chassis' children. It is kept
# clear of the physical UNI positions, which are numbered from one.
VEIP_REL_POS = 100


@dataclass
class OltContext:
    """The facts OMCI cannot carry and an operator has to supply.

    An OMCI trace shows what landed on the ONU; it says nothing about how the
    OLT refers to it. Without at least `onu_name` there is no way to emit an
    addressable configuration, so it is required rather than defaulted.
    """

    onu_name: str
    vendor: str = ""
    uni_count: int | None = None
    target: str = "running"
    queue_count: int = MAX_TRAFFIC_CLASS

    # An already-provisioned ONU has a chassis component whose name was chosen
    # when it was created and need not follow the convention below. Re-deriving
    # it produces a second chassis, which the OLT rejects outright ("Only one
    # chassis configuration is supported"), so it can be supplied instead.
    chassis_name: str = ""

    # A VEIP rides a virtual-UNI component that the ONU reports rather than one
    # the OLT configures, so its name has to be read off the box. An SFU has no
    # such component at all.
    veip_component: str = ""

    # Naming conventions, matching plugfest/rpc/request/create_onu_online.xml.
    def chassis(self) -> str:
        return self.chassis_name or f"ont_{self.onu_name}"

    def cage(self) -> str:
        return f"ontcage_{self.onu_name}"

    def transceiver(self) -> str:
        return f"ontsfp_{self.onu_name}"

    def ani_port(self) -> str:
        return f"ontaniport_{self.onu_name}"

    def ani_if(self) -> str:
        return self.onu_name

    def venet_if(self) -> str:
        return f"ontvenet_{self.onu_name}"

    def veip_if(self) -> str:
        return f"ontveip_{self.onu_name}"

    def veip_port(self) -> str:
        return f"ontveipport_{self.onu_name}"

    def uni_component(self, n: int) -> str:
        return f"ontuni_{self.onu_name}_uni{n}"

    def uni_if(self, n: int) -> str:
        return f"{self.onu_name}_uni{n}"

    def subif(self, uni: int, user: int) -> str:
        return f"enet_{self.onu_name}_uni{uni}_user{user}"

    def ip_host_if(self, n: int) -> str:
        return f"iphost_{self.onu_name}_uni{n}"

    def tcont(self, n: int) -> str:
        return f"tcont{n}_{self.onu_name}"

    def gemport(self, uni: int, user: int, tc: int) -> str:
        return f"{self.onu_name}_uni{uni}_user{user}_gemport{tc}"

    def classifier(self, pbit: int) -> str:
        return f"{self.onu_name}_classifier1_{pbit}"

    def policy(self) -> str:
        return f"{self.onu_name}_policy_pbit2tc"

    def qos_profile(self) -> str:
        return f"{self.onu_name}_QPP0"

    def tc2queue(self) -> str:
        return f"tc2queue1_{self.onu_name}"


@dataclass
class CompileResult:
    xml: str = ""
    onu_name: str = ""
    gaps: list[str] = field(default_factory=list)
    notes: list[str] = field(default_factory=list)
    services: int = 0
    validation_errors: list[str] = field(default_factory=list)
    validation_warnings: list[str] = field(default_factory=list)

    @property
    def valid(self) -> bool:
        return not self.validation_errors

    def summary(self) -> str:
        state = "valid" if self.valid else f"{len(self.validation_errors)} errors"
        return (f"{self.onu_name}: {self.services} service(s), {state}, "
                f"{len(self.gaps)} gap(s)")

    def to_dict(self) -> dict:
        """The whole outcome, for a caller that is not a terminal.

        Gaps travel with the XML rather than being logged away from it: a gap
        means the document is schema-correct but incomplete, so whoever
        receives the configuration is the one who needs to see them.
        """
        return {
            "onuName": self.onu_name,
            "xml": self.xml,
            "services": self.services,
            "valid": self.valid,
            "gaps": self.gaps,
            "notes": self.notes,
            "validationErrors": self.validation_errors,
            "validationWarnings": self.validation_warnings,
            "summary": self.summary(),
        }


def _service_rules(vlan_op: dict | None) -> list[dict]:
    """Keep the EVTOCD rows that actually define a VLAN translation.

    Real traces carry catch-all rows -- "any frame, add no tag" and
    "pop everything" defaults -- alongside the service rules. Compiling those
    would emit meaningless match/rewrite pairs, so a row only counts when it
    either matches a concrete VID or pushes a tag.
    """
    if not vlan_op:
        return []
    out = []
    for rule in vlan_op.get("rules") or []:
        pushes = (rule.get("treatment_outer_priority") != NULL_PRIORITY
                  or rule.get("treatment_inner_priority") != NULL_PRIORITY)
        matches = (rule.get("filter_outer_vid") not in (None, 0, NULL_VID)
                   or rule.get("filter_inner_vid") not in (None, 0, NULL_VID))
        if pushes or matches:
            out.append(rule)
    return out


def _bool(value) -> str:
    return "true" if value else "false"


def _traffic_classes(flow: dict, pmappers: dict) -> tuple[list, dict]:
    """Split a service into its GEM ports and the p-bits that feed each.

    An OMCI p-bit mapper is eight pointers, and several p-bits routinely point
    at the *same* GEM port. Emitting one gemport per p-bit therefore repeats a
    gemport-id, which the OLT refuses ("Gemport-id should be unique in onu").

    So the GEM port is the unit: each distinct one becomes a gemport with its
    own traffic class, and the p-bits that share it are all classified into
    that class.

    Returns `[(gem_port_id, traffic_class)]` and `{pbit: traffic_class}`.
    """
    pmapper = pmappers.get(flow.get("pmapper_instance")) or {}
    entries = pmapper.get("entries") or []

    order: list[int] = []
    pbits_of: dict[int, list[int]] = {}
    if entries:
        for entry in entries:
            gem = entry.get("gem_port_id")
            if gem is None:
                continue
            if gem not in pbits_of:
                pbits_of[gem] = []
                order.append(gem)
            pbits_of[gem].append(entry["pbit"])
    else:
        for gem in flow.get("gem_port_ids") or []:
            if gem not in pbits_of:
                pbits_of[gem] = []
                order.append(gem)

    gems = [(gem, tc) for tc, gem in enumerate(order)]
    pbit_class = {pbit: tc for gem, tc in gems for pbit in pbits_of[gem]}
    return gems, pbit_class


def _dedupe_flows(flows: list) -> list:
    """Collapse flows that describe the same service.

    Several bridge ports can front the same UNI with the same GEM ports and
    T-CONT, which is one service seen more than once. Compiling each copy emits
    the same sub-interface and gemport names again, and duplicate keys in one
    list are invalid.
    """
    seen: set[tuple] = set()
    out = []
    for flow in flows:
        signature = (
            flow.get("uni_kind"),
            flow.get("uni_instance"),
            tuple(flow.get("gem_port_ids") or []),
            tuple(flow.get("tcont_alloc_ids") or []),
            flow.get("pmapper_instance"),
            flow.get("vlan_operation_instance"),
        )
        if signature in seen:
            continue
        seen.add(signature)
        out.append(flow)
    return out


def _needs_qos_profile(flow: dict) -> bool:
    """Whether the OLT will insist on an ingress QoS policy profile here.

    Its rule covers a sub-interface layered on either an `ethernetCsmacd` UNI
    or a `vrefpoint`: both are user-facing ports where ingress traffic has to be
    classified into a traffic class. An IP host is the ONU's own interface and
    is exempt.
    """
    return flow.get("uni_kind") in ("pptp-ethernet-uni", "veip")


def _match_vid(rule: dict) -> int | None:
    for key in ("filter_inner_vid", "filter_outer_vid"):
        vid = rule.get(key)
        if vid not in (None, 0, NULL_VID):
            return vid
    return None


def _matched_tags(rule: dict) -> list[int]:
    """The VLAN tags the rule's filter actually describes, outermost first.

    A filter position with the don't-care VID (4096) matches no specific tag,
    so it is not one of them. This count is what decides how many tags the
    rewrite may pop: the OLT rejects a rewrite that pops more tags than the
    match describes.
    """
    tags = []
    for key in ("filter_outer_vid", "filter_inner_vid"):
        vid = rule.get(key)
        if vid not in (None, 0, NULL_VID):
            tags.append(vid)
    return tags


def _pop_count(rule: dict, matched: int) -> int:
    """How many tags the rewrite should pop.

    G.988 9.3.13 encodes "number of tags to remove" as 0-2 literally and 3 as
    "remove all". Reading 3 as three tags is what produced rewrites that popped
    two tags off a single-tagged match, which the OLT refuses.
    """
    pops = rule.get("tags_to_remove")
    if pops is None:
        return min(1, matched)
    if pops >= 3:
        return matched
    return min(pops, matched)


def _push_vid(rule: dict) -> int | None:
    """The VLAN the rule adds, if any."""
    if rule.get("treatment_outer_priority") != NULL_PRIORITY:
        vid = rule.get("treatment_outer_vid")
        if vid not in (None, NULL_VID):
            return vid
    if rule.get("treatment_inner_priority") != NULL_PRIORITY:
        vid = rule.get("treatment_inner_vid")
        if vid not in (None, NULL_VID):
            return vid
    return None


class Compiler:
    def __init__(self, index, validate: bool = True):
        self.index = index
        self.validator = XmlValidator(index) if validate else None

    # ------------------------------------------------------------------

    def compile(self, ir: dict, ctx: OltContext) -> CompileResult:
        result = CompileResult(onu_name=ctx.onu_name)
        builder = ConfigBuilder(self.index, target=ctx.target)

        flows = _dedupe_flows([f for f in (ir.get("service_flows") or [])
                               if f.get("reaches_ani")])
        local_only = [f for f in (ir.get("service_flows") or [])
                      if not f.get("reaches_ani")]
        if local_only:
            result.notes.append(
                f"{len(local_only)} UNI(s) are bridged locally with no PON "
                "service; no ANI configuration is emitted for them")
        if not flows:
            result.gaps.append(
                "the IR has no ANI-reaching service flow, so there is no "
                "subscriber service to configure")

        vlan_ops = {op["instance"]: op for op in (ir.get("vlan_operations") or [])}
        tconts = ir.get("tconts") or []
        gem_ports = {g["instance"]: g for g in (ir.get("gem_ports") or [])}
        pmappers = {p["instance"]: p for p in (ir.get("pmappers") or [])}
        ip_hosts = {h["instance"]: h for h in (ir.get("ip_hosts") or [])}

        # Assign a UNI ordinal per flow: OMCI instance numbers are not the
        # port numbers the OLT uses, so they are renumbered from 1 in a stable
        # order rather than passed through.
        uni_index: dict[int, int] = {}
        for flow in flows:
            key = flow.get("uni_instance")
            if key not in uni_index:
                uni_index[key] = len(uni_index) + 1

        try:
            onus = builder.top("onus")
            onu = onus.child("onu")
            onu.leaf("name", ctx.onu_name)
            root = onu.child("root")

            self._hardware(root, ctx, uni_index, flows, ir, result)
            self._interfaces(root, ctx, uni_index, flows, vlan_ops, result,
                             ip_hosts=ip_hosts)
            self._qos(root, ctx, flows, pmappers, result)
            self._xpongemtcont(root, ctx, uni_index, flows, tconts,
                               gem_ports, pmappers, result)
        except SchemaBuildError as exc:
            result.gaps.append(f"schema rejected the construction: {exc}")
            result.validation_errors.append(str(exc))
            return result

        result.services = len(flows)
        result.xml = builder.tostring()

        if self.validator is not None:
            findings = self.validator.validate_string(result.xml)
            result.validation_errors = [f"{f.path}: {f.message}"
                                        for f in findings
                                        if f.severity == "error"]
            result.validation_warnings = [f"{f.path}: {f.message}"
                                          for f in findings
                                          if f.severity == "warning"]
        return result

    # ------------------------------------------------------------------

    def _hardware(self, root, ctx: OltContext, uni_index: dict,
                  flows: list, ir: dict, result: CompileResult) -> None:
        """The ONU's physical components: chassis, cage, transceiver, UNI ports."""
        hardware = root.child("hardware")

        chassis = hardware.child("component")
        chassis.leaf("name", ctx.chassis())
        chassis.leaf("class", "ianahw:chassis")
        chassis.leaf("parent-rel-pos", 1)
        # The trace's vendor is only a safe default when this ONU is being
        # created. Reusing a chassis that the OLT already holds means real
        # hardware, and stamping the trace's vendor onto it would both misstate
        # what the ONU is and change which vendor YANGMAP overrides the OLT
        # selects -- so an existing chassis keeps whatever the box reports.
        vendor = ctx.vendor
        if not vendor and not ctx.chassis_name:
            vendor = (ir.get("onu") or {}).get("vendor_oui") or ""
        if vendor:
            chassis.leaf("mfg-name", vendor)
        else:
            result.gaps.append(
                "no vendor OUI available; hardware mfg-name is omitted and "
                "vendor-specific YANGMAP overrides cannot be selected")
        chassis.leaf("admin-state", "unlocked")

        cage = hardware.child("component")
        cage.leaves(name=ctx.cage(), **{"class": "bbf-hwt:cage"})
        cage.leaf("parent", ctx.chassis())
        cage.leaf("parent-rel-pos", 0)

        sfp = hardware.child("component")
        sfp.leaf("name", ctx.transceiver())
        sfp.leaf("class", "bbf-hwt:transceiver")
        sfp.leaf("parent", ctx.cage())
        sfp.leaf("parent-rel-pos", 0)

        ani = hardware.child("component")
        ani.leaf("name", ctx.ani_port())
        ani.leaf("class", "bbf-hwt:transceiver-link")
        ani.leaf("parent", ctx.transceiver())
        ani.leaf("parent-rel-pos", 1)

        for flow in flows:
            n = uni_index[flow.get("uni_instance")]
            if flow.get("uni_kind") != "pptp-ethernet-uni":
                continue
            comp = hardware.child("component")
            comp.leaf("name", ctx.uni_component(n))
            comp.leaf("class", "nokia-hwi:rj45-1G")
            comp.leaf("parent", ctx.chassis())
            comp.leaf("parent-rel-pos", n)
            helper = comp.child("omci-identifier-helper")
            helper.leaf("virtual-board-number", 1)

        # A VEIP has no externally visible port, so it hangs off a
        # `virtual-port` component: nokia-hardware-identities describes that
        # class as the underlying hardware port for virtual interfaces such as
        # a TR-385 onu-v-vrefpoint. Without it the v-vrefpoint has nothing to
        # reference and the ONU instantiates no VEIP.
        if any(f.get("uni_kind") == "veip" for f in flows):
            comp = hardware.child("component")
            comp.leaf("name", ctx.veip_component or ctx.veip_port())
            comp.leaf("class", "nokia-hwi:virtual-port")
            comp.leaf("parent", ctx.chassis())
            comp.leaf("parent-rel-pos", VEIP_REL_POS)

    def _interfaces(self, root, ctx: OltContext, uni_index: dict, flows: list,
                    vlan_ops: dict, result: CompileResult,
                    ip_hosts: dict | None = None) -> None:
        interfaces = root.child("interfaces")
        ip_hosts = ip_hosts or {}

        ani = interfaces.child("interface")
        ani.leaf("name", ctx.ani_if())
        ani.leaf("type", "bbf-xponift-mounted:ani")
        ani.leaf("port-layer-if", ctx.ani_port())
        ani.child("performance").leaf("enable", "false")

        # These two look alike but are different interfaces, and swapping them
        # silently produces the wrong ME on the wire. `onu-v-enet` is the ONU's
        # own IP host: `ietf-interface-onu-venet-omci` pins the bridge port's
        # TP type to 4 and points it at the IP host instance. A G.988 VEIP is
        # `onu-v-vrefpoint`, which `ietf-interface-veip` renders as TP type 11
        # against the VEIP instance. Compiling a VEIP as a v-enet therefore
        # yields an IP host, which is what the lab OLT was seen to do.
        if any(f.get("uni_kind") == "ip-host-config-data" for f in flows):
            venet = interfaces.child("interface")
            venet.leaf("name", ctx.venet_if())
            venet.leaf("type", "bbf-xponift-mounted:onu-v-enet")
            venet.child("onu-v-enet").leaf("ani", ctx.ani_if())

        if any(f.get("uni_kind") == "veip" for f in flows):
            vrp = interfaces.child("interface")
            vrp.leaf("name", ctx.veip_if())
            vrp.leaf("type", "bbf-xponift-mounted:onu-v-vrefpoint")
            vrp.child("onu-v-vrefpoint").leaf("related-onu", ctx.ani_if())
            vrp.child("performance").leaf("enable", "false")
            # `ietf-interface-veip` derives the bridge port from this port
            # reference, so without it the OLT accepts the interface and emits
            # no VEIP. It points at the `virtual-port` component emitted above,
            # or at the one the ONU already reports if it has its own.
            vrp.leaf("port-layer-if", ctx.veip_component or ctx.veip_port())

        for flow in flows:
            kind = flow.get("uni_kind")
            n = uni_index[flow.get("uni_instance")]

            if kind == "pptp-ethernet-uni":
                uni = interfaces.child("interface")
                uni.leaf("name", ctx.uni_if(n))
                uni.leaf("type", "ianaift-mounted:ethernetCsmacd")
                uni.leaf("port-layer-if", ctx.uni_component(n))
                uni.child("performance").leaf("enable", "false")
                eth = uni.child("ethernet")
                eth.child("auto-negotiation").leaf("status", "enabled")
                lower = ctx.uni_if(n)
            elif kind == "veip":
                lower = ctx.veip_if()
            elif kind == "ip-host-config-data":
                lower = ctx.venet_if()
            else:
                result.gaps.append(
                    f"UNI kind {kind!r} has no interface template; "
                    "only Ethernet UNIs, VEIPs and IP hosts are compiled")
                continue

            self._sub_interface(interfaces, ctx, n, flow, vlan_ops, lower,
                                result)

            if kind == "ip-host-config-data":
                self._ip_host(interfaces, ctx, n, flow, ip_hosts, result)

    def _ip_host(self, interfaces, ctx: OltContext, n: int, flow: dict,
                 ip_hosts: dict, result: CompileResult) -> None:
        """The ONU's own IP interface, layered on its VLAN sub-interface.

        Shape and field placement come from `ietf-interface-iphost.json`: an
        `ipForward` interface carrying `ietf-ip` addresses plus the
        `nokia-ip-aug` diagnose/acquisition augments. That map computes the
        OMCI `IP options` byte as
        `enableDhcp | respPin | RespTrace | enableIPStack`, and writes the
        address fields only when the origin is static -- both of which the IR
        has already undone, so this reads the decoded fields straight off it.
        """
        host = ip_hosts.get(flow.get("uni_instance"))
        if host is None:
            result.gaps.append(
                f"UNI {n} is an IP host but no IPHostConfigData was captured; "
                "the ipForward interface is omitted")
            return
        if host.get("from_defaults"):
            result.notes.append(
                f"IP host {n} was never Set over OMCI, so it runs on the ONU's "
                "own defaults; the interface is emitted without addressing")

        ip = interfaces.child("interface")
        ip.leaf("name", ctx.ip_host_if(n))
        ip.leaf("type", "ianaift-mounted:ipForward")
        ip.child("ipif-lower-layer").leaf("sub-interface", ctx.subif(n, 1))

        ipv4 = ip.child("ipv4")
        if host.get("ip_stack_enabled") is not None:
            ipv4.leaf("enabled", _bool(host["ip_stack_enabled"]))

        dhcp = host.get("dhcp")
        if dhcp is not None:
            ipv4.child("ip-address-acquisition-method").leaf(
                "origin", "dhcp" if dhcp else "static")

        # The map only writes an address when the origin is static, so a
        # DHCP host legitimately has none.
        address, netmask = host.get("ip_address"), host.get("netmask")
        if address:
            entry = ipv4.child("address")
            entry.leaf("ip", address)
            if netmask:
                entry.leaf("netmask", netmask)
        elif dhcp is False:
            result.gaps.append(
                f"IP host {n} is statically addressed but the trace carries "
                "no IP address")
        elif dhcp is None:
            result.notes.append(
                f"IP host {n} has no IP options byte in the trace, so neither "
                "its addressing mode nor its address is known")

        # The `dhcp` container is gated on the acquisition method ("Only
        # applicable when origin is DHCP"), so a statically addressed host must
        # not carry it even though OMCI still reports a DSCP mark for it.
        if host.get("dscp_mark") is not None and dhcp:
            ipv4.child("dhcp").leaf("dscp-mark", host["dscp_mark"])

        pings = host.get("respond_to_pings")
        trace = host.get("respond_to_traceroute")
        if pings is not None or trace is not None:
            diag = ip.child("ip-diagnose-control")
            if pings is not None:
                diag.leaf("respond-to-pings", _bool(pings))
            if trace is not None:
                diag.leaf("respond-to-traceroute-messages", _bool(trace))

        # Gateway and DNS live outside the interface (static route / DNS
        # resolver), which this compiler does not emit yet.
        if host.get("gateway"):
            result.gaps.append(
                f"IP host {n} has gateway {host['gateway']}, which belongs in "
                "an ietf-routing static route that is not compiled yet")
        if host.get("primary_dns") or host.get("secondary_dns"):
            result.gaps.append(
                f"IP host {n} has DNS servers, which belong in an "
                "ietf-system dns-resolver that is not compiled yet")

    def _sub_interface(self, interfaces, ctx: OltContext, n: int, flow: dict,
                       vlan_ops: dict, lower: str,
                       result: CompileResult) -> None:
        """The VLAN sub-interface: where the EVTOCD rules land."""
        subif = interfaces.child("interface")
        subif.leaf("name", ctx.subif(n, 1))
        subif.leaf("type", "bbfift-mounted:vlan-sub-interface")
        if flow.get("pmapper_instance") is not None or _needs_qos_profile(flow):
            subif.leaf("ingress-qos-policy-profile", ctx.qos_profile())
        subif.child("subif-lower-layer").leaf("interface", lower)

        vlan_op = vlan_ops.get(flow.get("vlan_operation_instance"))
        rules = _service_rules(vlan_op)
        if not rules:
            # No translation rule is not automatically a hole in the compiler.
            # A service that neither matches nor pushes a tag is transparent,
            # and a sub-interface without frame processing is the right
            # rendering of it. It is only a gap when the trace does show VLANs
            # for this service and none of them could be expressed.
            if flow.get("vlans"):
                result.gaps.append(
                    f"UNI {n} carries VLANs "
                    f"{sorted(set(flow['vlans']))} but no EVTOCD rule could be "
                    f"compiled (EVTOCD instance "
                    f"{flow.get('vlan_operation_instance')}); the "
                    "sub-interface is emitted without frame processing")
            else:
                result.notes.append(
                    f"UNI {n} has no VLAN operation, so the service is "
                    "transparent and the sub-interface carries no frame "
                    "processing")
            return

        processing = subif.child("inline-frame-processing")
        ingress = processing.child("ingress-rule")
        emitted = 0
        for i, rule in enumerate(rules):
            match_vid = _match_vid(rule)
            push_vid = _push_vid(rule)
            if match_vid is None and push_vid is None:
                continue
            entry = ingress.child("rule")
            entry.leaf("name", f"rule_translate_singletag{'' if i == 0 else i}")
            entry.leaf("priority", 100 + i)
            criteria = entry.child("flexible-match").child("match-criteria")

            # Describe every tag the OMCI filter names, so that the pop count
            # below stays consistent with the match.
            matched = _matched_tags(rule)
            if not matched and push_vid is not None:
                matched = [push_vid]
            for index, vid in enumerate(matched):
                tag = criteria.child("tag")
                tag.leaf("index", index)
                dot1q = tag.child("dot1q-tag")
                dot1q.leaf("tag-type", C_VLAN)
                dot1q.leaf("vlan-id", vid)
                dot1q.leaf("pbit", "any")
                dot1q.leaf("dei", "any")

            rewrite = entry.child("ingress-rewrite")
            rewrite.leaf("pop-tags", _pop_count(rule, len(matched)))
            if push_vid is not None:
                push = rewrite.child("push-tag")
                push.leaf("index", 0)
                pushed = push.child("dot1q-tag")
                pushed.leaf("tag-type", C_VLAN)
                pushed.leaf("vlan-id", push_vid)
                pushed.leaf("pbit-from-tag-index", 0)
                pushed.leaf("dei-from-tag-index", 0)
            emitted += 1

        # The egress direction mirrors the first rule: restore the customer tag.
        first_match = next((_match_vid(r) for r in rules
                            if _match_vid(r) is not None), None)
        if first_match is not None:
            egress = processing.child("egress-rewrite")
            egress.leaf("pop-tags", 1)
            push = egress.child("push-tag")
            push.leaf("index", 0)
            dot1q = push.child("dot1q-tag")
            dot1q.leaf("tag-type", C_VLAN)
            dot1q.leaf("vlan-id", first_match)
            dot1q.leaf("pbit-from-tag-index", 0)
            dot1q.leaf("dei-from-tag-index", 0)

        if emitted < len(rules):
            result.notes.append(
                f"UNI {n}: {len(rules) - emitted} EVTOCD rule(s) had neither a "
                "concrete match nor a pushed tag and were skipped")

    def _qos(self, root, ctx: OltContext, flows: list, pmappers: dict,
             result: CompileResult) -> None:
        """p-bit to traffic-class classification.

        The OMCI p-bit mapper is a table of eight pointers; on the Nokia side
        the same intent is expressed as a classifier per p-bit feeding a
        policy, referenced by the sub-interface.
        """
        used = {flow.get("pmapper_instance") for flow in flows
                if flow.get("pmapper_instance") is not None}
        # A sub-interface over a physical Ethernet UNI must carry an ingress
        # QoS policy profile whether or not OMCI used a p-bit mapper ("if
        # lower-layer interface is ethernetCsmacd or vrefpoint, then
        # ingress-qos-policy-profile must be configured"), so the policy has to
        # exist for those flows too.
        if not used and any(_needs_qos_profile(f) for f in flows):
            used = {None}
        if not used:
            # No p-bit classification to express, but the T-CONT queues still
            # need their traffic-class-to-queue profile: without it the OLT
            # rejects the T-CONT for referencing a profile that does not exist
            # ("tcont queue id should be included in tc2queue profile local
            # queue id").
            self._tm_profiles(root, ctx)
            return

        # Classify each p-bit into the traffic class of the GEM port it maps
        # to, so the classifiers and the gemports agree on what a class means.
        pbit_class: dict[int, int] = {}
        for flow in flows:
            for pbit, tc in _traffic_classes(flow, pmappers)[1].items():
                pbit_class.setdefault(pbit, tc)
        if not pbit_class:
            pbit_class = {p: p for p in range(MAX_TRAFFIC_CLASS)}

        classifiers = root.child("classifiers")
        for pbit in sorted(pbit_class):
            entry = classifiers.child("classifier-entry")
            entry.leaf("name", ctx.classifier(pbit))
            tag = entry.child("match-criteria").child("tag")
            tag.leaf("index", 0)
            tag.leaf("in-pbit-list", pbit)
            action = entry.child("classifier-action-entry-cfg")
            action.leaf("action-type", "scheduling-traffic-class")
            action.leaf("scheduling-traffic-class", pbit_class[pbit])

        policies = root.child("policies")
        policy = policies.child("policy")
        policy.leaf("name", ctx.policy())
        for pbit in sorted(pbit_class):
            policy.child("classifiers").leaf("name", ctx.classifier(pbit))

        profiles = root.child("qos-policy-profiles")
        profile = profiles.child("policy-profile")
        profile.leaf("name", ctx.qos_profile())
        profile.child("policy-list").leaf("name", ctx.policy())

        self._tm_profiles(root, ctx)

    def _tm_profiles(self, root, ctx: OltContext) -> None:
        """The traffic-class-to-queue profile every T-CONT's queues refer to."""
        tm = root.child("tm-profiles")
        mapping = tm.child("tc-id-2-queue-id-mapping-profile")
        mapping.leaf("name", ctx.tc2queue())
        for tc in range(ctx.queue_count):
            entry = mapping.child("mapping-entry")
            entry.leaf("traffic-class-id", tc)
            entry.leaf("local-queue-id", tc)

    def _xpongemtcont(self, root, ctx: OltContext, uni_index: dict, flows: list,
                      tconts: list, gem_ports: dict, pmappers: dict,
                      result: CompileResult) -> None:
        alloc_ids = [t.get("alloc_id") for t in tconts
                     if t.get("alloc_id") is not None]
        if not alloc_ids and not flows:
            return

        section = root.child("xpongemtcont")

        tcont_names: dict[int, str] = {}
        if alloc_ids:
            container = section.child("tconts")
            for i, alloc in enumerate(sorted(set(alloc_ids)), 1):
                name = ctx.tcont(i)
                tcont_names[alloc] = name
                tcont = container.child("tcont")
                tcont.leaf("name", name)
                tcont.leaf("alloc-id", alloc)
                tcont.leaf("interface-reference", ctx.ani_if())
                tm_root = tcont.child("tm-root")
                for queue_id in range(ctx.queue_count):
                    queue = tm_root.child("queue")
                    queue.leaf("local-queue-id", queue_id)
                    queue.leaf("priority", queue_id)
                    queue.leaf("weight", 1)
                tm_root.leaf("tc-id-2-queue-id-mapping-profile-name",
                             ctx.tc2queue())
        else:
            result.gaps.append(
                "no T-CONT alloc-id in the IR; upstream scheduling cannot be "
                "configured")

        gem_container = section.child("gemports")
        emitted = 0
        unreferenced = 0
        names: set[str] = set()
        # A GEM port-id has to be unique across the whole ONU, not just within
        # one service.
        gem_ids: set[int] = set()
        for flow in flows:
            n = uni_index[flow.get("uni_instance")]
            subif = ctx.subif(n, 1)

            # A gemport must name the T-CONT it drains into, so one that
            # cannot be resolved is skipped rather than emitted incomplete.
            allocs = [a for a in (flow.get("tcont_alloc_ids") or [])
                      if a in tcont_names]
            if not allocs:
                unreferenced += 1
                continue

            for port_id, tc in _traffic_classes(flow, pmappers)[0]:
                name = ctx.gemport(n, 1, tc)
                if name in names or port_id in gem_ids:
                    continue
                names.add(name)
                gem_ids.add(port_id)
                self._gemport(gem_container, ctx, n, tc, port_id, subif,
                              tcont_names[allocs[0]])
                emitted += 1

        if unreferenced:
            result.gaps.append(
                f"{unreferenced} service(s) name no T-CONT that the trace "
                "also created, so their GEM ports are omitted: a gemport "
                "without a tcont-ref is rejected")
        if not emitted:
            result.gaps.append(
                "no GEM port could be compiled; the IR has no resolved GEM "
                "port id for any service")

    def _gemport(self, container, ctx: OltContext, uni: int, tc: int,
                 port_id: int, subif: str, tcont_name: str) -> None:
        gem = container.child("gemport")
        gem.leaf("name", ctx.gemport(uni, 1, tc))
        gem.leaf("gemport-id", port_id)
        gem.leaf("interface", subif)
        gem.leaf("traffic-class", tc)
        gem.leaf("tcont-ref", tcont_name)


def main() -> int:
    ap = argparse.ArgumentParser(
        description=__doc__,
        formatter_class=argparse.RawDescriptionHelpFormatter)
    ap.add_argument("input", help="an ir.json, or an omci.json with --from-omci")
    ap.add_argument("--onu-name", required=True,
                    help="the OLT's name for this ONU (not present in OMCI)")
    ap.add_argument("--vendor", default="", help="vendor OUI, e.g. ALCL")
    ap.add_argument("--from-omci", action="store_true",
                    help="treat the input as an OMCI message log")
    ap.add_argument("--chassis-name", default="",
                    help="an existing chassis component to reuse, since a "
                         "second one is rejected outright")
    ap.add_argument("--veip-component", default="",
                    help="the virtual-UNI component a VEIP rides; only an HGU "
                         "reports one")
    ap.add_argument("--yang-dir", default=default_yang_dir())
    ap.add_argument("-o", "--output")
    ap.add_argument("--no-validate", action="store_true")
    ap.add_argument("--json", action="store_true", dest="as_json",
                    help="emit the whole outcome as JSON on stdout, for a "
                         "caller that is not a terminal")
    args = ap.parse_args()

    try:
        if args.from_omci:
            sys.path.insert(0, os.path.join(HERE, "..", "ir"))
            from ir_schema import ir_from_omci_file
            ir = ir_from_omci_file(args.input).to_dict()
        else:
            with open(args.input, encoding="utf-8") as fh:
                ir = json.load(fh)

        index = load_index(args.yang_dir)
        compiler = Compiler(index, validate=not args.no_validate)
        ctx = OltContext(onu_name=args.onu_name, vendor=args.vendor,
                         chassis_name=args.chassis_name,
                         veip_component=args.veip_component)
        result = compiler.compile(ir, ctx)
    except Exception as exc:                                # noqa: BLE001
        # A caller parsing JSON cannot act on a traceback, so failures are
        # reported in the same shape as success.
        if args.as_json:
            json.dump({"error": f"{type(exc).__name__}: {exc}"}, sys.stdout)
            return 2
        raise

    if args.as_json:
        json.dump(result.to_dict(), sys.stdout)
        return 0 if result.valid else 1

    print(result.summary())
    for note in result.notes:
        print(f"  note: {note}")
    for gap in result.gaps:
        print(f"  gap:  {gap}")
    for err in result.validation_errors:
        print(f"  ERROR {err}")

    if args.output:
        with open(args.output, "w", encoding="utf-8") as fh:
            fh.write(result.xml)
        print(f"\nwrote {args.output}")
    else:
        print()
        print(result.xml)
    return 0 if result.valid else 1


if __name__ == "__main__":
    sys.exit(main())
