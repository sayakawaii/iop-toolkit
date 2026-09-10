#!/usr/bin/env python3
"""Schema index over the Nokia LightSpan YANG modules.

The generated config has to be accepted by the OLT's confd, so every element
name, namespace, nesting and list key in it must exist in the real YANG.  This
module builds that authority out of the ~1056 `.yang` files pulled by
`fetch_yang.sh`.

All modules are loaded into a single pyang context and validated in one pass
(~13 s), then the resulting data tree is flattened into `SchemaNode` objects and
cached to disk so later runs start instantly.

Loading everything at once is not laziness avoided for its own sake -- it is
required for correctness.  Most of the interesting nodes arrive by `augment`
from a *different* module than their parent: `bbf-xpon-base` augments
`ietf-interfaces` with `channel-group`, `bbf-sub-interfaces` adds
`subif-lower-layer`, and the whole ONU tree under `onus/onu/root` is contributed
by 50 `-mounted` modules that augment `/onu:onus/onu:onu/onu:root`.  Parse only
the module named by an element's namespace and all of those children look
non-existent.

Because augmented children keep their own defining module's namespace, each
node records the namespace it must actually carry in XML, which is what makes
namespace checking meaningful.
"""

from __future__ import annotations

import json
import os
import pickle
import re
import sys
from dataclasses import dataclass, field

# The ONU configuration tree hangs off this path; the mounted modules augment
# it (documented as RFC 8528 use-schema, but expressed as plain augments).
MOUNT_MODULE = "bbf-fiber-onu-emulated-mount"
MOUNT_POINT_PATH = ("onus", "onu", "root")

CACHE_VERSION = 3

_MODULE_RE = re.compile(r"^\s*(?:sub)?module\s+([\w.-]+)\s*\{", re.MULTILINE)
_IS_SUBMODULE_RE = re.compile(r"^\s*submodule\s", re.MULTILINE)
_NAMESPACE_RE = re.compile(r"^[ \t]*namespace\b", re.MULTILINE)
_PREFIX_RE = re.compile(r"^\s*prefix\s+\"?([\w.-]+)\"?\s*;", re.MULTILINE)
_BELONGS_RE = re.compile(r"^\s*belongs-to\s+([\w.-]+)", re.MULTILINE)
_QUOTED_RE = re.compile(r"\"([^\"]*)\"")


def _parse_namespace(text: str) -> str | None:
    """Read the `namespace` argument, honouring YANG string concatenation.

    Several Nokia modules split the URI across lines with `+`, e.g.

        namespace "http://www.nokia.com/Fixed-Networks/BBA/yang/"
        + "nokia-hardware-extension-mounted";

    Taking only the first quoted run yields a truncated URI, which then fails to
    match the namespace the module actually uses in XML.
    """
    m = _NAMESPACE_RE.search(text)
    if not m:
        return None
    end = text.find(";", m.end())
    if end < 0:
        return None
    return "".join(_QUOTED_RE.findall(text[m.end():end])) or None

_SKIP_KEYWORDS = {"rpc", "notification", "action", "input", "output"}
_KIND_MAP = {
    "container": "container",
    "list": "list",
    "leaf": "leaf",
    "leaf-list": "leaf-list",
    "anydata": "anydata",
    "anyxml": "anydata",
}
_TRANSPARENT = {"choice", "case"}   # schema-only: children appear inline


@dataclass
class ModuleInfo:
    name: str
    path: str
    namespace: str | None = None
    prefix: str | None = None
    is_submodule: bool = False
    belongs_to: str | None = None


@dataclass
class SchemaNode:
    """One data node, flattened out of the pyang statement tree."""

    name: str
    namespace: str
    kind: str                       # container | list | leaf | leaf-list | anydata
    module: str = ""
    keys: list[str] = field(default_factory=list)
    type_name: str | None = None
    enums: list[str] = field(default_factory=list)
    config: bool = True
    mandatory: bool = False
    children: dict[tuple[str, str], "SchemaNode"] = field(default_factory=dict)

    def child(self, namespace: str, name: str) -> "SchemaNode | None":
        node = self.children.get((namespace, name))
        if node is not None:
            return node
        # Tolerate an unqualified/inherited namespace by matching on name alone
        # when it is unambiguous.
        matches = [n for (_, nm), n in self.children.items() if nm == name]
        if len(matches) == 1:
            return matches[0]
        return None

    def child_names(self) -> list[str]:
        return sorted({nm for _, nm in self.children})

    def to_dict(self) -> dict:
        return {
            "n": self.name,
            "ns": self.namespace,
            "k": self.kind,
            "m": self.module,
            "keys": self.keys,
            "t": self.type_name,
            "e": self.enums,
            "cfg": self.config,
            "man": self.mandatory,
            "c": [c.to_dict() for c in self.children.values()],
        }

    @classmethod
    def from_dict(cls, raw: dict) -> "SchemaNode":
        node = cls(
            name=raw["n"], namespace=raw["ns"], kind=raw["k"],
            module=raw.get("m", ""), keys=raw.get("keys") or [],
            type_name=raw.get("t"), enums=raw.get("e") or [],
            config=raw.get("cfg", True), mandatory=raw.get("man", False),
        )
        for sub in raw.get("c") or []:
            child = cls.from_dict(sub)
            node.children[(child.namespace, child.name)] = child
        return node


class SchemaIndex:
    def __init__(self, yang_dir: str, cache_dir: str | None = None,
                 verbose: bool = False):
        self.yang_dir = os.path.abspath(yang_dir)
        self.cache_dir = cache_dir or self.yang_dir
        self.verbose = verbose
        self.modules: dict[str, ModuleInfo] = {}
        self.by_namespace: dict[str, str] = {}
        self.roots: dict[tuple[str, str], SchemaNode] = {}
        self.load_errors: list[str] = []
        self._scan_modules()

    # ------------------------------------------------------------------
    # Module header scan (cheap, no YANG parsing)
    # ------------------------------------------------------------------

    def _scan_modules(self) -> None:
        for fname in sorted(os.listdir(self.yang_dir)):
            if not fname.endswith(".yang"):
                continue
            path = os.path.join(self.yang_dir, fname)
            with open(path, encoding="utf-8", errors="replace") as fh:
                head = fh.read(8192)
            m = _MODULE_RE.search(head)
            if not m:
                continue
            bt = _BELONGS_RE.search(head)
            px = _PREFIX_RE.search(head)
            info = ModuleInfo(
                name=m.group(1),
                path=path,
                namespace=_parse_namespace(head),
                prefix=px.group(1) if px else None,
                is_submodule=bool(_IS_SUBMODULE_RE.match(head)) or bool(bt),
                belongs_to=bt.group(1) if bt else None,
            )
            self.modules[info.name] = info
            if info.namespace and not info.is_submodule:
                self.by_namespace.setdefault(info.namespace, info.name)

    def namespace_of(self, module_name: str) -> str:
        info = self.modules.get(module_name)
        if info is None:
            return ""
        if info.namespace:
            return info.namespace
        if info.belongs_to:                      # submodule inherits
            return self.namespace_of(info.belongs_to)
        return ""

    def module_for_namespace(self, namespace: str) -> str | None:
        return self.by_namespace.get(namespace)

    # ------------------------------------------------------------------
    # Full schema build
    # ------------------------------------------------------------------

    @property
    def cache_path(self) -> str:
        return os.path.join(self.cache_dir, f".schema_cache_v{CACHE_VERSION}.pickle")

    def build(self, force: bool = False) -> None:
        """Populate `self.roots`, using the on-disk cache when it is valid."""
        if not force and self._load_cache():
            return
        self._build_from_yang()
        self._save_cache()

    def _fingerprint(self) -> str:
        newest = 0.0
        for info in self.modules.values():
            newest = max(newest, os.path.getmtime(info.path))
        return f"{len(self.modules)}:{newest:.0f}"

    def _load_cache(self) -> bool:
        try:
            with open(self.cache_path, "rb") as fh:
                blob = pickle.load(fh)
        except (OSError, pickle.UnpicklingError, EOFError):
            return False
        if blob.get("fingerprint") != self._fingerprint():
            return False
        self.roots = {
            tuple(k): SchemaNode.from_dict(v) for k, v in blob["roots"]
        }
        self.load_errors = blob.get("load_errors", [])
        if self.verbose:
            print(f"[schema] cache hit: {len(self.roots)} top-level nodes",
                  file=sys.stderr)
        return True

    def _save_cache(self) -> None:
        blob = {
            "fingerprint": self._fingerprint(),
            "roots": [[list(k), v.to_dict()] for k, v in self.roots.items()],
            "load_errors": self.load_errors,
        }
        try:
            with open(self.cache_path, "wb") as fh:
                pickle.dump(blob, fh, protocol=pickle.HIGHEST_PROTOCOL)
        except OSError:
            pass   # a read-only yang dir just means no caching

    def _build_from_yang(self) -> None:
        from pyang import context, repository

        repo = repository.FileRepository(self.yang_dir)
        ctx = repository and context.Context(repo)

        loaded = []
        for info in self.modules.values():
            if info.is_submodule:
                continue          # pulled in via `include`
            with open(info.path, encoding="utf-8", errors="replace") as fh:
                text = fh.read()
            mod = ctx.add_module(info.path, text)
            if mod is not None:
                loaded.append(mod)

        ctx.validate()
        for err in getattr(ctx, "errors", []) or []:
            self.load_errors.append(str(err))
        if self.verbose:
            print(f"[schema] parsed {len(loaded)} modules, "
                  f"{len(self.load_errors)} loader complaints", file=sys.stderr)

        for mod in loaded:
            for child in getattr(mod, "i_children", []) or []:
                node = self._convert(child)
                if node is None:
                    continue
                if node.kind in _TRANSPARENT:
                    self.roots.update(node.children)
                else:
                    self.roots[(node.namespace, node.name)] = node

    def _node_namespace(self, stmt) -> str:
        """Namespace the node must carry in XML.

        Augmented nodes belong to the augmenting module, so the namespace comes
        from the statement's own module rather than from its parent.
        """
        mod = getattr(stmt, "i_module", None)
        if mod is None:
            return ""
        name = getattr(mod, "arg", "") or ""
        ns = self.namespace_of(name)
        if ns:
            return ns
        nsstmt = mod.search_one("namespace") if hasattr(mod, "search_one") else None
        return nsstmt.arg if nsstmt is not None else ""

    def _convert(self, stmt, depth: int = 0) -> SchemaNode | None:
        kw = stmt.keyword
        if kw in _SKIP_KEYWORDS:
            return None
        if depth > 64:
            return None                       # guard against pathological recursion

        if kw in _TRANSPARENT:
            holder = SchemaNode(name=stmt.arg or "", namespace="", kind=kw)
            for child in getattr(stmt, "i_children", []) or []:
                sub = self._convert(child, depth + 1)
                if sub is None:
                    continue
                if sub.kind in _TRANSPARENT:
                    holder.children.update(sub.children)
                else:
                    holder.children[(sub.namespace, sub.name)] = sub
            return holder

        kind = _KIND_MAP.get(kw)
        if kind is None:
            return None

        mod = getattr(stmt, "i_module", None)
        node = SchemaNode(
            name=stmt.arg,
            namespace=self._node_namespace(stmt),
            kind=kind,
            module=getattr(mod, "arg", "") or "",
        )

        cfg = stmt.search_one("config")
        if cfg is not None and cfg.arg == "false":
            node.config = False
        man = stmt.search_one("mandatory")
        if man is not None and man.arg == "true":
            node.mandatory = True

        if kind == "list":
            key = stmt.search_one("key")
            if key and key.arg:
                node.keys = key.arg.split()

        if kind in ("leaf", "leaf-list"):
            t = stmt.search_one("type")
            if t is not None:
                node.type_name = t.arg
                if t.arg == "enumeration":
                    node.enums = [e.arg for e in t.search("enum")]

        for child in getattr(stmt, "i_children", []) or []:
            sub = self._convert(child, depth + 1)
            if sub is None:
                continue
            if sub.kind in _TRANSPARENT:
                node.children.update(sub.children)
            else:
                node.children[(sub.namespace, sub.name)] = sub
        return node

    # ------------------------------------------------------------------
    # Lookup helpers
    # ------------------------------------------------------------------

    def top_level(self, namespace: str, name: str) -> SchemaNode | None:
        node = self.roots.get((namespace, name))
        if node is not None:
            return node
        matches = [n for (_, nm), n in self.roots.items() if nm == name]
        return matches[0] if len(matches) == 1 else None

    def resolve_path(self, path: str) -> SchemaNode | None:
        """Resolve a `/a/b/c` name path (namespaces inferred) for debugging."""
        parts = [p for p in path.strip("/").split("/") if p]
        if not parts:
            return None
        node = self.top_level("", parts[0])
        for part in parts[1:]:
            if node is None:
                return None
            node = node.child("", part)
        return node

    def mount_point(self) -> SchemaNode | None:
        return self.resolve_path("/" + "/".join(MOUNT_POINT_PATH))

    def stats(self) -> dict:
        total = [0]

        def count(node: SchemaNode) -> None:
            total[0] += 1
            for child in node.children.values():
                count(child)

        for node in self.roots.values():
            count(node)
        return {
            "modules_on_disk": len(self.modules),
            "namespaces": len(self.by_namespace),
            "top_level_nodes": len(self.roots),
            "schema_nodes": total[0],
            "loader_complaints": len(self.load_errors),
        }


def default_yang_dir() -> str:
    here = os.path.dirname(os.path.abspath(__file__))
    return os.path.abspath(os.path.join(here, "..", "..", "yang", "fwlt-c"))


def load_index(yang_dir: str | None = None, force: bool = False,
               verbose: bool = False) -> SchemaIndex:
    idx = SchemaIndex(yang_dir or default_yang_dir(), verbose=verbose)
    idx.build(force=force)
    return idx


if __name__ == "__main__":
    import argparse

    ap = argparse.ArgumentParser(description="Build/inspect the YANG schema index")
    ap.add_argument("--yang-dir", default=default_yang_dir())
    ap.add_argument("--force", action="store_true", help="ignore the cache")
    ap.add_argument("--path", help="resolve a /a/b/c path and show its children")
    args = ap.parse_args()

    index = load_index(args.yang_dir, force=args.force, verbose=True)
    print(json.dumps(index.stats(), indent=2))

    mp = index.mount_point()
    if mp:
        print(f"\nmount point /{'/'.join(MOUNT_POINT_PATH)} has "
              f"{len(mp.children)} mounted top-level nodes:")
        print("  " + ", ".join(mp.child_names()))

    if args.path:
        node = index.resolve_path(args.path)
        if node is None:
            print(f"\n{args.path}: NOT FOUND")
        else:
            print(f"\n{args.path}: kind={node.kind} ns={node.namespace} "
                  f"module={node.module} keys={node.keys} type={node.type_name}")
            print("  children: " + ", ".join(node.child_names()))
