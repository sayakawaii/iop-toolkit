#!/usr/bin/env python3
"""Validate a NETCONF config document against the real LightSpan YANG schema.

This is the acceptance gate for everything the compiler emits.  confd rejects a
document for structural reasons long before it looks at semantics -- a
misspelled element, a leaf in the wrong namespace, a list without its key -- so
those are what is checked here:

* every element exists in the schema at the position where it appears
* it carries the namespace of the module that actually defines it (which, for
  augmented nodes, is not the parent's module)
* every list instance carries its key leaves
* leaves have no child elements, and containers/lists are not used as leaves
* enumeration leaves hold one of the declared values
* `config false` nodes are flagged, since they cannot appear in an edit-config

Namespaces outside the YANG set (notably tail-f's `confdConfig`) are reported as
warnings rather than errors: confd accepts them, but they are not part of the
published data model, so there is nothing to check them against.

Elements carrying `nc:operation="delete"` or `"remove"` are structurally checked
but exempt from the missing-key rule below the deleted node, matching how a
delete only needs enough of the tree to identify its target.

Usage:
    python3 validate_xml.py <file.xml> [more.xml ...]
    python3 validate_xml.py --quiet <file.xml>        # summary only
    python3 validate_xml.py --json <file.xml>         # machine-readable
"""

from __future__ import annotations

import argparse
import json
import os
import sys

from lxml import etree

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
from schema_index import SchemaIndex, SchemaNode, default_yang_dir, load_index  # noqa: E402

NETCONF_NS = "urn:ietf:params:xml:ns:netconf:base:1.0"

# Accepted by confd but not part of the modelled data tree.
TOLERATED_NS = {
    "http://tail-f.com/ns/confd_dyncfg/1.0",
}

WRAPPER_TAGS = {"rpc", "edit-config", "config", "data", "get-config", "copy-config"}
EDIT_CONFIG_META = {"target", "default-operation", "test-option", "error-option",
                    "source", "filter"}
DELETE_OPS = {"delete", "remove"}


class Finding:
    __slots__ = ("severity", "path", "message")

    def __init__(self, severity: str, path: str, message: str):
        self.severity = severity
        self.path = path
        self.message = message

    def __str__(self) -> str:
        return f"[{self.severity}] {self.path}: {self.message}"

    def to_dict(self) -> dict:
        return {"severity": self.severity, "path": self.path,
                "message": self.message}


def qname(el) -> tuple[str, str]:
    tag = el.tag
    if isinstance(tag, str) and tag.startswith("{"):
        ns, _, name = tag[1:].partition("}")
        return ns, name
    return "", str(tag)


def element_children(el) -> list:
    return [c for c in el if isinstance(c.tag, str)]


def operation_of(el) -> str | None:
    for key, value in el.attrib.items():
        if isinstance(key, str) and key.endswith("operation"):
            return value
    return None


class XmlValidator:
    def __init__(self, index: SchemaIndex):
        self.index = index
        self.findings: list[Finding] = []

    def add(self, severity: str, path: str, message: str) -> None:
        self.findings.append(Finding(severity, path, message))

    # ------------------------------------------------------------------

    def validate_file(self, path: str) -> list[Finding]:
        self.findings = []
        try:
            tree = etree.parse(path)
        except (etree.XMLSyntaxError, OSError) as exc:
            self.add("error", os.path.basename(path), f"cannot parse: {exc}")
            return self.findings
        return self.validate_element(tree.getroot())

    def validate_string(self, xml: str) -> list[Finding]:
        self.findings = []
        try:
            root = etree.fromstring(xml.encode("utf-8"))
        except etree.XMLSyntaxError as exc:
            self.add("error", "<string>", f"cannot parse: {exc}")
            return self.findings
        return self.validate_element(root)

    def validate_element(self, root) -> list[Finding]:
        data_roots = self._unwrap(root)
        if not data_roots:
            self.add("error", "/", "no configuration data found in document")
            return self.findings
        for el in data_roots:
            self._walk_top(el)
        return self.findings

    def _unwrap(self, root) -> list:
        """Strip <rpc>/<edit-config>/<config> down to the data nodes."""
        _, name = qname(root)
        if name not in WRAPPER_TAGS:
            return [root]

        for tag in ("config", "data"):
            found = root.find(f"{{{NETCONF_NS}}}{tag}")
            if found is None:
                found = root.find(tag)
            if found is not None:
                return element_children(found)

        if name in ("config", "data"):
            return element_children(root)

        out = []
        for child in element_children(root):
            _, cname = qname(child)
            if cname in EDIT_CONFIG_META:
                continue
            if cname in WRAPPER_TAGS:
                nested = self._unwrap(child)
                out.extend(nested)
                continue
            out.append(child)
        return out

    def _walk_top(self, el) -> None:
        ns, name = qname(el)
        if ns in TOLERATED_NS:
            self.add("warning", f"/{name}",
                     f"namespace {ns} is outside the YANG set "
                     "(confd internal); not validated")
            return
        if not ns:
            self.add("error", f"/{name}",
                     "top-level element has no namespace")
            return
        node = self.index.top_level(ns, name)
        if node is None:
            module = self.index.module_for_namespace(ns)
            if module is None:
                self.add("error", f"/{name}", f"unknown namespace {ns!r}")
            else:
                self.add("error", f"/{name}",
                         f"not a top-level data node "
                         f"(namespace maps to module {module})")
            return
        self._walk(el, node, f"/{name}", ns)

    def _walk(self, el, node: SchemaNode, path: str, parent_ns: str,
              in_delete: bool = False) -> None:
        children = element_children(el)
        op = operation_of(el)
        deleting = in_delete or (op in DELETE_OPS)

        if not node.config:
            self.add("error", path,
                     f"{node.kind} is 'config false' and cannot appear "
                     "in a configuration edit")

        if node.kind in ("leaf", "leaf-list"):
            if children:
                self.add("error", path,
                         f"{node.kind} has child elements "
                         f"{[qname(c)[1] for c in children]}")
                return
            self._check_leaf_value(el, node, path)
            return

        if node.kind == "anydata":
            return                      # opaque by definition

        if node.kind == "list" and not deleting:
            present = {qname(c)[1] for c in children}
            missing = [k for k in node.keys if k not in present]
            if missing:
                self.add("error", path,
                         f"list instance is missing key leaf(s) {missing} "
                         f"(keys: {node.keys})")

        if not children and node.kind == "container" and not deleting:
            # An empty container is legal (presence containers), so this is
            # informational rather than an error.
            pass

        for child in children:
            cns, cname = qname(child)
            if cns in TOLERATED_NS:
                self.add("warning", f"{path}/{cname}",
                         f"namespace {cns} is outside the YANG set; "
                         "not validated")
                continue
            lookup_ns = cns or parent_ns
            target = node.child(lookup_ns, cname)
            if target is None:
                known = node.child_names()
                shown = ", ".join(known[:12]) + (" ..." if len(known) > 12 else "")
                hint = ""
                if cns and self.index.module_for_namespace(cns) is None:
                    hint = f" (namespace {cns!r} is not in the YANG set)"
                self.add("error", f"{path}/{cname}",
                         f"no such node under {node.name!r}{hint}; "
                         f"known children: [{shown}]")
                continue
            if cns and target.namespace and cns != target.namespace:
                self.add("error", f"{path}/{cname}",
                         f"wrong namespace {cns!r}; "
                         f"{cname!r} is defined by {target.module} "
                         f"as {target.namespace!r}")
                continue
            self._walk(child, target, f"{path}/{cname}",
                       target.namespace or lookup_ns, deleting)

    def _check_leaf_value(self, el, node: SchemaNode, path: str) -> None:
        text = (el.text or "").strip()
        if not text:
            return
        if node.enums and text not in node.enums:
            allowed = ", ".join(node.enums[:10])
            self.add("error", path,
                     f"value {text!r} is not a valid enum; allowed: [{allowed}]")
            return
        if node.type_name in ("uint8", "uint16", "uint32", "uint64",
                              "int8", "int16", "int32", "int64"):
            probe = text[1:] if text[:1] in "+-" else text
            if not probe.isdigit():
                self.add("error", path,
                         f"value {text!r} is not a valid {node.type_name}")
            elif node.type_name.startswith("uint") and text.startswith("-"):
                self.add("error", path,
                         f"value {text!r} is negative but type is "
                         f"{node.type_name}")
        elif node.type_name == "boolean" and text not in ("true", "false"):
            self.add("error", path,
                     f"value {text!r} is not a boolean (expected true/false)")


def main() -> int:
    ap = argparse.ArgumentParser(
        description=__doc__,
        formatter_class=argparse.RawDescriptionHelpFormatter)
    ap.add_argument("files", nargs="+")
    ap.add_argument("--yang-dir", default=default_yang_dir())
    ap.add_argument("--quiet", action="store_true",
                    help="print only the per-file summary")
    ap.add_argument("--json", action="store_true",
                    help="emit findings as JSON")
    ap.add_argument("--warnings-as-errors", action="store_true")
    ap.add_argument("--max-findings", type=int, default=40,
                    help="cap printed findings per file (0 = no cap)")
    args = ap.parse_args()

    index = load_index(args.yang_dir, verbose=not args.json)
    validator = XmlValidator(index)

    report = []
    total_err = total_warn = 0
    for path in args.files:
        findings = validator.validate_file(path)
        errors = [f for f in findings if f.severity == "error"]
        warnings = [f for f in findings if f.severity == "warning"]
        total_err += len(errors)
        total_warn += len(warnings)
        report.append({
            "file": path,
            "status": "pass" if not errors else "fail",
            "errors": len(errors),
            "warnings": len(warnings),
            "findings": [f.to_dict() for f in findings],
        })
        if args.json:
            continue
        status = "PASS" if not errors else "FAIL"
        print(f"{status}  {os.path.basename(path)}  "
              f"({len(errors)} errors, {len(warnings)} warnings)")
        if not args.quiet:
            shown = findings if args.max_findings == 0 else findings[:args.max_findings]
            for f in shown:
                print(f"   {f}")
            if len(findings) > len(shown):
                print(f"   ... {len(findings) - len(shown)} more")

    if args.json:
        print(json.dumps(report, indent=2))
    else:
        print(f"\ntotal: {total_err} errors, {total_warn} warnings")

    if total_err or (args.warnings_as_errors and total_warn):
        return 1
    return 0


if __name__ == "__main__":
    sys.exit(main())
