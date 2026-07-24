import { title } from "@/components/primitives";
import DefaultLayout from "@/layouts/default";
import {Form,Input,Button,Table, TableHeader, TableColumn, TableBody, TableRow, TableCell, getKeyValue,CircularProgress} from "@heroui/react";
import { button as buttonStyles } from "@heroui/theme";
import React, { useEffect, useRef, useCallback, Key } from "react";
import axios from "axios";
import { API_BASE } from "@/config/api";

export default function OmciAnalyzerPage() {
    const [submitted, setSubmitted] = React.useState<null | { [k: string]: FormDataEntryValue }>(null);
    // 新增：logs 数据和显示控制
    const [logs, setLogs] = React.useState<any[]>([]);
    const [showLogs, setShowLogs] = React.useState<boolean>(false);

    // polling ref to clear interval when needed
    const pollingRef = useRef<ReturnType<typeof setInterval> | null>(null);

    const onSubmit = (e: React.FormEvent<HTMLFormElement>) => {
      e.preventDefault();
      const data = Object.fromEntries(new FormData(e.currentTarget));
      setSubmitted(data);
    };

    const baseUrl = API_BASE + "/api/omcianalyzer";

    // merge progress updates into logs (by requestKey)
    const applyProgressUpdates = useCallback((updates: any[]) => {
      setLogs((prev) => {
        const map = new Map(prev.map((r: any) => [r.requestKey, r]));
        for (const u of updates) {
          const k = u.requestKey;
          if (!k) continue;
          const existing = map.get(k) || {};
          // merge shallowly: progress/status/etc overwrite existing
          map.set(k, { ...existing, ...u });
        }
        return Array.from(map.values());
      });
    }, []);

    // start polling given requestKeys; clears previous poll if any
    const startPolling = useCallback((requestKeys: string[]) => {
      // clear existing
      if (pollingRef.current) {
        globalThis.clearInterval(pollingRef.current);
        pollingRef.current = null;
      }

      const poll = async () => {
        try {
          const fd = new FormData();
          for (const k of requestKeys) fd.append("requestKey", k);
          const resp = await axios.post(`${baseUrl}/progress`, fd, {
            headers: { "Content-Type": "multipart/form-data" },
          });
          const data = resp.data;
          if (Array.isArray(data)) {
            applyProgressUpdates(data);

            // 如果任一记录返回 error，则停止轮询（错误已被记录到 logs）
            const anyError = data.some((d: any) => String(d.status) === "error");
            if (anyError && pollingRef.current) {
              globalThis.clearInterval(pollingRef.current);
              pollingRef.current = null;
              return;
            }

            // 仍在处理的条件：status 为 processing 或 progress 未到 100
            const stillRunning = data.some((d: any) => {
              const s = String(d.status ?? "").toLowerCase();
              if (s === "processing") return true;
              if (s === "success") return false;
              const p = Number(d.progress);
              return Number.isNaN(p) ? true : p !== 100;
            });

            if (!stillRunning && pollingRef.current) {
              globalThis.clearInterval(pollingRef.current);
              pollingRef.current = null;
            }
          }
        } catch (err) {
          console.error("Progress poll error:", err);
          // keep retrying; could add retry/backoff logic here
        }
      };

      // immediate poll then interval
      poll();
      pollingRef.current = globalThis.setInterval(poll, 1000);
      return () => {
        if (pollingRef.current) {
          globalThis.clearInterval(pollingRef.current);
          pollingRef.current = null;
        }
      };
    }, [applyProgressUpdates]);

    // History 按钮处理：发送空 POST 到 `${baseUrl}/history`，不启动轮询
    const handleHistoryClick = async () => {
      try {
        const formData = new FormData(); // empty POST
        const response = await axios.post(`${baseUrl}/history`, formData, {
          headers: { "Content-Type": "multipart/form-data" },
        });
        // backend 返回格式与 /request 相同
        let returned: any[] = [];
        if (Array.isArray(response.data)) returned = response.data;
        else if (response.data) returned = [response.data];

        if (returned.length) {
          setLogs((prev) => {
            const map = new Map(prev.map((r: any) => [r.requestKey, r]));
            for (const rec of returned) {
              if (!rec.requestKey) continue;
              map.set(rec.requestKey, { ...(map.get(rec.requestKey) || {}), ...rec });
            }
            return Array.from(map.values());
          });
          setShowLogs(true);
        } else {
          if (response.data && typeof response.data === "object") {
            setLogs([response.data]);
            setShowLogs(true);
          }
        }
      } catch (err) {
        console.error("History fetch error:", err);
      }
    };

    // upload request and kick off polling for returned requestKey(s)
    useEffect(() => {
        if (!submitted) return;

        let cleanup: (() => void) | undefined;
        const fetchData = async () => {
            try {
                const formData = new FormData();
                for (const key in submitted) {
                    const value = submitted[key];
                    if (value instanceof File) {
                        formData.append(key, value);
                    } else {
                        formData.append(key, value.toString());
                    }
                }

                const response = await axios.post(`${baseUrl}/request`, formData, {
                    headers: {
                        'Content-Type': 'multipart/form-data',
                    },
                });
                console.log('Response data:', response.data);

                // backend may return array of records or single record
                let returned: any[] = [];
                if (Array.isArray(response.data)) returned = response.data;
                else if (response.data) returned = [response.data];

                if (returned.length) {
                  // initialize logs with returned records (keep fields from response)
                  setLogs((prev) => {
                    const map = new Map(prev.map((r: any) => [r.requestKey, r]));
                    for (const rec of returned) {
                      if (!rec.requestKey) continue;
                      map.set(rec.requestKey, { ...(map.get(rec.requestKey) || {}), ...rec });
                    }
                    return Array.from(map.values());
                  });
                  setShowLogs(true);

                  const requestKeys = returned.map((r) => r.requestKey).filter(Boolean);
                  if (requestKeys.length) {
                    cleanup = startPolling(requestKeys);
                  }
                } else {
                  // no returned items
                  if (response.data && typeof response.data === "object") {
                    setLogs([response.data]);
                    setShowLogs(true);
                  }
                }
            } catch (error) {
                console.error('Error submitting form:', error);
            }
        };

        fetchData();

        return () => {
          if (cleanup) cleanup();
        };
    }, [submitted, startPolling]);

    // cleanup polling on unmount
    useEffect(() => {
      return () => {
        if (pollingRef.current) {
          globalThis.clearInterval(pollingRef.current);
          pollingRef.current = null;
        }
      };
    }, []);

  // 排序状态：支持按 logName / uploadTime 排序
  const [sortKey, setSortKey] = React.useState<string | null>(null);
  const [sortAsc, setSortAsc] = React.useState<boolean>(true);

  const toggleSort = (key: string) => {
    if (sortKey === key) {
      setSortAsc((s) => !s);
    } else {
      setSortKey(key);
      setSortAsc(true);
    }
  };

  // 根据 sortKey/sortAsc 对 logs 做稳定排序，返回新的数组（不修改原 logs）
  const sortedLogs = React.useMemo(() => {
    if (!sortKey) return logs.slice();
    const arr = logs.slice();
    arr.sort((a: any, b: any) => {
      const va = a[sortKey];
      const vb = b[sortKey];
      // uploadTime 尝试按时间比较；否则按字符串/数字比较
      if (sortKey === 'uploadTime') {
        const da = va ? Date.parse(String(va)) : Number.NaN;
        const db = vb ? Date.parse(String(vb)) : Number.NaN;
        if (!Number.isNaN(da) && !Number.isNaN(db)) return sortAsc ? da - db : db - da;
        // fallback numeric
        const na = Number(va || 0);
        const nb = Number(vb || 0);
        if (!Number.isNaN(na) && !Number.isNaN(nb)) return sortAsc ? na - nb : nb - na;
        // final fallback string
        const sa = String(va ?? '');
        const sb = String(vb ?? '');
        return sortAsc ? sa.localeCompare(sb) : sb.localeCompare(sa);
      }

      // default: string compare
      const sa = String(va ?? '');
      const sb = String(vb ?? '');
      return sortAsc ? sa.localeCompare(sb) : sb.localeCompare(sa);
    });
    return arr;
  }, [logs, sortKey, sortAsc]);

  // 定义用于自动渲染的列（前三列自动填充）
  const columns = [
    { key: 'logName', label: 'Name' },
    { key: 'uploadTime', label: 'UploadTime' },
    { key: 'progress', label: 'Progress' },
    { key: 'status', label: 'Status' },     // 新增 status 列
    { key: 'errorMsg', label: 'Error' },    // 新增 errorMsg 列
    { key: 'action', label: 'Action' } // 用于手动渲染 action 列
  ];

  return (
    <DefaultLayout>
      <section className="flex flex-col items-center justify-center gap-4 py-8 md:py-10">
        <div className="inline-block max-w-lg text-center justify-center">
          <span className={title({size:"sm"})}>OMCI Analyzer&nbsp;</span>
        </div>
        <Form className="w-full max-w-xs" onSubmit={onSubmit}>
            <div className="w-full mt-8 justify-center">
                <Input name="omcianalyzerFile" className="justify-center" color="primary" type="file" multiple />
            </div>
            <div className="flex gap-10 justify-center">
                <Button type="submit" className="bg-linear-to-tr from-pink-500 to-yellow-500 text-white shadow-lg" radius="full">
                    Start
                </Button>
                <Button type="reset" className={buttonStyles({
                    color: "default",
                    radius: "full",
                    variant: "shadow",
                    })}>
                    Reset
                </Button>
                <Button type="button" onPress={handleHistoryClick} className={buttonStyles({
                    color: "default",
                    radius: "full",
                    variant: "shadow",
                    })}>
                    History
                </Button>
            </div>
        </Form>
      </section>

            {/* 新增：heroui Table，初始隐藏，收到 response 时显示（全宽且水平可滚动） */}
            <div
                style={{
                    width: '100%',
                    marginTop: 24,
                    display: showLogs ? 'block' : 'none',
                    overflowX: 'auto' // 允许横向滚动，避免列换行
                }}
            >
                {/* 保持表格最小宽度，避免内容换行 */}
                <div style={{ minWidth: 900 }}>
                    <Table aria-label="Logs table" style={{ width: '100%' }}>
                        <TableHeader columns={columns}>
                          {(column) => {
                            const key = String(column.key);
                            // 仅为 Name / UploadTime 添加排序交互
                            if (key === 'logName' || key === 'uploadTime') {
                              const indicator = sortKey === key ? (sortAsc ? ' ▲' : ' ▼') : '';
                              return (
                                <TableColumn key={key}>
                                  <button type="button" onClick={() => toggleSort(key)} style={{ cursor: 'pointer', background: 'transparent', border: 'none', padding: 0 }}>
                                    {column.label}{indicator}
                                  </button>
                                </TableColumn>
                              );
                            }
                            return <TableColumn key={key}>{column.label}</TableColumn>;
                          }}
                        </TableHeader>

                        <TableBody items={sortedLogs}>
                          {(item: any) => (
                            <TableRow key={item.requestKey ?? item.logName}>
                              {(columnKey: Key) => {
                                // 手动为 action 列渲染按钮，其它列自动填充
                                if (columnKey === 'action') {
                                  return (
                                    <TableCell style={{ whiteSpace: 'nowrap' }}>
                                      {Number(item.progress) === 100 ? (
                                          <a href={`/omcianalyzer/onus?requestKey=${item.requestKey}`} target="_blank" rel="noreferrer">
                                              <Button type="button" className={buttonStyles({color: "success",radius: "full",variant: "shadow",})}>Check Result</Button>
                                          </a>
                                      ) : null}
                                    </TableCell>
                                  );
                                }

                                // 对 progress 列使用进度条呈现，同时保持自动取值风格
                                if (columnKey === 'progress') {
                                  const p = Number(getKeyValue(item, columnKey) ?? 0);
                                  return (
                                    <TableCell>
                                        <CircularProgress
                                        aria-label="Loading..."
                                        color="success"
                                        showValueLabel={true}
                                        size="lg"
                                        value={p}
                                        />
                                    </TableCell>
                                  );
                                }

                                // status 列：显示带颜色的 badge（success=green, processing=orange, error=red）
                                if (columnKey === 'status') {
                                  const st = String(getKeyValue(item, 'status') ?? '').toLowerCase();
                                  const bg = st === 'success' ? '#16a34a' : st === 'processing' ? '#f59e0b' : st === 'error' ? '#dc2626' : '#6b7280';
                                  const label = st ? st.charAt(0).toUpperCase() + st.slice(1) : '';
                                  return (
                                    <TableCell>
                                      <span style={{ display: 'inline-block', padding: '4px 8px', borderRadius: 12, color: '#fff', background: bg, fontSize: 12 }}>
                                        {label}
                                      </span>
                                    </TableCell>
                                  );
                                }

                                // errorMsg 列：显示错误信息，允许换行/截断
                                if (columnKey === 'errorMsg') {
                                  return (
                                    <TableCell style={{ maxWidth: 400, whiteSpace: 'normal', wordBreak: 'break-word' }}>
                                      {String(getKeyValue(item, 'errorMsg') ?? '')}
                                    </TableCell>
                                  );
                                }

                                // 默认自动填充其他列
                                return (
                                  <TableCell style={{ whiteSpace: 'nowrap' }}>
                                    {String(getKeyValue(item, String(columnKey)) ?? '')}
                                  </TableCell>
                                );
                              }}
                            </TableRow>
                          )}
                        </TableBody>
                    </Table>
                </div>
            </div>
     </DefaultLayout>
   );
 }
