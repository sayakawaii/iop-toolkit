import React from "react";
import axios from "axios";
import DefaultLayout from "@/layouts/default";
import { title } from "@/components/primitives";
import { Button, Input, Table, TableHeader, TableColumn, TableBody, TableRow, TableCell, Pagination} from "@heroui/react";
import { button as buttonStyles } from "@heroui/theme";
import { API_BASE } from "@/config/api";

type LibraryRow = {
  ID?: number | string;
  Customer?: string;
  RcrNum?: string;
  Vendor?: string;
  OnuType?: string;
  PonType?: string;
  RGMode?: string;
  EquipID?: string;
  HwVersion?: string;
  SwVersion?: string;
  LSRelease?: string;
  Comments?: string;
  InputLog?: string; // used for download requestKey
};

const baseUrl = API_BASE;
const SEARCH_API = baseUrl + "/api/library/search";
const ADD_API = baseUrl + "/api/library/addRecord";
const UPDATE_API = baseUrl + "/api/library/updateRecord";
const DELETE_API = baseUrl + "/api/library/deleteRecord";
const CUSTOMERS_API = baseUrl + "/api/library/customers";
const VENDORS_API = baseUrl + "/api/library/vendors";
const ADD_CUSTOMER_API = baseUrl + "/api/library/addCustomer";
const ADD_VENDOR_API = baseUrl + "/api/library/addVendor";
const DOWNLOAD_BASE = baseUrl + "/api/library/download?requestKey=";

export default function LibraryPage() {
  const [rows, setRows] = React.useState<LibraryRow[]>([]);
  const [loading, setLoading] = React.useState(false);
  const [searchTerm, setSearchTerm] = React.useState("");
  const [error, setError] = React.useState<string | null>(null);

  // insert / edit modal state
  const [showInsert, setShowInsert] = React.useState(false);
  const [insertLoading, setInsertLoading] = React.useState(false);
  const [editingId, setEditingId] = React.useState<string | number | null>(null);
  const [formValues, setFormValues] = React.useState<Record<string, any>>({
    Customer: "",
    RcrNum: "",
    Vendor: "",
    OnuType: "",
    PonType: "",
    RGMode: "",
    EquipID: "",
    HwVersion: "",
    SwVersion: "",
    LSRelease: "",
    Comments: "",
    // InputLog file handled separately
  });
  const fileInputRef = React.useRef<HTMLInputElement | null>(null);

  // customers / vendors lists
  const [customers, setCustomers] = React.useState<string[]>([]);
  const [vendors, setVendors] = React.useState<string[]>([]);
  const [listsLoading, setListsLoading] = React.useState(false);

  // insert-customer/vendor modals
  const [showAddCustomer, setShowAddCustomer] = React.useState(false);
  const [showAddVendor, setShowAddVendor] = React.useState(false);
  const [newCustomerName, setNewCustomerName] = React.useState("");
  const [newVendorName, setNewVendorName] = React.useState("");
  const [addEntityLoading, setAddEntityLoading] = React.useState(false);

  // delete confirm
  const [deleteTarget, setDeleteTarget] = React.useState<LibraryRow | null>(null);
  const [deleteLoading, setDeleteLoading] = React.useState(false);

  // validation state
  const [errors, setErrors] = React.useState<Record<string, string>>({});

  // stable key generator
  const rowKey = (r: LibraryRow) => String(r.ID ?? r.RcrNum ?? r.InputLog ?? Math.random());
  const [page, setPage] = React.useState(1);
  const rowsPerPage = 10;

  const pages = Math.ceil(rows.length / rowsPerPage);

  const items = React.useMemo(() => {
    const start = (page - 1) * rowsPerPage;
    const end = start + rowsPerPage;

    return rows.slice(start, end);
  }, [page, rows]);

  // icons (use currentColor so button styles apply)
  const DownloadIcon = ({ className = "w-4 h-4" }: { className?: string }) => (
    <svg className={className} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round" aria-hidden>
      <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4" />
      <polyline points="7 10 12 15 17 10" />
      <line x1="12" y1="15" x2="12" y2="3" />
    </svg>
  );

  // Analyzer (chart) icon — placed left of Edit
  const AnalyzerIcon = ({ className = "w-4 h-4" }: { className?: string }) => (
    <svg className={className} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round" aria-hidden>
      <path d="M3 3v18h18" />
      <rect x="6" y="10" width="2" height="7" rx="1" />
      <rect x="10.5" y="7" width="2" height="10" rx="1" />
      <rect x="15" y="4" width="2" height="13" rx="1" />
    </svg>
  );

  // stylized pen pointing left-down
  const EditIcon = ({ className = "w-4 h-4" }: { className?: string }) => (
    <svg className={className} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round" aria-hidden>
      <path d="M3 21v-3.6a2 2 0 0 1 .6-1.4L14 6.6a2 2 0 0 1 2.8 0l1.6 1.6a2 2 0 0 1 0 2.8L7 21.4A2 2 0 0 1 5.6 22H3z" />
      <path d="M14 6l4 4" />
    </svg>
  );

  const DeleteIcon = ({ className = "w-4 h-4" }: { className?: string }) => (
    <svg className={className} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2.2" strokeLinecap="round" strokeLinejoin="round" aria-hidden>
      <line x1="18" y1="6" x2="6" y2="18" />
      <line x1="6" y1="6" x2="18" y2="18" />
    </svg>
  );

  // download helper: fetch binary (blob) and trigger download with filename from headers if present
  const downloadFile = async (requestKey?: string | null) => {
    if (!requestKey) {
      setError("No file key");
      return;
    }
    try {
      setError(null);
      const url = `${DOWNLOAD_BASE}${encodeURIComponent(String(requestKey))}`;
      const resp = await axios.get(url, { responseType: "blob" });
      const blob = resp.data as Blob;
      let filename = "download";
      const cd = (resp.headers && (resp.headers["content-disposition"] || resp.headers["Content-Disposition"])) as string | undefined;
      if (cd) {
        const m = /filename\*?=(?:UTF-8'')?["']?([^;"']+)/i.exec(cd);
        if (m && m[1]) {
          try {
            filename = decodeURIComponent(m[1]);
          } catch {
            filename = m[1];
          }
        }
      }
      if (!filename || filename === "download") {
        const ext = blob.type ? blob.type.split("/")[1] : "";
        if (ext) filename = `download.${ext}`;
      }
      const blobUrl = window.URL.createObjectURL(blob);
      const a = document.createElement("a");
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

  // fetch customers/vendors (server returns []string)
  const fetchLists = React.useCallback(async () => {
    setListsLoading(true);
    try {
      const [cResp, vResp] = await Promise.allSettled([axios.get(CUSTOMERS_API), axios.get(VENDORS_API)]);
      if (cResp.status === "fulfilled" && Array.isArray(cResp.value.data)) {
        setCustomers(cResp.value.data.map(String));
      } else {
        setCustomers([]);
      }
      if (vResp.status === "fulfilled" && Array.isArray(vResp.value.data)) {
        setVendors(vResp.value.data.map(String));
      } else {
        setVendors([]);
      }
    } catch (err) {
      console.warn("fetch lists failed", err);
      setCustomers([]);
      setVendors([]);
    } finally {
      setListsLoading(false);
    }
  }, []);

  // fetch rows
  const fetchRows = React.useCallback(
    async (term = "") => {
      setLoading(true);
      setError(null);
      try {
        const fd = new FormData();
        fd.append("search", term);
        const resp = await axios.post(SEARCH_API, fd, {
          headers: { "Content-Type": "multipart/form-data" },
        });
        if (Array.isArray(resp.data)) setRows(resp.data);
        else if (resp.data) setRows([resp.data]);
        else setRows([]);
      } catch (err: any) {
        console.error("Search error", err);
        setError(err?.message ?? "Search failed");
        setRows([]);
      } finally {
        setLoading(false);
      }
    },
    []
  );

  // initial load
  React.useEffect(() => {
    fetchRows("");
    fetchLists();
  }, [fetchRows, fetchLists]);

  const doSearch = (e?: React.FormEvent) => {
    e?.preventDefault();
    fetchRows(searchTerm);
  };

  // open insert (or edit)
  const onOpenInsert = (row?: LibraryRow) => {
    setErrors({});
    if (row) {
      setEditingId(row.ID ?? null);
      setFormValues({
        Customer: row.Customer ?? "",
        RcrNum: row.RcrNum ?? "",
        Vendor: row.Vendor ?? "",
        OnuType: row.OnuType ?? "",
        PonType: row.PonType ?? "",
        RGMode: row.RGMode ?? "",
        EquipID: row.EquipID ?? "",
        HwVersion: row.HwVersion ?? "",
        SwVersion: row.SwVersion ?? "",
        LSRelease: row.LSRelease ?? "",
        Comments: row.Comments ?? "",
      });
      if (fileInputRef.current) fileInputRef.current.value = "";
    } else {
      setEditingId(null);
      setFormValues({
        Customer: "",
        RcrNum: "",
        Vendor: "",
        OnuType: "",
        PonType: "",
        RGMode: "",
        EquipID: "",
        HwVersion: "",
        SwVersion: "",
        LSRelease: "",
        Comments: "",
      });
      if (fileInputRef.current) fileInputRef.current.value = "";
    }
    setShowInsert(true);
  };

  // client-side validation for required fields
  const validateForm = (): boolean => {
    const req = ["Customer", "Vendor", "RcrNum", "OnuType", "PonType", "RGMode"];
    const next: Record<string, string> = {};
    for (const k of req) {
      if (!formValues[k] || String(formValues[k]).trim() === "") next[k] = "Required";
    }
    setErrors(next);
    return Object.keys(next).length === 0;
  };

  const submitInsert = async (e?: React.FormEvent) => {
    e?.preventDefault();
    if (!validateForm()) return;
    setInsertLoading(true);
    setError(null);
    try {
      const form = new FormData();
      for (const k of Object.keys(formValues)) form.append(k, formValues[k] ?? "");
      if (editingId) form.append("ID", String(editingId));
      const fileEl = fileInputRef.current;
      if (fileEl && fileEl.files && fileEl.files[0]) {
        form.append("file-input-log", fileEl.files[0]);
      }
      const api = editingId ? UPDATE_API : ADD_API;
      await axios.post(api, form, { headers: { "Content-Type": "multipart/form-data" } });
      setShowInsert(false);
      setEditingId(null);
      fetchRows(searchTerm);
      fetchLists();
    } catch (err: any) {
      console.error("Save failed", err);
      setError(err?.message ?? "Save failed");
    } finally {
      setInsertLoading(false);
    }
  };

  // delete flow
  const onRequestDelete = (r: LibraryRow) => setDeleteTarget(r);
  const confirmDelete = async () => {
    if (!deleteTarget) return;
    setDeleteLoading(true);
    setError(null);
    try {
      const fd = new FormData();
      fd.append("ID", String(deleteTarget.ID ?? ""));
      await axios.post(DELETE_API, fd, { headers: { "Content-Type": "multipart/form-data" } });
      setDeleteTarget(null);
      fetchRows(searchTerm);
    } catch (err: any) {
      console.error("Delete failed", err);
      setError(err?.message ?? "Delete failed");
    } finally {
      setDeleteLoading(false);
    }
  };

  // add customer/vendor
  const submitAddCustomer = async () => {
    if (!newCustomerName.trim()) return;
    setAddEntityLoading(true);
    try {
      const fd = new FormData();
      fd.append("name", newCustomerName);
      await axios.post(ADD_CUSTOMER_API, fd, { headers: { "Content-Type": "multipart/form-data" } });
      setNewCustomerName("");
      setShowAddCustomer(false);
      await fetchLists();
    } catch (err: any) {
      console.error("Add customer failed", err);
      setError(err?.message ?? "Add customer failed");
    } finally {
      setAddEntityLoading(false);
    }
  };
  const submitAddVendor = async () => {
    if (!newVendorName.trim()) return;
    setAddEntityLoading(true);
    try {
      const fd = new FormData();
      fd.append("name", newVendorName);
      await axios.post(ADD_VENDOR_API, fd, { headers: { "Content-Type": "multipart/form-data" } });
      setNewVendorName("");
      setShowAddVendor(false);
      await fetchLists();
    } catch (err: any) {
      console.error("Add vendor failed", err);
      setError(err?.message ?? "Add vendor failed");
    } finally {
      setAddEntityLoading(false);
    }
  };

  // define columns with responsive classNames: show more on larger screens
  const columns = [
    { key: "ID", label: "ID", className: "whitespace-nowrap" },
    { key: "Customer", label: "Customer", className: "whitespace-nowrap" },
    { key: "RcrNum", label: "RCR", className: "whitespace-nowrap" },
    { key: "Vendor", label: "Vendor", className: "hidden md:table-cell whitespace-nowrap" },
    { key: "OnuType", label: "ONU Type", className: "hidden md:table-cell whitespace-nowrap" },
    { key: "PonType", label: "PON Type", className: "hidden lg:table-cell whitespace-nowrap" },
    { key: "RGMode", label: "RG Mode", className: "hidden lg:table-cell whitespace-nowrap" },
    { key: "EquipID", label: "EquipID", className: "hidden xl:table-cell whitespace-nowrap" },
    { key: "HwVersion", label: "HwVersion", className: "hidden xl:table-cell whitespace-nowrap" },
    { key: "SwVersion", label: "SwVersion", className: "hidden xl:table-cell whitespace-nowrap" },
    { key: "LSRelease", label: "LSRelease", className: "hidden xl:table-cell whitespace-nowrap" },
    { key: "Comments", label: "Comments", className: "hidden 2xl:table-cell" },
    { key: "InputLog", label: "File", className: "whitespace-nowrap" },
    { key: "Actions", label: "Actions", className: "whitespace-nowrap" },
  ];

  return (
    <DefaultLayout fullWidth>
      <section className="flex flex-col gap-6 py-8 md:py-10">
        <header className="flex items-center justify-between">
          <span className={title({ size: "sm" })}>Library&nbsp;</span>
          <div className="flex items-center gap-2">
            <form onSubmit={doSearch} className="flex items-center gap-2">
              <input
                aria-label="Search library"
                placeholder="Search library..."
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
                className="px-3 py-2 border rounded bg-white dark:bg-gray-800"
              />
              <Button type="submit" className={buttonStyles({ color: "default", radius: "full" })}>
                Search
              </Button>
            </form>

            <Button onPress={() => onOpenInsert()} className={buttonStyles({ color: "default", radius: "full" })}>
              Insert Record
            </Button>
          </div>
        </header>

        {error && <div className="text-sm text-red-500">{error}</div>}

        <div className="overflow-auto rounded border">
          <div className="min-w-full">
            <Table
              aria-label="Library table" className="w-full" color="primary" selectionMode="single"
              bottomContent={
                <div className="flex w-full justify-center">
                <Pagination
                    isCompact
                    showControls
                    showShadow
                    color="secondary"
                    page={page}
                    total={pages}
                    onChange={(page) => setPage(page)}
                />
                </div>
              }>
              <TableHeader columns={columns}>
                {(column) => (
                  <TableColumn key={String(column.key)}>
                    <div className={column.className ?? ""}>{column.label}</div>
                  </TableColumn>
                )}
              </TableHeader>

              <TableBody emptyContent={"No data to display."} items={items}>
                {(item: LibraryRow) => {
                  return (
                    <TableRow key={rowKey(item)}>
                      {(columnKey: any) => {
                        const col = columns.find((c) => c.key === String(columnKey));
                        const cellClass = col?.className ?? "";
                        if (columnKey === "Actions") {
                          return (
                            <TableCell className={cellClass + " flex gap-2"}>
                              {/* Analyzer: opens omcianalyzer result view in a new tab using same requestKey as download */}
                              {item.InputLog ? (
                                <Button
                                    isIconOnly size="sm" radius="full" variant="light"
                                    onPress={() => {
                                    const key = item.InputLog ?? "";
                                    if (!key) return;
                                    const url = `/omcianalyzer/onus?requestKey=${encodeURIComponent(String(key))}`;
                                    window.open(url, "_blank", "noopener,noreferrer");
                                    }}
                                    title="Analyze"
                                    aria-label="Analyze"
                                >
                                <AnalyzerIcon />
                                </Button>
                              ) : null}
                               <Button
                                 isIconOnly  size="sm" radius="full" variant="light"
                                 onPress={() => onOpenInsert(item)}
                                 title="Edit"
                                 aria-label="Edit"
                               >
                                 <EditIcon />
                               </Button>
                               <Button
                                 isIconOnly  size="sm" radius="full" variant="light" color="danger"
                                 onPress={() => onRequestDelete(item)}
                                 title="Delete"
                                 aria-label="Delete"
                               >
                                 <DeleteIcon />
                               </Button>
                             </TableCell>
                           );
                         }

                        if (columnKey === "InputLog") {
                          return (
                            <TableCell className={cellClass}>
                              {item.InputLog ? (
                                <Button
                                  isIconOnly  size="sm" radius="full" variant="light"
                                  onPress={() => downloadFile(item.InputLog)}
                                  title="Download"
                                  aria-label="Download"
                                >
                                  <DownloadIcon />
                                </Button>
                              ) : (
                                <span className="text-xs text-gray-400">No file</span>
                              )}
                            </TableCell>
                          );
                        }

                        // default render field value
                        const val = (item as any)[String(columnKey)];
                        return <TableCell className="felx max-w-[60px] whitespace-normal break-words">{String(val ?? "")}</TableCell>;
                      }}
                    </TableRow>
                  );
                }}
              </TableBody>
            </Table>

            {/* Loading / empty states below table for accessibility */}
            {loading && <div className="p-3 text-sm text-gray-600">Loading...</div>}
            {!loading && rows.length === 0 && <div className="p-3 text-sm text-gray-600">No records</div>}
          </div>
        </div>

        {/* Insert / Edit Modal */}
        {showInsert && (
          <div className="fixed inset-0 z-50 flex items-start md:items-center justify-center p-4">
            <div
              className="absolute inset-0 bg-black/40"
              onClick={() => {
                setShowInsert(false);
                setEditingId(null);
              }}
            />
            <div className="relative z-10 w-full max-w-3xl bg-white dark:bg-gray-800 rounded shadow-lg p-4">
              <div className="flex items-center justify-between mb-4">
                <h3 className="text-lg font-medium">{editingId ? "Edit Record" : "Insert Record"}</h3>
                <button
                  onClick={() => {
                    setShowInsert(false);
                    setEditingId(null);
                  }}
                  className="px-2 py-1 rounded bg-gray-100 dark:bg-gray-700"
                >
                  Close
                </button>
              </div>

              <form onSubmit={submitInsert} className="grid grid-cols-1 md:grid-cols-2 gap-3">
                {/* Customer + add */}
                <label className="flex flex-col">
                  <div className="flex items-center justify-between">
                    <span className="text-xs text-gray-600">Customer <span className="text-red-500">*</span></span>
                    <div className="flex gap-2">
                      <button type="button" onClick={() => setShowAddCustomer(true)} className="text-xs px-2 py-0.5 bg-white dark:bg-gray-700 rounded">
                        + Add
                      </button>
                    </div>
                  </div>
                  <select
                    value={formValues.Customer}
                    onChange={(e) => setFormValues((s) => ({ ...s, Customer: e.target.value }))}
                    className="px-2 py-2 border rounded bg-white dark:bg-gray-700"
                    required
                    aria-required
                    disabled={listsLoading}
                  >
                    <option value="">{listsLoading ? "Loading..." : "Select customer"}</option>
                    {customers.map((c) => (
                      <option key={c} value={c}>
                        {c}
                      </option>
                    ))}
                  </select>
                  {errors.Customer && <span className="text-xs text-red-500">{errors.Customer}</span>}
                </label>

                {/* Vendor + add */}
                <label className="flex flex-col">
                  <div className="flex items-center justify-between">
                    <span className="text-xs text-gray-600">Vendor <span className="text-red-500">*</span></span>
                    <div className="flex gap-2">
                      <button type="button" onClick={() => setShowAddVendor(true)} className="text-xs px-2 py-0.5 bg-white dark:bg-gray-700 rounded">
                        + Add
                      </button>
                    </div>
                  </div>
                  <select
                    value={formValues.Vendor}
                    onChange={(e) => setFormValues((s) => ({ ...s, Vendor: e.target.value }))}
                    className="px-2 py-2 border rounded bg-white dark:bg-gray-700"
                    required
                    aria-required
                    disabled={listsLoading}
                  >
                    <option value="">{listsLoading ? "Loading..." : "Select vendor"}</option>
                    {vendors.map((v) => (
                      <option key={v} value={v}>
                        {v}
                      </option>
                    ))}
                  </select>
                  {errors.Vendor && <span className="text-xs text-red-500">{errors.Vendor}</span>}
                </label>

                {/* PON Type (select) */}
                <label className="flex flex-col">
                  <span className="text-xs text-gray-600">PON Type <span className="text-red-500">*</span></span>
                  <select value={formValues.PonType} onChange={(e) => setFormValues((s) => ({ ...s, PonType: e.target.value }))} className="px-2 py-2 border rounded bg-white dark:bg-gray-700" required aria-required>
                    <option value="">Select PON</option>
                    <option value="GPON">GPON</option>
                    <option value="XGSPON">XGSPON</option>
                    <option value="25GPON">25GPON</option>
                    <option value="HSPON">HSPON</option>
                  </select>
                  {errors.PonType && <span className="text-xs text-red-500">{errors.PonType}</span>}
                </label>

                {/* RG Mode (select) */}
                <label className="flex flex-col">
                  <span className="text-xs text-gray-600">RG Mode <span className="text-red-500">*</span></span>
                  <select value={formValues.RGMode} onChange={(e) => setFormValues((s) => ({ ...s, RGMode: e.target.value }))} className="px-2 py-2 border rounded bg-white dark:bg-gray-700" required aria-required>
                    <option value="">Select RG Mode</option>
                    <option value="SFU">SFU</option>
                    <option value="HGU">HGU</option>
                    <option value="MDU">MDU</option>
                  </select>
                  {errors.RGMode && <span className="text-xs text-red-500">{errors.RGMode}</span>}
                </label>

                {/* RCR */}
                <label className="flex flex-col">
                  <span className="text-xs text-gray-600">RCR <span className="text-red-500">*</span></span>
                  <input
                    value={formValues.RcrNum}
                    onChange={(e) => setFormValues((s) => ({ ...s, RcrNum: e.target.value }))}
                    className="px-2 py-2 border rounded bg-white dark:bg-gray-700"
                    required
                    aria-required
                  />
                  {errors.RcrNum && <span className="text-xs text-red-500">{errors.RcrNum}</span>}
                </label>

                {/* ONU Type */}
                <label className="flex flex-col">
                  <span className="text-xs text-gray-600">ONU Type <span className="text-red-500">*</span></span>
                  <input
                    value={formValues.OnuType}
                    onChange={(e) => setFormValues((s) => ({ ...s, OnuType: e.target.value }))}
                    className="px-2 py-2 border rounded bg-white dark:bg-gray-700"
                    required
                    aria-required
                  />
                  {errors.OnuType && <span className="text-xs text-red-500">{errors.OnuType}</span>}
                </label>

                {/* HwVersion */}
                <label className="flex flex-col">
                  <span className="text-xs text-gray-600">HwVersion</span>
                  <input value={formValues.HwVersion} onChange={(e) => setFormValues((s) => ({ ...s, HwVersion: e.target.value }))} className="px-2 py-2 border rounded bg-white dark:bg-gray-700" />
                </label>

                {/* SwVersion */}
                <label className="flex flex-col">
                  <span className="text-xs text-gray-600">SwVersion</span>
                  <input value={formValues.SwVersion} onChange={(e) => setFormValues((s) => ({ ...s, SwVersion: e.target.value }))} className="px-2 py-2 border rounded bg-white dark:bg-gray-700" />
                </label>

                {/* EquipID */}
                <label className="flex flex-col">
                  <span className="text-xs text-gray-600">EquipID</span>
                  <input value={formValues.EquipID} onChange={(e) => setFormValues((s) => ({ ...s, EquipID: e.target.value }))} className="px-2 py-2 border rounded bg-white dark:bg-gray-700" />
                </label>

                {/* Comments */}
                <label className="flex flex-col md:col-span-2">
                  <span className="text-xs text-gray-600">Comments</span>
                  <input value={formValues.Comments} onChange={(e) => setFormValues((s) => ({ ...s, Comments: e.target.value }))} className="px-2 py-2 border rounded bg-white dark:bg-gray-700" />
                </label>

                {/* Attach Log File */}
                <label className="flex flex-col md:col-span-2">
                  <span className="text-xs text-gray-600">Attach Log File</span>
                  <Input ref={fileInputRef} type="file" name="file-input-log" className="mt-1" color="primary" />
                </label>

                <div className="md:col-span-2 flex justify-end gap-2">
                  <Button
                    onPress={() => {
                      setShowInsert(false);
                      setEditingId(null);
                    }}
                    className={buttonStyles({ color: "default" })}
                  >
                    Cancel
                  </Button>
                  <Button type="submit" className={buttonStyles({ color: "primary" })} disabled={insertLoading}>
                    {insertLoading ? "Saving..." : "Save"}
                  </Button>
                </div>
              </form>
            </div>
          </div>
        )}

        {/* Add Customer Modal */}
        {showAddCustomer && (
          <div className="fixed inset-0 z-50 flex items-center justify-center p-4">
            <div className="absolute inset-0 bg-black/40" onClick={() => setShowAddCustomer(false)} />
            <div className="relative z-10 w-full max-w-md bg-white dark:bg-gray-800 rounded shadow-lg p-4">
              <h4 className="mb-2">Add Customer</h4>
              <input className="w-full px-2 py-2 border rounded bg-white dark:bg-gray-700" value={newCustomerName} onChange={(e) => setNewCustomerName(e.target.value)} />
              <div className="flex justify-end gap-2 mt-3">
                <Button onPress={() => setShowAddCustomer(false)} className={buttonStyles({ color: "default" })}>
                  Cancel
                </Button>
                <Button onPress={submitAddCustomer} className={buttonStyles({ color: "primary" })} disabled={addEntityLoading}>
                  Add
                </Button>
              </div>
            </div>
          </div>
        )}

        {/* Add Vendor Modal */}
        {showAddVendor && (
          <div className="fixed inset-0 z-50 flex items-center justify-center p-4">
            <div className="absolute inset-0 bg-black/40" onClick={() => setShowAddVendor(false)} />
            <div className="relative z-10 w-full max-w-md bg-white dark:bg-gray-800 rounded shadow-lg p-4">
              <h4 className="mb-2">Add Vendor</h4>
              <input className="w-full px-2 py-2 border rounded bg-white dark:bg-gray-700" value={newVendorName} onChange={(e) => setNewVendorName(e.target.value)} />
              <div className="flex justify-end gap-2 mt-3">
                <Button onPress={() => setShowAddVendor(false)} className={buttonStyles({ color: "default" })}>
                  Cancel
                </Button>
                <Button onPress={submitAddVendor} className={buttonStyles({ color: "primary" })} disabled={addEntityLoading}>
                  Add
                </Button>
              </div>
            </div>
          </div>
        )}

        {/* Delete Confirm */}
        {deleteTarget && (
          <div className="fixed inset-0 z-50 flex items-center justify-center p-4">
            <div className="absolute inset-0 bg-black/40" onClick={() => setDeleteTarget(null)} />
            <div className="relative z-10 w-full max-w-md bg-white dark:bg-gray-800 rounded shadow-lg p-4">
              <h4 className="mb-2">Confirm Delete</h4>
              <p className="text-sm text-gray-600">Delete record {deleteTarget.RcrNum ?? deleteTarget.ID}? This action cannot be undone.</p>
              <div className="flex justify-end gap-2 mt-4">
                <Button onPress={() => setDeleteTarget(null)} className={buttonStyles({ color: "default" })}>
                  Cancel
                </Button>
                <Button onPress={confirmDelete} className={buttonStyles({ color: "danger" })} disabled={deleteLoading}>
                  {deleteLoading ? "Deleting..." : "Delete"}
                </Button>
              </div>
            </div>
          </div>
        )}
      </section>
    </DefaultLayout>
  );
}
