import { useState, useMemo, useCallback, useEffect } from "react";
import axios from "axios";
import {
  Button,
  Input,
  Table,
  TableHeader,
  TableColumn,
  TableBody,
  TableRow,
  TableCell,
  DropdownTrigger,
  Dropdown,
  DropdownMenu,
  DropdownItem,
  Chip,
  Pagination,
  Spinner,
} from "@heroui/react";

import DefaultLayout from "@/layouts/default";
import LoggerModal from "@/pages/collector/LoggerModal";
import { API_BASE, wsBaseUrl } from "@/config/api";

const baseUrl = API_BASE;
const CONNECT_API = baseUrl + "/api/collector/connect";
const QUERY_API = baseUrl + "/api/collector/query";

const VerticalDotsIcon = ({
  size = 24,
  width,
  height,
  ...props
}: {
  size?: number;
  width?: number;
  height?: number;
  [key: string]: any;
}) => {
  return (
    <svg
      aria-hidden="true"
      fill="none"
      focusable="false"
      height={size || height}
      role="presentation"
      viewBox="0 0 24 24"
      width={size || width}
      {...props}
    >
      <path
        d="M12 10c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2zm0-6c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2zm0 12c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"
        fill="currentColor"
      />
    </svg>
  );
};

const SearchIcon = (props: { [key: string]: any }) => {
  return (
    <svg
      aria-hidden="true"
      fill="none"
      focusable="false"
      height="1em"
      role="presentation"
      viewBox="0 0 24 24"
      width="1em"
      {...props}
    >
      <path
        d="M11.5 21C16.7467 21 21 16.7467 21 11.5C21 6.25329 16.7467 2 11.5 2C6.25329 2 2 6.25329 2 11.5C2 16.7467 6.25329 21 11.5 21Z"
        stroke="currentColor"
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth="2"
      />
      <path
        d="M22 22L20 20"
        stroke="currentColor"
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth="2"
      />
    </svg>
  );
};

const ChevronDownIcon = ({
  strokeWidth = 1.5,
  ...otherProps
}: {
  strokeWidth?: number;
  [key: string]: any;
}) => {
  return (
    <svg
      aria-hidden="true"
      fill="none"
      focusable="false"
      height="1em"
      role="presentation"
      viewBox="0 0 24 24"
      width="1em"
      {...otherProps}
    >
      <path
        d="m19.92 8.95-6.52 6.52c-.77.77-2.03.77-2.8 0L4.08 8.95"
        stroke="currentColor"
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeMiterlimit={10}
        strokeWidth={strokeWidth}
      />
    </svg>
  );
};

const statusColorMap: { [key: string]: string } = {
  active: "success",
  paused: "danger",
  vacation: "warning",
};

const MatrixRain = () => {
  useEffect(() => {
    const canvas = document.getElementById(
      "matrix-canvas",
    ) as HTMLCanvasElement;

    if (!canvas) return;
    const ctx = canvas.getContext("2d");

    if (!ctx) return;

    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;

    const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#$%^&*()";
    const fontSize = 16;
    const columns = Math.floor(canvas.width / fontSize);
    const drops: number[] = [];

    for (let i = 0; i < columns; i++) {
      drops[i] = Math.floor((Math.random() * canvas.height) / fontSize);
    }

    const draw = () => {
      ctx.fillStyle = "rgba(0, 0, 0, 0.04)";
      ctx.fillRect(0, 0, canvas.width, canvas.height);

      ctx.fillStyle = "#0F0";
      ctx.font = fontSize + "px monospace";

      for (let i = 0; i < drops.length; i++) {
        const text = chars[Math.floor(Math.random() * chars.length)];

        ctx.fillText(text, i * fontSize, drops[i] * fontSize);

        if (drops[i] * fontSize > canvas.height && Math.random() > 0.975) {
          drops[i] = 0;
        }
        drops[i]++;
      }
    };

    const interval = setInterval(draw, 33);

    return () => clearInterval(interval);
  }, []);

  return <canvas className="absolute inset-0 z-0" id="matrix-canvas" />;
};

const TreeNode = ({
  node,
  onSelect,
}: {
  node: any;
  onSelect: (node: any) => void;
}) => {
  const visibleChildren = node.children
    ? node.children.filter((child: any) => child.type !== "v-ani")
    : [];

  return (
    <div>
      <Button variant="light" onPress={() => onSelect(node)}>
        {node.title}
      </Button>
      {visibleChildren.length > 0 && (
        <div className="ml-4">
          {visibleChildren.map((child: any) => (
            <TreeNode key={child.key} node={child} onSelect={onSelect} />
          ))}
        </div>
      )}
    </div>
  );
};

export default function CollectorPage() {
  const [oamIp, setOamIp] = useState("");
  const [username, setUsername] = useState("admin");
  const [password, setPassword] = useState("admin");
  const [data, setData] = useState<any>(null);
  const [tableData, setTableData] = useState<any[]>([]);
  const [loading, setLoading] = useState(false);
  const [filterValue, setFilterValue] = useState("");
  const [selectedKeys, setSelectedKeys] = useState<any>(new Set([]));
  const [visibleColumns, setVisibleColumns] = useState<any>(new Set());
  const [statusFilter, setStatusFilter] = useState<any>("all");
  const [rowsPerPage, setRowsPerPage] = useState(5);
  const [sortDescriptor, setSortDescriptor] = useState<{
    column: string;
    direction: "ascending" | "descending";
  }>({
    column: "name",
    direction: "ascending",
  });
  const [page, setPage] = useState(1);
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [showLoggerModal, setShowLoggerModal] = useState(false);
  const [currentItem, setCurrentItem] = useState<any>(null);
  const [ws, setWs] = useState<WebSocket | null>(null);

  const setupWebSocket = useCallback((requestId: string) => {
    const wsUrl = wsBaseUrl() + `/api/collector/ws?request_id=${requestId}`;
    const websocket = new WebSocket(wsUrl);

    websocket.onmessage = async (event) => {
      try {
        const message = JSON.parse(event.data);
        if (message.type === "update") {
          const response = await axios.get(`${QUERY_API}?request_id=${requestId}`);
          if (response.data.status === "Done" && response.data.result) {
            const payload = response.data.result;
            let normalized: any = payload?.oltTopology || payload?.ntinfo || payload || [];
            normalized = Array.isArray(normalized) ? normalized : [normalized];
            
            // Only update if we have valid data
            if (normalized && (Array.isArray(normalized) ? normalized.length > 0 : Object.keys(normalized).length > 0)) {
              setData(normalized);
              console.log("Data updated from WebSocket");
            }
          }
        }
      } catch (error) {
        console.error("Error refreshing data:", error);
        // Keep existing data on error
      }
    };

    websocket.onerror = (error) => {
      console.error("WebSocket error:", error);
      // Keep existing data on error
    };

    websocket.onclose = () => {
      console.log("WebSocket connection closed");
    };

    setWs(websocket);
  }, []);

  useEffect(() => {
    return () => {
      if (ws) {
        ws.close();
      }
    };
  }, [ws]);

  const pollQuery = useCallback(
    async (requestId: string, startTime: number) => {
      try {
        const response = await axios.get(
          `${QUERY_API}?request_id=${requestId}`,
        );

        if (response.data.status === "Done") {
          const payload = response.data.result;
          let normalized: any = payload?.oltTopology || payload?.ntinfo || payload || [];
          // Ensure the result is always an array
          normalized = Array.isArray(normalized) ? normalized : [normalized];

          setData(normalized);
          setLoading(false);

          // Check if data is from cache and being updated
          if (response.data.from_cache && response.data.updating) {
            setupWebSocket(requestId);
          }
        } else if (response.data.status === "Running") {
          // Check if data is from cache and being updated
          if (response.data.from_cache && response.data.updating) {
            // Setup WebSocket to wait for update while continuing to poll
            setupWebSocket(requestId);
          }
          
          if (Date.now() - startTime < 30000) {
            setTimeout(() => pollQuery(requestId, startTime), 1000);
          } else {
            console.error("Polling timeout");
            setLoading(false);
          }
        } else if (response.data.status === "Failed") {
          setErrorMessage(response.data.error || "Unknown error occurred");
          setLoading(false);
        } else {
          console.error("Unknown status:", response.data.status);
          setLoading(false);
        }
      } catch (error) {
        console.error("Error polling:", error);
        setLoading(false);
      }
    },
    [setupWebSocket],
  );

  const handleConnect = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    try {
      const response = await axios.post(CONNECT_API, {
        oamIp,
        username,
        password,
      });
      const requestId = response.data.request_id;

      sessionStorage.setItem("request_id", requestId);
      const startTime = Date.now();

      pollQuery(requestId, startTime);
    } catch (error) {
      console.error("Error connecting:", error);
      setLoading(false);
    }
  };

  const buildTree = (items: any[]): any[] => {
    return items.map((item) => {
      // Backend may use different property names for nested arrays: child, ltinfo, poninfo, onuinfo
      let childrenArray = item.child ?? item.ltinfo ?? item.poninfo ?? item.onuinfo ?? [];
      // Normalize to array if it's an object
      if (childrenArray && !Array.isArray(childrenArray)) {
        childrenArray = [childrenArray];
      }

      return {
        key: item.name || item.id || Math.random().toString(36).slice(2),
        title: item.name || item.model || item.type,
        children: Array.isArray(childrenArray) && childrenArray.length > 0 ? buildTree(childrenArray) : [],
        ...item,
      };
    });
  };

  const treeData = data ? buildTree(data) : [];

  const handleNodeSelect = (node: any) => {
    let newTableData = node.children || [];

    // Expand state object if it exists
    newTableData = newTableData.map((item: any) => {
      if (
        item.state &&
        typeof item.state === "object" &&
        !Array.isArray(item.state)
      ) {
        return { ...item, ...item.state };
      }

      return item;
    });
    setTableData(newTableData);
    const newColumns =
      newTableData.length > 0
        ? Object.keys(newTableData[0]).filter((key) => {
            if (key === "child" || key === "key" || key === "title")
              return false;
            const val = newTableData[0][key];

            return val === null || typeof val !== "object";
          })
        : [];

    setVisibleColumns(new Set([...newColumns, "actions"]));
    setPage(1); // reset page on new data
  };

  const tableColumns =
    tableData.length > 0
      ? Object.keys(tableData[0]).filter((key) => {
          if (key === "child" || key === "key" || key === "title") return false;
          const val = tableData[0][key];

          return val === null || typeof val !== "object";
        })
      : [];

  const columns = [
    ...tableColumns.map((col: string) => ({
      name: col.toUpperCase(),
      uid: col,
      sortable: true,
    })),
    { name: "ACTIONS", uid: "actions" },
  ];

  const statusOptions = [
    { name: "Active", uid: "active" },
    { name: "Paused", uid: "paused" },
    { name: "Vacation", uid: "vacation" },
  ];

  const hasSearchFilter = Boolean(filterValue);

  const headerColumns = useMemo(() => {
    return columns.filter((column) =>
      Array.from(visibleColumns).includes(column.uid),
    );
  }, [visibleColumns, columns]);

  const filteredItems = useMemo(() => {
    let filtered = [...tableData];

    if (hasSearchFilter) {
      filtered = filtered.filter((item) =>
        (item.name || "").toLowerCase().includes(filterValue.toLowerCase()),
      );
    }
    if (
      statusFilter !== "all" &&
      Array.from(statusFilter).length !== statusOptions.length
    ) {
      filtered = filtered.filter((item) =>
        Array.from(statusFilter).includes(item.status),
      );
    }

    return filtered;
  }, [tableData, filterValue, statusFilter, statusOptions.length]);

  const pages = Math.ceil(filteredItems.length / rowsPerPage) || 1;

  const items = useMemo(() => {
    const start = (page - 1) * rowsPerPage;
    const end = start + rowsPerPage;

    return filteredItems.slice(start, end);
  }, [page, filteredItems, rowsPerPage]);

  const sortedItems = useMemo(() => {
    return [...items].sort((a, b) => {
      const first = a[sortDescriptor.column];
      const second = b[sortDescriptor.column];
      const cmp = first < second ? -1 : first > second ? 1 : 0;

      return sortDescriptor.direction === "descending" ? -cmp : cmp;
    });
  }, [sortDescriptor, items]);

  const renderCell = useCallback((item: any, columnKey: string) => {
    if (columnKey === "actions") {
      let menuItems: string[] = [];

      if (item.type === "lt") {
        menuItems = ["lock", "unlock", "logger"];
      } else if (item.type === "v-ani") {
        menuItems = ["up", "down", "mibreset"];
      } else {
        menuItems = ["lock", "unlock"];
      }

      return (
        <Dropdown>
          <DropdownTrigger>
            <Button isIconOnly size="sm" variant="light">
              <VerticalDotsIcon className="text-default-300" />
            </Button>
          </DropdownTrigger>
          <DropdownMenu>
            {menuItems.map((menuItem) => (
              <DropdownItem
                key={menuItem}
                onPress={() => {
                  if (menuItem === 'logger') {
                    setCurrentItem(item);
                    setShowLoggerModal(true);
                  }
                }}
              >
                {menuItem}
              </DropdownItem>
            ))}
          </DropdownMenu>
        </Dropdown>
      );
    }
    if (columnKey === "status") {
      return (
        <Chip
          className="capitalize"
          color={(statusColorMap[item.status as string] as any) || "default"}
          size="sm"
          variant="flat"
        >
          {item[columnKey]}
        </Chip>
      );
    }
    const val = item[columnKey];

    if (val === null || val === undefined) return "";
    if (typeof val === "object") {
      return (
        <pre className="whitespace-pre-wrap max-h-40 overflow-auto">
          {JSON.stringify(val, null, 2)}
        </pre>
      );
    }

    return String(val);
  }, []);

  const onNextPage = useCallback(() => {
    if (page < pages) {
      setPage(page + 1);
    }
  }, [page, pages]);

  const onPreviousPage = useCallback(() => {
    if (page > 1) {
      setPage(page - 1);
    }
  }, [page]);

  const onRowsPerPageChange = useCallback((e: any) => {
    setRowsPerPage(Number(e.target.value));
    setPage(1);
  }, []);

  const onSearchChange = useCallback((value: any) => {
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

  const topContent = useMemo(() => {
    return (
      <div className="flex flex-col gap-4">
        <div className="flex justify-between gap-3 items-end">
          <Input
            isClearable
            className="w-full sm:max-w-[44%]"
            placeholder="Search by name..."
            startContent={<SearchIcon />}
            value={filterValue}
            onClear={() => onClear()}
            onValueChange={onSearchChange}
          />
          <div className="flex gap-3">
            <Dropdown>
              <DropdownTrigger className="hidden sm:flex">
                <Button
                  endContent={<ChevronDownIcon className="text-small" />}
                  variant="flat"
                >
                  Status
                </Button>
              </DropdownTrigger>
              <DropdownMenu
                disallowEmptySelection
                aria-label="Table Columns"
                closeOnSelect={false}
                selectedKeys={statusFilter}
                selectionMode="multiple"
                onSelectionChange={setStatusFilter}
              >
                {statusOptions.map((status) => (
                  <DropdownItem key={status.uid} className="capitalize">
                    {status.name}
                  </DropdownItem>
                ))}
              </DropdownMenu>
            </Dropdown>
            <Dropdown>
              <DropdownTrigger className="hidden sm:flex">
                <Button
                  endContent={<ChevronDownIcon className="text-small" />}
                  variant="flat"
                >
                  Columns
                </Button>
              </DropdownTrigger>
              <DropdownMenu
                disallowEmptySelection
                aria-label="Table Columns"
                closeOnSelect={false}
                selectedKeys={visibleColumns}
                selectionMode="multiple"
                onSelectionChange={setVisibleColumns}
              >
                {columns.map((column) => (
                  <DropdownItem key={column.uid} className="capitalize">
                    {column.name}
                  </DropdownItem>
                ))}
              </DropdownMenu>
            </Dropdown>
          </div>
        </div>
        <div className="flex justify-between items-center">
          <span className="text-default-400 text-small">
            Total {tableData.length} items
          </span>
          <label className="flex items-center text-default-400 text-small">
            Rows per page:
            <select
              className="bg-transparent outline-solid outline-transparent text-default-400 text-small"
              onChange={onRowsPerPageChange}
            >
              <option value="5">5</option>
              <option value="10">10</option>
              <option value="15">15</option>
            </select>
          </label>
        </div>
      </div>
    );
  }, [
    filterValue,
    statusFilter,
    visibleColumns,
    onRowsPerPageChange,
    tableData.length,
    onSearchChange,
    hasSearchFilter,
    columns,
    statusOptions,
  ]);

  const bottomContent = useMemo(() => {
    return (
      <div className="py-2 px-2 flex justify-between items-center">
        <span className="w-[30%] text-small text-default-400">
          {selectedKeys === "all"
            ? "All items selected"
            : `${selectedKeys.size} of ${filteredItems.length} selected`}
        </span>
        <Pagination
          isCompact
          showControls
          showShadow
          color="primary"
          page={page}
          total={pages}
          onChange={setPage}
        />
        <div className="hidden sm:flex w-[30%] justify-end gap-2">
          <Button
            isDisabled={pages === 1}
            size="sm"
            variant="flat"
            onPress={onPreviousPage}
          >
            Previous
          </Button>
          <Button
            isDisabled={pages === 1}
            size="sm"
            variant="flat"
            onPress={onNextPage}
          >
            Next
          </Button>
        </div>
      </div>
    );
  }, [selectedKeys, filteredItems.length, page, pages]);

  return (
    <DefaultLayout>
      <section className="flex flex-col gap-6 py-8 md:py-10">
        {!data && !loading && (
          <div className="max-w-4xl mx-auto">
            <form className="flex gap-4 items-end" onSubmit={handleConnect}>
              <div className="flex-1">
                <label htmlFor="oamIp" id="oamIp-label">
                  OAM IP
                </label>
                <Input
                  aria-labelledby="oamIp-label"
                  autoComplete="off"
                  id="oamIp"
                  placeholder="Enter OAM IP"
                  value={oamIp}
                  onValueChange={setOamIp}
                />
              </div>
              <div className="flex-1">
                <label htmlFor="username" id="username-label">
                  NetConf Username
                </label>
                <Input
                  aria-labelledby="username-label"
                  autoComplete="username"
                  id="username"
                  placeholder="Enter Username"
                  value={username}
                  onValueChange={setUsername}
                />
              </div>
              <div className="flex-1">
                <label htmlFor="password" id="password-label">
                  NetConf Password
                </label>
                <Input
                  aria-labelledby="password-label"
                  autoComplete="current-password"
                  id="password"
                  placeholder="Enter Password"
                  type="password"
                  value={password}
                  onValueChange={setPassword}
                />
              </div>
              <Button
                className="bg-linear-to-tr from-pink-500 to-yellow-500 text-white shadow-lg"
                radius="full"
                type="submit"
              >
                Start
              </Button>
            </form>
          </div>
        )}

        {loading && (
          <div className="fixed inset-0 bg-black flex items-center justify-center z-50">
            <MatrixRain />
            <div className="relative z-10 text-center">
              <Spinner color="success" size="lg" />
              <p className="text-white mt-4 text-xl">Connecting</p>
            </div>
          </div>
        )}

        {data && (
          <div className="flex gap-6">
            <div className="w-1/8">
              {treeData.map((node: any) => (
                <TreeNode
                  key={node.key}
                  node={node}
                  onSelect={handleNodeSelect}
                />
              ))}
            </div>
            <div className="w-7/8">
              <Table
                isHeaderSticky
                aria-label="Collector Table"
                bottomContent={bottomContent}
                bottomContentPlacement="outside"
                classNames={{
                  wrapper: "max-h-[382px]",
                }}
                selectedKeys={selectedKeys}
                selectionMode="multiple"
                sortDescriptor={sortDescriptor}
                topContent={topContent}
                topContentPlacement="outside"
                onSelectionChange={setSelectedKeys}
                onSortChange={(descriptor) =>
                  setSortDescriptor(descriptor as any)
                }
              >
                <TableHeader columns={headerColumns}>
                  {(column) => (
                    <TableColumn
                      key={column.uid}
                      align={column.uid === "actions" ? "center" : "start"}
                      allowsSorting={!!(column as any).sortable}
                    >
                      {column.name}
                    </TableColumn>
                  )}
                </TableHeader>
                <TableBody emptyContent={"No items found"} items={sortedItems}>
                  {(item) => (
                    <TableRow key={item.name || item.id || Math.random()}>
                      {(columnKey) => (
                        <TableCell>
                          {renderCell(item, columnKey as string)}
                        </TableCell>
                      )}
                    </TableRow>
                  )}
                </TableBody>
              </Table>
            </div>
          </div>
        )}

        {errorMessage && (
          <div className="fixed bottom-0 left-0 right-0 bg-red-600 text-white p-4 flex justify-between items-center z-50">
            <span>{errorMessage}</span>
            <Button
              size="sm"
              variant="light"
              onPress={() => setErrorMessage(null)}
            >
              Close
            </Button>
          </div>
        )}
        <LoggerModal
          isOpen={showLoggerModal}
          onClose={() => setShowLoggerModal(false)}
          requestId={sessionStorage.getItem("request_id")}
          requestParent={currentItem?.parent}
        />
      </section>
    </DefaultLayout>
  );
}
