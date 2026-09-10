#!/usr/bin/env python3
"""Build configuration XML by walking the YANG schema.

Namespaces are the most common way a generated Nokia config gets rejected, and
hardcoding them does not survive contact with this data model: below
`onus/onu/root`, children routinely belong to a different module than their
parent. `subif-lower-layer` sits inside an `ietf-interfaces-mounted` interface
but belongs to `bbf-sub-interfaces-mounted`; `channel-group` comes from
`bbf-xpon-base`. A template that inherits the parent's namespace is wrong in
exactly those places, and wrong in a way that only shows up on the device.

So this builder never takes a namespace as input. It resolves each element
against the schema index and uses the namespace the schema says the node has.
The result is that a document is namespace-correct by construction, and an
element that does not exist fails here -- at the point of construction, with a
list of what was valid -- rather than in the validator afterwards.
"""

from __future__ import annotations

import os
import sys

from lxml import etree

sys.path.insert(0, os.path.join(os.path.dirname(os.path.abspath(__file__)),
                                "..", "validate"))
from schema_index import SchemaIndex, SchemaNode  # noqa: E402

NETCONF_NS = "urn:ietf:params:xml:ns:netconf:base:1.0"


class SchemaBuildError(RuntimeError):
    pass


class Cursor:
    """A position in both the XML document and the YANG schema."""

    __slots__ = ("element", "node", "index", "path")

    def __init__(self, element, node: SchemaNode, index: SchemaIndex, path: str):
        self.element = element
        self.node = node
        self.index = index
        self.path = path

    def child(self, name: str, text=None) -> "Cursor":
        """Add a child element, taking its namespace from the schema."""
        target = self.node.child(self.node.namespace, name)
        if target is None:
            matches = [n for (_, nm), n in self.node.children.items()
                       if nm == name]
            if len(matches) == 1:
                target = matches[0]
            elif len(matches) > 1:
                raise SchemaBuildError(
                    f"{self.path}/{name} is ambiguous across namespaces "
                    f"{[m.namespace for m in matches]}")
        if target is None:
            known = ", ".join(self.node.child_names()[:15])
            raise SchemaBuildError(
                f"{self.path}: no child {name!r} in the schema; "
                f"valid children include [{known}]")

        el = etree.SubElement(self.element, f"{{{target.namespace}}}{name}")
        if text is not None:
            el.text = str(text)
        return Cursor(el, target, self.index, f"{self.path}/{name}")

    def leaf(self, name: str, value) -> "Cursor":
        """Set a leaf, skipping it when the value is absent.

        Emitting an empty element for a missing value is worse than omitting
        it: confd reads it as "set this leaf to the empty string".
        """
        if value is None or value == "":
            return self
        self.child(name, value)
        return self

    def leaves(self, **values) -> "Cursor":
        for name, value in values.items():
            self.leaf(name.replace("_", "-"), value)
        return self


class ConfigBuilder:
    """An `<edit-config>` document under construction."""

    def __init__(self, index: SchemaIndex, target: str = "running"):
        self.index = index
        self.rpc = etree.Element(f"{{{NETCONF_NS}}}edit-config",
                                 nsmap={"nc": NETCONF_NS})
        target_el = etree.SubElement(self.rpc, f"{{{NETCONF_NS}}}target")
        etree.SubElement(target_el, f"{{{NETCONF_NS}}}{target}")
        self.config = etree.SubElement(self.rpc, f"{{{NETCONF_NS}}}config")

    def top(self, name: str) -> Cursor:
        """Start a top-level data tree, e.g. `onus`."""
        matches = [(ns, node) for (ns, nm), node in self.index.roots.items()
                   if nm == name]
        if not matches:
            raise SchemaBuildError(f"{name!r} is not a top-level data node")
        if len(matches) > 1:
            raise SchemaBuildError(
                f"{name!r} is ambiguous across {[ns for ns, _ in matches]}")
        ns, node = matches[0]
        el = etree.SubElement(self.config, f"{{{ns}}}{name}")
        return Cursor(el, node, self.index, f"/{name}")

    def tostring(self, pretty: bool = True) -> str:
        return etree.tostring(self.rpc, pretty_print=pretty,
                              encoding="unicode")

    def write(self, path: str) -> None:
        with open(path, "w", encoding="utf-8") as fh:
            fh.write(self.tostring())
