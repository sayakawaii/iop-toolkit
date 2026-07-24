import React from "react";
import axios from "axios";
import DefaultLayout from "@/layouts/default";
import { title } from "@/components/primitives";
import { Button} from "@heroui/react";
import { button as buttonStyles } from "@heroui/theme";
import { API_BASE } from "@/config/api";

type NodeItem = {
  id: string;
  name: string;
  type: "dir" | "file";
  path: string;
  size?: number;
  mtime?: string;
  hasChildren?: boolean;
};

const baseUrl = API_BASE + "/api/downloads";

function IconFolder() {
  return (
    <svg className="w-5 h-5 mr-2 inline-block" viewBox="0 0 24 24" fill="none" stroke="currentColor">
      <path d="M3 7a2 2 0 0 1 2-2h3l2 2h7a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V7z" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"></path>
    </svg>
  );
}
function IconFile() {
  return (
    <svg className="w-5 h-5 mr-2 inline-block" viewBox="0 0 24 24" fill="none" stroke="currentColor">
      <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"></path>
      <path d="M14 2v6h6" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"></path>
    </svg>
  );
}

export default function DownloadsPage() {
  // helper to produce stable unique key for nodes
  const getNodeKey = (n: NodeItem) => n.path || n.id || n.name;

  // tree state
  const [roots, setRoots] = React.useState<NodeItem[]>([]);
  const cacheRef = React.useRef<Map<string, NodeItem[]>>(new Map());
  const loadingRef = React.useRef<Map<string, boolean>>(new Map());
  const [expanded, setExpanded] = React.useState<Record<string, boolean>>({});
  const [selected, setSelected] = React.useState<Record<string, boolean>>({});
  const [search, setSearch] = React.useState("");
  const [loadingRoot, setLoadingRoot] = React.useState(false);
  const [previewContent, setPreviewContent] = React.useState<string | null>(null);
  const [previewOpen, setPreviewOpen] = React.useState(false);
  // previewUrl used when preview content is an SVG (render as image via blob URL)
  const [previewUrl, setPreviewUrl] = React.useState<string | null>(null);
  const [error, setError] = React.useState<string | null>(null);
  const [refreshToken, setRefreshToken] = React.useState(0);

  // fetch listing for a path ("" for root)
  const listPath = React.useCallback(async (path: string) => {
    if (loadingRef.current.get(path)) return;
    loadingRef.current.set(path, true);
    try {
      const resp = await axios.get(`${baseUrl}/list`, { params: { path } });
      const data = resp.data;
      const arr: NodeItem[] = Array.isArray(data) ? data : data ? [data] : [];
      cacheRef.current.set(path, arr);
      // ensure root state is populated immediately when listing root
      if (path === "") {
        setRoots(arr);
      }
      return arr;
    } catch (err: any) {
      console.error("List error", err);
      setError(err?.message ?? "List failed");
      return [];
    } finally {
      loadingRef.current.set(path, false);
    }
  }, []);

  // initial root load or refresh
  React.useEffect(() => {
    let mounted = true;
    (async () => {
      setLoadingRoot(true);
      setError(null);
      try {
        const items = await listPath("");
        if (!mounted) return;
        setRoots(items || []);
      } finally {
        if (mounted) setLoadingRoot(false);
      }
    })();
    return () => {
      mounted = false;
    };
  }, [listPath, refreshToken]);

  // expand handler with lazy load
  const onToggle = React.useCallback(
    async (node: NodeItem) => {
      if (node.type !== "dir") return;
      const newExpanded = { ...expanded, [node.path]: !expanded[node.path] };
      setExpanded(newExpanded);
      if (!cacheRef.current.has(node.path) && node.hasChildren !== false) {
        await listPath(node.path);
        // force re-render to show newly loaded children
        setRefreshToken((t) => t + 1);
      }
    },
    [expanded, listPath]
  );

  // select toggles
  const toggleSelect = React.useCallback((p: string, checked?: boolean) => {
    setSelected((s) => {
      const copy = { ...s };
      if (checked === undefined) copy[p] = !copy[p];
      else if (!checked) delete copy[p];
      else copy[p] = true;
      return copy;
    });
  }, []);

  const selectedPaths = React.useMemo(() => Object.keys(selected).filter(Boolean), [selected]);

  // download single file (blob) and save
  const downloadFile = React.useCallback(async (node: NodeItem) => {
    try {
      const resp = await axios.get(`${baseUrl}/download`, { params: { path: node.path }, responseType: "blob" });
      const blob = resp.data;
      const a = document.createElement("a");
      const url = globalThis.URL.createObjectURL(blob);
      a.href = url;
      a.download = node.name;
      document.body.appendChild(a);
      a.click();
      a.remove();
      globalThis.URL.revokeObjectURL(url);
    } catch (err: any) {
      console.error("Download failed", err);
      setError(err?.message ?? "Download failed");
    }
  }, []);

  // batch download: request server zip or zip locally (use server zip API)
  const downloadSelected = React.useCallback(async () => {
    if (!selectedPaths.length) return;
    try {
      const resp = await axios.post(
        `${baseUrl}/zip`,
        { paths: selectedPaths },
        { responseType: "blob", headers: { "Content-Type": "application/json" } }
      );
      const blob = resp.data;
      const a = document.createElement("a");
      const url = globalThis.URL.createObjectURL(blob);
      a.href = url;
      a.download = `download_${Date.now()}.zip`;
      document.body.appendChild(a);
      a.click();
      a.remove();
      globalThis.URL.revokeObjectURL(url);
    } catch (err: any) {
      console.error("Zip failed", err);
      setError(err?.message ?? "Zip failed");
    }
  }, [selectedPaths]);

  // preview small text file
  const previewFile = React.useCallback(async (node: NodeItem) => {
    try {
      setPreviewContent(null);
      setPreviewUrl(null);
      setPreviewOpen(true);
      const resp = await axios.get(`${baseUrl}/preview`, { params: { path: node.path } });
      const data = resp.data ?? "";
      const text = String(data);
      // detect SVG content (simple heuristic)
      const trimmed = text.trim();
      const isSvg = trimmed.startsWith("<svg") || trimmed.startsWith("<?xml") || trimmed.includes("<svg");
      if (isSvg) {
        // create blob URL to render as image (avoids inserting raw HTML)
        const blob = new Blob([text], { type: "image/svg+xml;charset=utf-8" });
        const url = globalThis.URL.createObjectURL(blob);
        setPreviewUrl(url);
        setPreviewContent(text); // keep raw if needed
      } else {
        setPreviewContent(text);
      }
    } catch (err: any) {
      console.error("Preview failed", err);
      setPreviewContent("Preview not available");
    }
  }, []);

  // revoke previewUrl when it's changed/unmounted or when modal closed
  React.useEffect(() => {
    return () => {
      if (previewUrl) {
        globalThis.URL.revokeObjectURL(previewUrl);
      }
    };
  }, [previewUrl]);

  // refresh a node path (clear cache then refetch)
  const refreshPath = React.useCallback(async (path: string) => {
    cacheRef.current.delete(path);
    if (path === "") {
      setRefreshToken((t) => t + 1);
    } else {
      await listPath(path);
      setRefreshToken((t) => t + 1);
    }
  }, [listPath]);

  // search: simple client-side search across cached nodes; if not cached, fetch root children
  const getVisibleRoots = React.useCallback(() => {
    const q = search.trim().toLowerCase();
    if (!q) return roots;
    const results: NodeItem[] = [];
    const seen = new Set<string>();
    function visit(path: string, items?: NodeItem[]) {
      const list = items ?? cacheRef.current.get(path) ?? [];
      for (const it of list) {
        if (seen.has(it.path)) continue;
        seen.add(it.path);
        if (it.name.toLowerCase().includes(q) || it.path.toLowerCase().includes(q)) {
          results.push(it);
        }
        if (it.type === "dir") {
          visit(it.path, cacheRef.current.get(it.path));
        }
      }
    }
    // ensure at least root cached
    visit("");
    return results;
  }, [roots, search]);

  // recursive TreeNode
  function TreeNode({ node, depth = 0 }: { node: NodeItem; depth?: number }) {
    const isExpanded = !!expanded[node.path];
    const children = cacheRef.current.get(node.path) ?? [];
    const isSelected = !!selected[node.path];

    return (
      <div role="treeitem" aria-expanded={node.type === "dir" ? isExpanded : undefined} className={`w-full`} aria-selected={isSelected}>
        <div
          className={`flex items-center gap-2 px-3 py-2 rounded-md hover:shadow-sm`}
          style={{ marginLeft: depth * 12 }}
        >
          {node.type === "dir" ? (
            <Button
              isIconOnly
              aria-label={isExpanded ? "Collapse" : "Expand"}
              onPress={() => onToggle(node)} size="sm" radius="full" variant="light"
              title={isExpanded ? "Collapse" : "Expand"}
            >
              <svg className={`w-4 h-4 transform ${isExpanded ? "rotate-90" : ""}`} viewBox="0 0 24 24" fill="none" stroke="currentColor">
                <path d="M6 9l6 6 6-6" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"></path>
              </svg>
            </Button>
          ) : (
            <span className="w-6 inline-block" />
          )}

          <input
            type="checkbox"
            checked={isSelected}
            onChange={(e) => toggleSelect(node.path, e.target.checked)}
            className="w-4 h-4"
            aria-label={`select ${node.name}`}
          />

          <span className="flex-1 truncate" title={node.path}>
            {node.type === "dir" ? <IconFolder /> : <IconFile />} <span className="align-middle">{node.name}</span>
          </span>

          <span className="text-sm text-gray-500 dark:text-gray-300 mr-3">{node.mtime ? new Date(node.mtime).toLocaleString() : ""}</span>

          <div className="flex gap-2">
            {node.type === "file" && (
              <>
                <Button
                  onPress={() => downloadFile(node)}
                  className={buttonStyles({
                                      color: "default",
                                      radius: "full",
                                      variant: "shadow",
                                      })}>
                  Download
                </Button>
                <Button
                  onPress={() => previewFile(node)}
                  className={buttonStyles({
                                      color: "default",
                                      radius: "full",
                                      variant: "shadow",
                                      })}>
                  Preview
                </Button>
              </>
            )}
            {node.type === "dir" && (
              <Button onPress={() => refreshPath(node.path)} className={buttonStyles({
                                  color: "default",
                                  radius: "full",
                                  variant: "shadow",
                                  })}>
                Refresh
              </Button>
            )}
          </div>
        </div>

        {/* children */}
        {node.type === "dir" && isExpanded && (
          <div role="group" className="mt-1">
            {loadingRef.current.get(node.path) ? (
              <div className="px-6 py-2 text-sm text-gray-500">Loading...</div>
            ) : children.length ? (
              children.map((c) => <TreeNode key={getNodeKey(c)} node={c} depth={depth + 1} />)
            ) : (
              <div className="px-6 py-2 text-sm text-gray-500">Empty</div>
            )}
          </div>
        )}
      </div>
    );
  }

  const visibleRoots = getVisibleRoots();

  return (
    <DefaultLayout>
      <section className="flex flex-col gap-4 py-8 md:py-10">
        <header className="flex items-center justify-between">
          <span className={title({size:"sm"})}>Downloads&nbsp;</span>
          <div className="flex items-center gap-2">
            <input
              placeholder="Search files / paths"
              value={search}
              onChange={(e) => setSearch(e.target.value)}
              className="px-3 py-2 border rounded bg-white dark:bg-gray-800"
            />
            <Button
              onPress={() => {
                setSelected({});
              }}
              className={buttonStyles({
                                  color: "default",
                                  radius: "full",
                                  variant: "shadow",
                                  })}>
              Clear
            </Button>
            <Button onPress={() => downloadSelected()} disabled={!selectedPaths.length} className={buttonStyles({
                                color: "success",
                                radius: "full",
                                variant: "shadow",
                                })}>
              Download Selected ({selectedPaths.length})
            </Button>
            <Button onPress={() => refreshPath("")} className={buttonStyles({
                                color: "default",
                                radius: "full",
                                variant: "shadow",
                                })}>
              Refresh
            </Button>
          </div>
        </header>

        {error && <div className="text-sm text-red-500">{error}</div>}

        <div className="border rounded-md overflow-hidden" role="tree" aria-label="downloads tree">
          <div className="px-3 py-2 bg-gray-50 dark:bg-gray-900 text-sm text-gray-600 dark:text-gray-300 flex items-center justify-between">
            <div>Path</div>
            <div className="text-xs text-gray-500">Actions</div>
          </div>

          <div className="p-2 max-h-[60vh] overflow-auto">
            {loadingRoot ? (
              <div className="text-sm text-gray-500">Loading root...</div>
            ) : visibleRoots.length ? (
              visibleRoots.map((r) => <TreeNode key={getNodeKey(r)} node={r} />)
            ) : (
              <div className="text-sm text-gray-500">No files</div>
            )}
          </div>
        </div>

        {/* Preview modal */}
        {previewOpen && (
          <div className="fixed inset-0 z-50 flex items-center justify-center">
            <div className="absolute inset-0 bg-black/40" onClick={() => { setPreviewOpen(false); if (previewUrl) { globalThis.URL.revokeObjectURL(previewUrl); setPreviewUrl(null); } }} />
            <div className="relative max-w-3xl w-full bg-white dark:bg-gray-800 rounded shadow-lg p-4 z-10">
              <div className="flex justify-between items-center mb-2">
                <h3 className="text-lg font-medium">Preview</h3>
                <Button onPress={() => { setPreviewOpen(false); if (previewUrl) { globalThis.URL.revokeObjectURL(previewUrl); setPreviewUrl(null); } }} className={buttonStyles({
                                    color: "default",
                                    radius: "full",
                                    variant: "shadow",
                                    })}>
                  Close
                </Button>
              </div>
              <div className="overflow-auto max-h-[60vh] text-sm text-gray-900 dark:text-gray-100">
                {previewUrl ? (
                  <img src={previewUrl} alt="svg preview" className="max-w-full max-h-[60vh] mx-auto" />
                ) : (
                  <pre className="whitespace-pre-wrap">{previewContent ?? "Loading..."}</pre>
                )}
              </div>
            </div>
          </div>
        )}

        <footer className="text-sm text-gray-500">Tip: use checkboxes to select multiple files and Download Selected to get a ZIP.</footer>
      </section>
    </DefaultLayout>
  );
}
