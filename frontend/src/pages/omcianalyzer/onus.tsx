import DefaultLayout from "@/layouts/default";
import {Button, Table, TableHeader, TableColumn, TableBody, TableRow, TableCell, Pagination, getKeyValue} from "@heroui/react";
import { button as buttonStyles } from "@heroui/theme";
import React, { useEffect, Key } from "react";
import axios from "axios";
import { useLocation } from "react-router-dom";
import { title } from "@/components/primitives";
import { API_BASE } from "@/config/api";
import GenerateYangModal from "@/pages/omcianalyzer/GenerateYangModal";

export default function OmciAnalyzerPageOnusPage() {
  const [rows, setRows] = React.useState<any[]>([]);
  const [showTable, setShowTable] = React.useState(false);
  const [loading, setLoading] = React.useState(false);
  const [error, setError] = React.useState<string | null>(null);
  const [yangOnu, setYangOnu] = React.useState<string | null>(null);

  // client-side features
  const [search, setSearch] = React.useState("");
  const [page, setPage] = React.useState(1);
  const [pageSize, setPageSize] = React.useState<number>(10);
  const pageSizeOptions = [10, 20, 50, 100];
  const [sortKey, setSortKey] = React.useState<string | null>(null);
  const [sortAsc, setSortAsc] = React.useState(true);

  const location = useLocation();
  const params = new URLSearchParams(location.search);
  const requestKey = params.get("requestKey") ?? "";

  const baseUrl = API_BASE + "/api/omcianalyzer";

  // 定义用于自动渲染的列（前4列自动填充）
  const columns = [
    { key: 'onuName', label: 'OnuName' },
    { key: 'swVersion', label: 'SwVersion' },
    { key: 'hwVersion', label: 'HwVersion' },
    { key: 'status', label: 'Status' },
    { key: 'action', label: 'Action' } // 用于手动渲染 action 列
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
        // 使用 GET 获取 ONUs（语义正确）
        const resp = await axios.get(`${baseUrl}/onus`, {
          params: { requestKey },
        });

        const data = resp.data;
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
        console.error("Failed to load ONUs:", err);
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
        // ONU names are often bare numbers, where a plain string sort puts
        // "10" ahead of "9" -- which is exactly the case this page is for.
        const res = va.localeCompare(vb, undefined, { numeric: true });
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

  // A shrinking result set can leave the current page past the end.
  React.useEffect(() => {
    if (page > totalPages) setPage(totalPages);
  }, [page, totalPages]);

  const toggleSort = (key: string) => {
    if (sortKey === key) {
      setSortAsc(!sortAsc);
    } else {
      setSortKey(key);
      setSortAsc(true);
    }
    setPage(1);
  };

  return (
    <DefaultLayout>
      <div className="w-full p-5">
        <span className={title({size:"sm"})}>ONUs&nbsp;</span>

        <div className={`${showTable ? "flex" : "hidden"} mt-3 mb-2 gap-2 items-center`}>
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

          <span className="text-sm text-gray-600 whitespace-nowrap">
            {search.trim() ? `${filtered.length} / ${rows.length} ONUs` : `${rows.length} ONUs`}
          </span>
        </div>

        {loading && <div className="text-sm text-gray-600 mt-2">Loading...</div>}
        {error && <div className="text-sm text-red-600 mt-2">Error: {error}</div>}

        {/* 表格初始隐藏，收到 response 后显示 */}
        <div className={`${showTable ? "block" : "hidden"} w-full mt-4 overflow-auto`}>
          <div className="min-w-[800px]">
            <Table aria-label="ONUs table" style={{ width: "100%" }}
              bottomContent={
                // One page needs no pager, and most logs carry a handful of ONUs.
                totalPages > 1 ? (
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
                ) : null
              }>
              <TableHeader columns={columns}>
                {(column) => (
                  <TableColumn key={String(column.key)}>
                    <div
                      className={column.key !== "action" ? "inline-flex gap-2 items-center cursor-pointer" : "inline-flex gap-2 items-center"}
                      onClick={() => column.key !== "action" && toggleSort(String(column.key))}
                    >
                      <span className="text-sm font-medium">{column.label}</span>
                      {sortKey === column.key ? (sortAsc ? " ▲" : " ▼") : null}
                    </div>
                  </TableColumn>
                )}
              </TableHeader>

              <TableBody items={pageItems} emptyContent="No ONU matches the search.">
                {(item: any) => (
                  <TableRow key={item.onuName ?? item.key}>
                    {(columnKey: Key) => {
                      if (columnKey === "action") {
                        const onuName = item.onuName ?? "";
                        const encOnu = encodeURIComponent(onuName);
                        const diagramUrl = `/omcianalyzer/diagram?requestKey=${encodeURIComponent(requestKey)}&onuName=${encOnu}`;
                        const omciUrl = `/omcianalyzer/omci?requestKey=${encodeURIComponent(requestKey)}&onuName=${encOnu}`;

                        return (
                          <TableCell className="whitespace-nowrap">
                            <div className="flex gap-2 items-center">
                              <a href={diagramUrl} target="_blank" rel="noreferrer" className="inline-block">
                                <Button type="button" className={buttonStyles({ color: "primary", radius: "full", variant: "shadow" })}>
                                  Diagram
                                </Button>
                              </a>
                              <a href={omciUrl} target="_blank" rel="noreferrer" className="inline-block">
                                <Button type="button" className={buttonStyles({ color: "default", radius: "full", variant: "shadow" })}>
                                  OmciData
                                </Button>
                              </a>
                              <Button
                                type="button"
                                className={buttonStyles({ color: "secondary", radius: "full", variant: "shadow" })}
                                onPress={() => setYangOnu(onuName)}
                              >
                                LS Config
                              </Button>
                            </div>
                          </TableCell>
                        );
                      }

                      // 默认自动填充其他列
                      return (
                        <TableCell className="whitespace-nowrap text-sm">
                          {String(getKeyValue(item, String(columnKey)) ?? "")}
                        </TableCell>
                      );
                    }}
                  </TableRow>
                )}
              </TableBody>
            </Table>
          </div>
        </div>

        {/* 如果没有数据且非加载状态，显示提示 */}
        {!loading && !showTable && !error && requestKey && (
          <div className="mt-3 text-sm text-gray-600">No ONUs returned for requestKey.</div>
        )}

        <GenerateYangModal
          isOpen={yangOnu !== null}
          onClose={() => setYangOnu(null)}
          requestKey={requestKey}
          onuName={yangOnu ?? ""}
        />
      </div>
    </DefaultLayout>
  );
}
