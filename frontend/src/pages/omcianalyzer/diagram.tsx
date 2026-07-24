import DefaultLayout from "@/layouts/default";
import React, { useEffect } from "react";
import axios from "axios";
import { useLocation } from "react-router-dom";
import { Image } from "@heroui/react";
import { title } from "@/components/primitives";
import { API_BASE } from "@/config/api";

type ImgItem = {
  url: string;
  category?: string;
  title?: string;
  [k: string]: any;
};

export default function OmciAnalyzerDiagramPage() {
  const location = useLocation();
  const params = new URLSearchParams(location.search);
  const requestKey = params.get("requestKey") ?? "";
  const onuName = params.get("onuName") ?? "";

  const baseUrl = API_BASE + "/api/omcianalyzer";
  const apiOrigin = (() => {
    try {
      return new URL(baseUrl).origin;
    } catch {
      return "";
    }
  })();

  const [items, setItems] = React.useState<ImgItem[]>([]);
  const [loading, setLoading] = React.useState(false);
  const [error, setError] = React.useState<string | null>(null);
  const [category, setCategory] = React.useState<string>("all");

  useEffect(() => {
    if (!requestKey) {
      setItems([]);
      setError("missing requestKey");
      return;
    }

    const fetchImages = async () => {
      setLoading(true);
      setError(null);
      try {
        const resp = await axios.get(`${baseUrl}/diagram`, {
          params: { requestKey, onuName },
        });

        let data = resp.data;
        if (!Array.isArray(data)) data = [];

        // backend now returns objects: { path: string, category: string }
        const normalized: ImgItem[] = (data as any[]).map((it) => {
          const rawPath = it?.path ?? "";
          let url = String(rawPath || "");
          // if path is root-relative, prefix API origin
          if (url && !/^https?:\/\//i.test(url) && url.startsWith("/")) {
            url = apiOrigin + url;
          }
          return {
            url,
            category: it?.category ?? "latest",
            title: it?.title ?? "",
            ...it,
          };
        }).filter((it) => it.url);

        setItems(normalized);
        setLoading(false);
        setError(null);
      } catch (err: any) {
        console.error("Failed to load diagrams:", err);
        setError(err?.message ?? "Request failed");
        setLoading(false);
        setItems([]);
      }
    };

    fetchImages();
  }, [requestKey, onuName, apiOrigin]);

  const categories = React.useMemo(() => {
    const set = new Set<string>(["all", "latest", "tcont", "update", "similar"]);
    items.forEach((it) => {
      if (it.category) set.add(it.category);
    });
    return Array.from(set);
  }, [items]);

  const filtered = React.useMemo(() => {
    if (category === "all") return items;
    return items.filter((it) => (it.category ?? "latest") === category);
  }, [items, category]);

  return (
    <DefaultLayout>
      <div className="w-full p-5">
        <span className={title({size:"sm"})}>Diagram Gallery&nbsp;</span>

        <div className="mt-2 mb-3 flex gap-2 flex-wrap items-center">
          {categories.map((c) => (
            <button
              key={c}
              data-category={c}
              onClick={() => setCategory(c)}
              className={
                category === c
                  ? "px-3 py-1 rounded-md border-2 border-blue-500 bg-blue-50 text-sm"
                  : "px-3 py-1 rounded-md border border-gray-300 bg-white text-sm"
              }
            >
              {c[0].toUpperCase() + c.slice(1)}
            </button>
          ))}
        </div>

        {loading && <div className="text-sm text-gray-600 mt-3">Loading diagrams...</div>}
        {error && <div className="text-sm text-red-600 mt-3">Error: {error}</div>}
        {!loading && !error && filtered.length === 0 && (
          <div className="text-sm text-gray-600 mt-3">No diagrams available.</div>
        )}

        <div className={filtered.length ? "mt-3 flex flex-wrap gap-3" : "hidden"}>
          {filtered.map((it, idx) => (
            <div
              key={it.url + idx}
              className="w-64 border border-gray-200 rounded-md p-2 bg-white flex flex-col gap-2"
            >
              <div className="h-40 flex items-center justify-center overflow-hidden bg-gray-50">
                <Image
                  isZoomed
                  alt={it.title ?? `diagram-${idx}`}
                  src={it.url}
                  className="max-w-full max-h-full object-contain block"
                  onClick={() => globalThis.open(it.url, "_blank", "noopener")}
                />
              </div>

              <div className="flex justify-between items-center gap-2">
                <div className="flex-1 overflow-hidden text-sm truncate">
                  <strong>{it.title ?? (it.category ?? "")}</strong>
                </div>
              </div>

              <div className="text-xs text-gray-600">Category: {it.category ?? "latest"}</div>
            </div>
          ))}
        </div>
      </div>
    </DefaultLayout>
  );
}