import DefaultLayout from "@/layouts/default";
import { useState, useEffect, useMemo, useCallback, useRef, Key } from "react";
import { useSearchParams } from "react-router-dom";
import axios from "axios";
import { Button, Input, Table, TableHeader, TableColumn, TableBody, TableRow, TableCell, Modal, ModalContent, ModalHeader, ModalBody, ModalFooter, CircularProgress, getKeyValue } from "@heroui/react";
import { API_BASE } from "@/config/api";

const baseUrl = API_BASE;
const LOGGER_GET_LOGS_API = baseUrl + "/api/collector/loggerGetLogs";
const LOGGER_POST_LOGS_TO_ANALYZER_API = baseUrl + "/api/omcianalyzer/minio";

type LogEntry = {
  app: string;
  minio: string;
  size: number;
  url: string;
};

const DownloadIcon = ({ className = "w-4 h-4" }: { className?: string }) => (
  <svg className={className} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round" aria-hidden>
    <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4" />
    <polyline points="7 10 12 15 17 10" />
    <line x1="12" y1="15" x2="12" y2="3" />
  </svg>
);

const AnalyzerIcon = ({ className = "w-4 h-4" }: { className?: string }) => (
  <svg className={className} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round" aria-hidden>
    <path d="M3 3v18h18" />
    <rect x="6" y="10" width="2" height="7" rx="1" />
    <rect x="10.5" y="7" width="2" height="10" rx="1" />
    <rect x="15" y="4" width="2" height="13" rx="1" />
  </svg>
);

export default function CollectorLogsPage() {
  const [searchParams] = useSearchParams();
  const requestId = searchParams.get('request_id');
  
  const [logs, setLogs] = useState<LogEntry[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [filterValue, setFilterValue] = useState("");
  const [page, setPage] = useState(1);
  const rowsPerPage = 10;

  // Modal states for analyzer
  const [showModal, setShowModal] = useState(false);
  const [modalLogs, setModalLogs] = useState<any[]>([]);
  const [modalSortKey, setModalSortKey] = useState<string | null>(null);
  const [modalSortAsc, setModalSortAsc] = useState<boolean>(true);
  const pollingRef = useRef<ReturnType<typeof setInterval> | null>(null);

  // Analyzer base URL (same as omcianalyzer)
  const analyzerBaseUrl = API_BASE + "/api/omcianalyzer";

  useEffect(() => {
    if (requestId) {
      fetchLogs();
    }
  }, [requestId]);

  const fetchLogs = async () => {
    setLoading(true);
    setError(null);
    try {
      const response = await axios.get(`${LOGGER_GET_LOGS_API}?request_id=${requestId}`);
      setLogs(response.data.logs || []);
    } catch (err: any) {
      console.error("Error fetching logs:", err);
      setError(err?.message ?? "Failed to fetch logs");
    } finally {
      setLoading(false);
    }
  };

  // Progress polling logic (same as omcianalyzer/index.tsx)
  const applyProgressUpdates = useCallback((updates: any[]) => {
    setModalLogs((prev) => {
      const map = new Map(prev.map((r: any) => [r.requestKey, r]));
      for (const u of updates) {
        const k = u.requestKey;
        if (!k) continue;
        const existing = map.get(k) || {};
        map.set(k, { ...existing, ...u });
      }
      return Array.from(map.values());
    });
  }, []);

  const startPolling = useCallback((requestKeys: string[]) => {
    if (pollingRef.current) {
      globalThis.clearInterval(pollingRef.current);
      pollingRef.current = null;
    }

    const poll = async () => {
      try {
        const fd = new FormData();
        for (const k of requestKeys) fd.append("requestKey", k);
        const resp = await axios.post(`${analyzerBaseUrl}/progress`, fd, {
          headers: { "Content-Type": "multipart/form-data" },
        });
        const data = resp.data;
        if (Array.isArray(data)) {
          applyProgressUpdates(data);

          const anyError = data.some((d: any) => String(d.status) === "error");
          if (anyError && pollingRef.current) {
            globalThis.clearInterval(pollingRef.current);
            pollingRef.current = null;
            return;
          }

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
      }
    };

    poll();
    pollingRef.current = globalThis.setInterval(poll, 1000);
    return () => {
      if (pollingRef.current) {
        globalThis.clearInterval(pollingRef.current);
        pollingRef.current = null;
      }
    };
  }, [applyProgressUpdates, analyzerBaseUrl]);

  // Cleanup polling on unmount
  useEffect(() => {
    return () => {
      if (pollingRef.current) {
        globalThis.clearInterval(pollingRef.current);
        pollingRef.current = null;
      }
    };
  }, []);

  const handleDownload = async (log: LogEntry) => {
    try {
      const response = await axios.get(log.url, { responseType: 'blob' });
      const blob = response.data as Blob;
      const filename = log.minio.split('/').pop() || 'download.log';
      const blobUrl = window.URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = blobUrl;
      a.download = filename;
      document.body.appendChild(a);
      a.click();
      a.remove();
      window.URL.revokeObjectURL(blobUrl);
    } catch (err: any) {
      console.error("Download failed", err);
      setError(err?.message ?? "Download failed");
    }
  };

  const handleAnalyze = async (log: LogEntry) => {
    // Check if app is onu_engine
    if (log.app !== 'onu_engine') {
      console.log('Analyze: Only onu_engine logs are supported');
      setError('Only onu_engine logs can be analyzed');
      return;
    }

    try {
      // Post log.minio to analyzer API (same logic as omcianalyzer/index.tsx line 148)
      const payload: {
      minio: string;
    } = { minio: log.minio };

      const response = await axios.post(LOGGER_POST_LOGS_TO_ANALYZER_API, payload);
      console.log('Analyzer response:', response.data);

      // Process response data (same as omcianalyzer/index.tsx)
      let returned: any[] = [];
      if (Array.isArray(response.data)) returned = response.data;
      else if (response.data) returned = [response.data];

      if (returned.length) {
        setModalLogs((prev) => {
          const map = new Map(prev.map((r: any) => [r.requestKey, r]));
          for (const rec of returned) {
            if (!rec.requestKey) continue;
            map.set(rec.requestKey, { ...(map.get(rec.requestKey) || {}), ...rec });
          }
          return Array.from(map.values());
        });
        setShowModal(true);

        const requestKeys = returned.map((r) => r.requestKey).filter(Boolean);
        if (requestKeys.length) {
          startPolling(requestKeys);
        }
      } else {
        if (response.data && typeof response.data === "object") {
          setModalLogs([response.data]);
          setShowModal(true);
        }
      }
    } catch (err: any) {
      console.error("Analyze failed", err);
      setError(err?.message ?? "Analysis failed");
    }
  };

  const hasSearchFilter = Boolean(filterValue);

  const filteredItems = useMemo(() => {
    let filtered = [...logs];
    if (hasSearchFilter) {
      filtered = filtered.filter((item) =>
        item.app.toLowerCase().includes(filterValue.toLowerCase()) ||
        item.minio.toLowerCase().includes(filterValue.toLowerCase())
      );
    }
    return filtered;
  }, [logs, filterValue]);

  const items = useMemo(() => {
    const start = (page - 1) * rowsPerPage;
    const end = start + rowsPerPage;
    return filteredItems.slice(start, end);
  }, [page, filteredItems]);

  const onSearchChange = useCallback((value: string) => {
    if (value) {
      setFilterValue(value);
      setPage(1);
    } else {
      setFilterValue("");
    }
  }, []);

  const onClear = useCallback(() => {
    setFilterValue("");
    setPage(1);
  }, []);

  // Modal sorting logic
  const toggleModalSort = (key: string) => {
    if (modalSortKey === key) {
      setModalSortAsc((s) => !s);
    } else {
      setModalSortKey(key);
      setModalSortAsc(true);
    }
  };

  const sortedModalLogs = useMemo(() => {
    if (!modalSortKey) return modalLogs.slice();
    const arr = modalLogs.slice();
    arr.sort((a: any, b: any) => {
      const va = a[modalSortKey];
      const vb = b[modalSortKey];
      if (modalSortKey === 'uploadTime') {
        const da = va ? Date.parse(String(va)) : Number.NaN;
        const db = vb ? Date.parse(String(vb)) : Number.NaN;
        if (!Number.isNaN(da) && !Number.isNaN(db)) return modalSortAsc ? da - db : db - da;
        const na = Number(va || 0);
        const nb = Number(vb || 0);
        if (!Number.isNaN(na) && !Number.isNaN(nb)) return modalSortAsc ? na - nb : nb - na;
        const sa = String(va ?? '');
        const sb = String(vb ?? '');
        return modalSortAsc ? sa.localeCompare(sb) : sb.localeCompare(sa);
      }
      const sa = String(va ?? '');
      const sb = String(vb ?? '');
      return modalSortAsc ? sa.localeCompare(sb) : sb.localeCompare(sa);
    });
    return arr;
  }, [modalLogs, modalSortKey, modalSortAsc]);

  const modalColumns = [
    { key: 'logName', label: 'Name' },
    { key: 'uploadTime', label: 'UploadTime' },
    { key: 'progress', label: 'Progress' },
    { key: 'status', label: 'Status' },
    { key: 'errorMsg', label: 'Error' },
    { key: 'action', label: 'Action' }
  ];

  const topContent = useMemo(() => {
    return (
      <div className="flex flex-col gap-4">
        <div className="flex justify-between gap-3 items-end">
          <Input
            isClearable
            className="w-full sm:max-w-[44%]"
            placeholder="Search by app or minio..."
            value={filterValue}
            onClear={onClear}
            onValueChange={onSearchChange}
          />
          <div className="flex gap-3">
            <span className="text-default-400 text-small">
              Total {filteredItems.length} logs
            </span>
          </div>
        </div>
      </div>
    );
  }, [filterValue, filteredItems.length, onSearchChange, onClear]);

  return (
    <DefaultLayout>
      <section className="flex flex-col gap-6 py-8 md:py-10">
        <div>
          <h1 className="text-2xl font-bold">Logger Logs</h1>
          {requestId && (
            <p className="text-small text-default-400 mt-1">Request ID: {requestId}</p>
          )}
        </div>
        
        {error && (
          <div className="bg-danger-50 text-danger p-4 rounded-lg">
            {error}
          </div>
        )}

        <Table
          aria-label="Logger logs table"
          topContent={topContent}
          topContentPlacement="outside"
        >
          <TableHeader>
            <TableColumn key="app">APP</TableColumn>
            <TableColumn key="minio">MINIO</TableColumn>
            <TableColumn key="size">SIZE (MB)</TableColumn>
            <TableColumn key="actions">ACTIONS</TableColumn>
          </TableHeader>
          <TableBody
            items={items}
            isLoading={loading}
            emptyContent={loading ? "Loading..." : "No logs found"}
          >
            {(item) => (
              <TableRow key={item.minio}>
                <TableCell>{item.app}</TableCell>
                <TableCell>{item.minio}</TableCell>
                <TableCell>{item.size.toFixed(2)}</TableCell>
                <TableCell>
                  <div className="flex gap-2">
                    <Button
                      isIconOnly
                      size="sm"
                      variant="light"
                      onPress={() => handleDownload(item)}
                      title="Download"
                      aria-label="Download"
                    >
                      <DownloadIcon />
                    </Button>
                    <Button
                      isIconOnly
                      size="sm"
                      variant="light"
                      onPress={() => handleAnalyze(item)}
                      title="Analyze"
                      aria-label="Analyze"
                    >
                      <AnalyzerIcon />
                    </Button>
                  </div>
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>

        {/* Modal for analyzer results */}
        <Modal 
          isOpen={showModal} 
          onOpenChange={setShowModal}
          size="5xl"
          scrollBehavior="inside"
        >
          <ModalContent>
            {(onClose) => (
              <>
                <ModalHeader className="flex flex-col gap-1">
                  Analyzer Results
                </ModalHeader>
                <ModalBody>
                  <div style={{ minWidth: 900 }}>
                    <Table aria-label="Analyzer results table" style={{ width: '100%' }}>
                      <TableHeader columns={modalColumns}>
                        {(column) => {
                          const key = String(column.key);
                          if (key === 'logName' || key === 'uploadTime') {
                            const indicator = modalSortKey === key ? (modalSortAsc ? ' ▲' : ' ▼') : '';
                            return (
                              <TableColumn key={key}>
                                <button 
                                  type="button" 
                                  onClick={() => toggleModalSort(key)} 
                                  style={{ cursor: 'pointer', background: 'transparent', border: 'none', padding: 0 }}
                                >
                                  {column.label}{indicator}
                                </button>
                              </TableColumn>
                            );
                          }
                          return <TableColumn key={key}>{column.label}</TableColumn>;
                        }}
                      </TableHeader>

                      <TableBody items={sortedModalLogs}>
                        {(item: any) => (
                          <TableRow key={item.requestKey ?? item.logName}>
                            {(columnKey: Key) => {
                              if (columnKey === 'action') {
                                return (
                                  <TableCell style={{ whiteSpace: 'nowrap' }}>
                                    {Number(item.progress) === 100 ? (
                                      <a href={`/omcianalyzer/onus?requestKey=${item.requestKey}`} target="_blank" rel="noreferrer">
                                        <Button 
                                          type="button" 
                                          color="success"
                                          radius="full"
                                          variant="shadow"
                                        >
                                          Check Result
                                        </Button>
                                      </a>
                                    ) : null}
                                  </TableCell>
                                );
                              }

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

                              if (columnKey === 'status') {
                                const st = String(getKeyValue(item, 'status') ?? '').toLowerCase();
                                const bg = st === 'success' ? '#16a34a' : st === 'processing' ? '#f59e0b' : st === 'error' ? '#dc2626' : '#6b7280';
                                const label = st ? st.charAt(0).toUpperCase() + st.slice(1) : '';
                                return (
                                  <TableCell>
                                    <span style={{ 
                                      display: 'inline-block', 
                                      padding: '4px 8px', 
                                      borderRadius: 12, 
                                      color: '#fff', 
                                      background: bg, 
                                      fontSize: 12 
                                    }}>
                                      {label}
                                    </span>
                                  </TableCell>
                                );
                              }

                              if (columnKey === 'errorMsg') {
                                return (
                                  <TableCell style={{ maxWidth: 400, whiteSpace: 'normal', wordBreak: 'break-word' }}>
                                    {String(getKeyValue(item, 'errorMsg') ?? '')}
                                  </TableCell>
                                );
                              }

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
                </ModalBody>
                <ModalFooter>
                  <Button color="danger" variant="light" onPress={onClose}>
                    Close
                  </Button>
                </ModalFooter>
              </>
            )}
          </ModalContent>
        </Modal>
      </section>
    </DefaultLayout>
  );
}
