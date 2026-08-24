import DefaultLayout from "@/layouts/default";
import { title } from "@/components/primitives";
import {
  Button,
  Table,
  TableHeader,
  TableColumn,
  TableBody,
  TableRow,
  TableCell,
  Pagination,
  getKeyValue,
} from "@heroui/react";
import { button as buttonStyles } from "@heroui/theme";
import React, { useEffect, Key } from "react";
import axios from "axios";
import { useLocation } from "react-router-dom";
import { API_BASE } from "@/config/api";

/* --- 新增：递归渲染任意结构为嵌套表格的组件 --- */
type AnyData = null | boolean | number | string | AnyData[] | { [k: string]: AnyData };

const cellTableStyle: React.CSSProperties = {
  borderCollapse: "collapse",
  width: "100%",
  fontSize: 13,
};

const cellTdStyle: React.CSSProperties = {
  border: "1px solid #e5e7eb",
  padding: "6px 8px",
  verticalAlign: "top",
  background: "#fff",
  whiteSpace: "normal",
};

function isPrimitive(v: any): v is null | boolean | number | string {
  return v === null || ["string", "number", "boolean"].includes(typeof v);
}

function tryParseJSON(s: any): AnyData | null {
  if (typeof s !== "string") return null;
  try {
    return JSON.parse(s);
  } catch {
    return null;
  }
}

function formatDevId(devId: unknown): string {
  const n = Number(devId);
  if (n === 10) return "Baseline OMCI";
  if (n === 11) return "Extended OMCI";
  if (Number.isFinite(n)) {
    return `Unknown (0x${n.toString(16).toUpperCase().padStart(2, "0")})`;
  }
  return "Unknown";
}

/** Ensure Details shows MsgFormat even when backend has not been redeployed yet. */
function enrichDetailContent(content: AnyData): AnyData {
  if (!content || typeof content !== "object" || Array.isArray(content)) return content;

  const root = content as Record<string, AnyData>;
  const header = root.header;
  if (!header || typeof header !== "object" || Array.isArray(header)) return content;

  const h = header as Record<string, AnyData>;
  const msgFormat = String(h.MsgFormat ?? h.msgFormat ?? formatDevId(h.DevId ?? h.devId));
  const enrichedHeader: Record<string, AnyData> = { ...h, MsgFormat: msgFormat };

  let enrichedPayload = root.payload;
  if (enrichedPayload && typeof enrichedPayload === "object" && !Array.isArray(enrichedPayload)) {
    const p = enrichedPayload as Record<string, AnyData>;
    if (p.MsgFormat === undefined && p.msgFormat === undefined) {
      enrichedPayload = { MsgFormat: msgFormat, ...p };
    }
  }

  return { ...root, header: enrichedHeader, payload: enrichedPayload };
}

const RecursiveTable: React.FC<{ data: AnyData }> = ({ data }) => {
  if (isPrimitive(data)) {
    return <span>{String(data ?? "")}</span>;
  }

  if (Array.isArray(data)) {
    if (data.length === 0) return <span>[]</span>;

    const allObjects = data.every((it) => it && typeof it === "object" && !Array.isArray(it));
    if (allObjects) {
      const keySet = new Set<string>();
      (data as any[]).forEach((it) => Object.keys(it).forEach((k) => keySet.add(k)));
      const keys = Array.from(keySet);

      return (
        <table style={cellTableStyle}>
          <thead>
            <tr>
              {keys.map((k) => (
                <th key={k} style={cellTdStyle}>{k}</th>
              ))}
            </tr>
          </thead>
          <tbody>
            {(data as any[]).map((row, ri) => (
              <tr key={ri}>
                {keys.map((k) => (
                  <td key={k} style={cellTdStyle}>
                    <RecursiveTable data={(row && (row[k] ?? null)) as AnyData} />
                  </td>
                ))}
              </tr>
            ))}
          </tbody>
        </table>
      );
    }

    return (
      <table style={cellTableStyle}>
        <tbody>
          {(data as any[]).map((it, i) => (
            <tr key={i}>
              <td style={cellTdStyle}>
                <RecursiveTable data={it as AnyData} />
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    );
  }

  return (
    <table style={cellTableStyle}>
      <tbody>
        {Object.keys(data as object).map((k) => (
          <tr key={k}>
            <td style={{ ...cellTdStyle, fontWeight: 600, width: 160 }}>{k}</td>
            <td style={cellTdStyle}>
              <RecursiveTable data={(data as any)[k] as AnyData} />
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default function OmciAnalyzerPageOmciPage() {
  const [rows, setRows] = React.useState<any[]>([]);
  const [showTable, setShowTable] = React.useState(false);
  const [loading, setLoading] = React.useState(false);
  const [error, setError] = React.useState<string | null>(null);

  // client-side features
  const [search, setSearch] = React.useState("");
  const [page, setPage] = React.useState(1);
  const [pageSize, setPageSize] = React.useState<number>(20);
  const pageSizeOptions = [20, 100, 1000, 10000];
  const [sortKey, setSortKey] = React.useState<string | null>(null);
  const [sortAsc, setSortAsc] = React.useState(true);

  // modal for Details (lazy show)
  const [modalOpen, setModalOpen] = React.useState(false);
  // changed: modalContent can be raw object (AnyData) or string
  const [modalContent, setModalContent] = React.useState<AnyData | string | null>(null);
  const [modalLoading, setModalLoading] = React.useState(false);

  const location = useLocation();
  const params = new URLSearchParams(location.search);
  const requestKey = params.get("requestKey") ?? "";
  const onuName = params.get("onuName") ?? "";

  const baseUrl = API_BASE + "/api/omcianalyzer";

  // columns (note: action handled manually)
  const columns = [
    { key: "id", label: "ID" },
    { key: "latency", label: "Latency(ms)" },
    { key: "tcid", label: "TCID(hex)" },
    { key: "name", label: "Name" },
    { key: "class", label: "Class" },
    { key: "type", label: "Type" },
    { key: "msgFormat", label: "Msg Format" },
    { key: "direction", label: "Direction" },
    { key: "status", label: "Status" },
    { key: "action", label: "Action" },
  ];

  useEffect(() => {
    if (!requestKey) {
      setShowTable(false);
      return;
    }

    const fetchData = async () => {
      setLoading(true);
      setError(null);
      try {
        const resp = await axios.get(`${baseUrl}/omci`, {
          params: { requestKey, onuName },
        });

        const data = resp.data;
        console.error("Fetched OMCI data:", data);
        if (Array.isArray(data)) {
          setRows(data);
          setShowTable(true);
        } else if (data) {
          setRows([data]);
          setShowTable(true);
        } else {
          setRows([]);
          setShowTable(false);
        }
      } catch (err: any) {
        console.error("Failed to load OMCI data:", err);
        setError(err?.message ?? "Request failed");
        setShowTable(false);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
    setPage(1);
    setSearch("");
    setSortKey(null);
  }, [requestKey]);

  const filtered = React.useMemo(() => {
    const q = search.trim().toLowerCase();
    let list = rows.slice();

    if (q) {
      list = list.filter((r) =>
        Object.values(r).some((v) =>
          String(v ?? "").toLowerCase().includes(q)
        )
      );
    }

    if (sortKey) {
      list.sort((a: any, b: any) => {
        const va = String(getKeyValue(a, sortKey) ?? "");
        const vb = String(getKeyValue(b, sortKey) ?? "");
        if (va === vb) return 0;
        const res = va > vb ? 1 : -1;
        return sortAsc ? res : -res;
      });
    }

    return list;
  }, [rows, search, sortKey, sortAsc]);

  const totalPages = Math.max(1, Math.ceil(filtered.length / pageSize));
  const pageItems = filtered.slice((page - 1) * pageSize, page * pageSize);

  React.useEffect(() => {
    setPage(1);
  }, [pageSize]);

  const toggleSort = (key: string) => {
    if (sortKey === key) {
      setSortAsc(!sortAsc);
    } else {
      setSortKey(key);
      setSortAsc(true);
    }
    setPage(1);
  };

  const exportCsv = () => {
    const cols = columns.filter((c) => c.key !== "action");
    const rowsToExport = pageItems;
    const csvLines = [
      cols.map((c) => `"${c.label.replace(/"/g, '""')}"`).join(","),
      ...rowsToExport.map((r) =>
        cols
          .map((c) => `"${String(getKeyValue(r, c.key) ?? "").replace(/"/g, '""')}"`)
          .join(",")
      ),
    ];
    const blob = new Blob([csvLines.join("\n")], { type: "text/csv;charset=utf-8;" });
    const url = URL.createObjectURL(blob);
    const a = document.createElement("a");
    a.href = url;
    a.download = `omci_${requestKey || "export"}.csv`;
    document.body.appendChild(a);
    a.click();
    a.remove();
    URL.revokeObjectURL(url);
  };

  const openDetails = async (item: any) => {
    setModalOpen(true);
    setModalContent(null);

    // if content present, show immediately; handle object/string properly
    if (item.content !== undefined && item.content !== null) {
      if (typeof item.content === "object") {
        setModalContent(enrichDetailContent(item.content as AnyData));
      } else {
        const parsed = tryParseJSON(item.content);
        setModalContent(parsed ? enrichDetailContent(parsed) : String(item.content));
      }
      return;
    }

    setModalLoading(true);
    try {
      const resp = await axios.get(`${baseUrl}/omci/detail`, {
        params: { requestKey, id: item.id ?? item.tcid ?? item.name },
      });

      // resp.data may contain .content as object or string; preserve object when available
      const raw = resp.data?.content ?? resp.data ?? null;
      if (raw !== null && typeof raw === "object") {
        setModalContent(enrichDetailContent(raw as AnyData));
      } else if (typeof raw === "string") {
        const parsed = tryParseJSON(raw);
        setModalContent(parsed ? enrichDetailContent(parsed) : raw);
      } else {
        setModalContent(null);
      }
    } catch (err: any) {
      setModalContent("Failed to load detail: " + (err?.message ?? ""));
    } finally {
      setModalLoading(false);
    }
  };

  return (
    <DefaultLayout>
      <div className="w-full p-5">
        <span className={title({size:"sm"})}>OMCI Records&nbsp;</span>

        <div className="mt-3 mb-2 flex gap-2 items-center">
          <input
            placeholder="Search all fields..."
            value={search}
            onChange={(e) => { setSearch(e.target.value); setPage(1); }}
            className="flex-1 p-2 rounded border border-gray-300"
          />

          <div className="inline-flex items-center gap-2">
            <label htmlFor="perPageSelect" className="text-sm text-gray-700">Per page:</label>
            <select
              id="perPageSelect"
              value={pageSize}
              onChange={(e) => setPageSize(Number(e.target.value))}
              className="p-1 rounded border border-gray-300"
            >
              {pageSizeOptions.map((opt) => (
                <option key={opt} value={opt}>{opt}</option>
              ))}
            </select>
          </div>

          <Button onPress={exportCsv} className={buttonStyles({ color: "default", radius: "full", variant: "shadow" })}>
            Export CSV (page)
          </Button>
        </div>

        {loading && <div className="text-sm text-gray-600">Loading...</div>}
        {error && <div className="text-sm text-red-600">Error: {error}</div>}

        <div className={`${showTable ? "block" : "hidden"} w-full mt-4 overflow-x-auto`}>
          <div className="min-w-[900px]">
            <Table aria-label="OMCI table" className="w-full"
            bottomContent={
                <div className="flex w-full justify-center">
                <Pagination
                    isCompact
                    showControls
                    showShadow
                    color="secondary"
                    page={page}
                    total={totalPages}
                    onChange={(page) => setPage(page)}
                />
                </div>
            }>
              <TableHeader columns={columns}>
                {(column) => (
                  <TableColumn key={String(column.key)}>
                    <div
                      className={column.key !== "action" ? "inline-flex gap-2 items-center cursor-pointer" : "inline-flex gap-2 items-center"}
                      onClick={() => column.key !== "action" && toggleSort(String(column.key))}
                    >
                      {column.label}
                      {sortKey === column.key ? (sortAsc ? " ▲" : " ▼") : null}
                    </div>
                  </TableColumn>
                )}
              </TableHeader>

              <TableBody items={pageItems}>
                {(item: any) => {
                  const dir = String(item.direction ?? "");
                  // light / dark variants to keep distinct colors and readable contrast in dark mode
                  const rowBg =
                    dir.includes("OLT --> ONU")
                      ? "bg-blue-300/50 dark:bg-blue-300 dark:text-blue-800"
                      : dir.includes("OLT <-- ONU") ? "" : "";
                  return (
                    <TableRow key={item.id ?? item.tcid ?? item.name} className={rowBg}>
                      {(columnKey: Key) => {
                        if (columnKey === "action") {
                          return (
                            <TableCell className="whitespace-nowrap">
                              <Button onPress={() => openDetails(item)} className={buttonStyles({ color: "default", radius: "full", variant: "shadow" })}>
                                Details
                              </Button>
                            </TableCell>
                          );
                        }

                        if (columnKey === "status") {
                          const s = String(getKeyValue(item, "status") ?? "").toLowerCase();
                          const colorClass = s.includes("success") ? "text-green-600" : s.includes("warning") ? "text-orange-500" : "";
                          return <TableCell className={colorClass + " "}>{String(getKeyValue(item, "status") ?? "")}</TableCell>;
                        }

                        if (columnKey === "tcid") {
                          const v = getKeyValue(item, "tcid");
                          const str = v == null ? "" : String(v);
                          return <TableCell>{str}</TableCell>;
                        }

                        const val = getKeyValue(item, String(columnKey));
                        const parsed = typeof val === "string" ? tryParseJSON(val) : null;
                        const toRender = parsed ?? (val === undefined ? null : val);

                        if (toRender !== null && typeof toRender === "object") {
                          return (
                            <TableCell className="whitespace-normal p-2">
                              <RecursiveTable data={toRender as AnyData} />
                            </TableCell>
                          );
                        }

                        return <TableCell className="whitespace-nowrap">{String(toRender ?? "")}</TableCell>;
                      }}
                    </TableRow>
                  );
                }}
              </TableBody>
            </Table>
          </div>
        </div>

        {!loading && !showTable && !error && requestKey && (
          <div className="mt-3 text-sm text-gray-600">No OMCI records returned for requestKey.</div>
        )}

        {modalOpen && (
          <div
            role="dialog"
            aria-modal="true"
            className="fixed inset-0 bg-black/50 flex items-center justify-center z-50"
            onClick={() => setModalOpen(false)}
          >
            <div
              onClick={(e) => e.stopPropagation()}
              className="max-w-[80%] max-h-[80%] overflow-auto bg-white p-4 rounded-md"
            >
              <div className="flex justify-between items-center mb-3">
                <h3>Details</h3>
                <Button onPress={() => setModalOpen(false)}>Close</Button>
              </div>
              {modalLoading ? (
                <div>Loading...</div>
              ) : (
                // support modalContent being object (render via RecursiveTable) or string
                (() => {
                  if (modalContent === null) return <div>No content</div>;
                  if (typeof modalContent === "string") {
                    const parsed = tryParseJSON(modalContent);
                    if (parsed) return <RecursiveTable data={parsed} />;
                    return <pre className="whitespace-pre-wrap break-words">{modalContent}</pre>;
                  }
                  // object / array
                  return <RecursiveTable data={modalContent as AnyData} />;
                })()
              )}
            </div>
          </div>
        )}
      </div>
    </DefaultLayout>
  );
}
