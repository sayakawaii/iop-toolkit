import DefaultLayout from "@/layouts/default";
import {Button, Table, TableHeader, TableColumn, TableBody, TableRow, TableCell, getKeyValue} from "@heroui/react";
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
  }, [requestKey]);

  return (
    <DefaultLayout>
      <div className="w-full p-5">
        <span className={title({size:"sm"})}>ONUs&nbsp;</span>

        {loading && <div className="text-sm text-gray-600 mt-2">Loading...</div>}
        {error && <div className="text-sm text-red-600 mt-2">Error: {error}</div>}

        {/* 表格初始隐藏，收到 response 后显示 */}
        <div className={`${showTable ? "block" : "hidden"} w-full mt-4 overflow-auto`}>
          <div className="min-w-[800px]">
            <Table aria-label="ONUs table" style={{ width: "100%" }}>
              <TableHeader columns={columns}>
                {(column) => (
                  <TableColumn key={String(column.key)}>
                    <span className="text-sm font-medium">{column.label}</span>
                  </TableColumn>
                )}
              </TableHeader>

              <TableBody items={rows}>
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
