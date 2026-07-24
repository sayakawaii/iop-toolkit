import { useState, useEffect, useMemo, useCallback } from "react";
import axios from "axios";
import { Button, Input, Table, TableHeader, TableColumn, TableBody, TableRow, TableCell, Dropdown, DropdownTrigger, DropdownMenu, DropdownItem, Tabs, Tab, Chip, Modal, ModalContent, ModalHeader, ModalBody, Checkbox } from "@heroui/react";
import { API_BASE } from "@/config/api";

const baseUrl = API_BASE;
const LOGGER_GET_API = baseUrl + "/api/collector/loggerGet";
const LOGGER_SET_API = baseUrl + "/api/collector/loggerSet";

const levelOptions = [
  { label: "CRITICAL", value: 2 },
  { label: "ERROR", value: 3 },
  { label: "WARNING", value: 4 },
  { label: "INFO", value: 6 },
  { label: "DEBUG", value: 7 },
];

const levelColorMap: { [key: number]: string } = {
  2: "danger",
  3: "danger",
  4: "warning",
  6: "primary",
  7: "secondary",
};

const getLevelName = (val: number) => {
  return levelOptions.find(o => o.value === val)?.label || "WARNING";
};

const AppTable = ({ app, data, levels, setLevels, filterValue, sortDescriptor, setSortDescriptor, onSearchChange, onClear, onModified, appStates, handleToggle, handleCheckLogs, overlay, setOverlay }: any) => {
  const hasSearchFilter = Boolean(filterValue);

  const filteredItems = useMemo(() => {
    let filtered = [...data];

    if (hasSearchFilter) {
      filtered = filtered.filter((item) =>
        item.id.toString().toLowerCase().includes(filterValue.toLowerCase()) ||
        item.module_name.toLowerCase().includes(filterValue.toLowerCase()) ||
        getLevelName(levels[app][item.id]).toLowerCase().includes(filterValue.toLowerCase())
      );
    }

    return filtered;
  }, [data, filterValue, levels, app]);

  const sortedItems = useMemo(() => {
    return [...filteredItems].sort((a, b) => {
      const first = a[sortDescriptor.column];
      const second = b[sortDescriptor.column];
      const cmp = first < second ? -1 : first > second ? 1 : 0;

      return sortDescriptor.direction === "descending" ? -cmp : cmp;
    });
  }, [sortDescriptor, filteredItems]);

  const topContent = useMemo(() => {
    return (
      <div className="flex flex-col gap-4">
        <div className="flex justify-between gap-3 items-end">
          <Input
            isClearable
            className="w-full sm:max-w-[44%]"
            placeholder="Search by id, name, or level..."
            onClear={() => onClear()}
            onValueChange={onSearchChange}
          />
          <div className="flex gap-3 items-center">
            <Checkbox
              isSelected={overlay}
              onValueChange={setOverlay}
            >
              Overlay
            </Checkbox>
            {appStates[app] === 'disabled' ? (
              <Button onPress={() => handleToggle(overlay ? 'overlay' : 'enable')}>Start</Button>
            ) : (
              <Button onPress={() => handleToggle('disable')}>Stop</Button>
            )}
            <Button onPress={() => handleCheckLogs(app)}>CheckLogs</Button>
          </div>
        </div>
      </div>
    );
  }, [filterValue, app, onSearchChange, onClear, appStates, handleToggle, handleCheckLogs]);

  return (
    <Table
      aria-label={`${app} Logger Modules`}
      topContent={topContent}
      topContentPlacement="outside"
      sortDescriptor={sortDescriptor}
      onSortChange={(descriptor) => setSortDescriptor(descriptor as any)}
    >
      <TableHeader>
        <TableColumn key="id" allowsSorting>
          ID
        </TableColumn>
        <TableColumn key="module_name" allowsSorting>
          Module Name
        </TableColumn>
        <TableColumn key="level" allowsSorting>
          Level
        </TableColumn>
      </TableHeader>
      <TableBody>
        {sortedItems.map((mod: any) => (
          <TableRow key={mod.id}>
            <TableCell>{mod.id}</TableCell>
            <TableCell>{mod.module_name}</TableCell>
            <TableCell>
              <Dropdown>
                <DropdownTrigger>
                  <Button variant="flat">
                    <Chip
                      className="capitalize"
                      color={(levelColorMap[levels[app]?.[mod.id] || 4] as any) || "default"}
                      size="sm"
                      variant="flat"
                    >
                      {getLevelName(levels[app]?.[mod.id] || 4)}
                    </Chip>
                  </Button>
                </DropdownTrigger>
                <DropdownMenu
                  onAction={(key) => {
                    setLevels((prev: any) => ({
                      ...prev,
                      [app]: {
                        ...prev[app],
                        [mod.id]: parseInt(key as string)
                      }
                    }));
                    onModified(app, true);
                  }}
                >
                  {levelOptions.map(option => (
                    <DropdownItem key={option.value}>
                      {option.label}
                    </DropdownItem>
                  ))}
                </DropdownMenu>
              </Dropdown>
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
};

interface LoggerModalProps {
  isOpen: boolean;
  onClose: () => void;
  requestId: string | null;
  requestParent: any;
}

export default function LoggerModal({ isOpen, onClose, requestId, requestParent}: LoggerModalProps) {
  const [loggerData, setLoggerData] = useState<any>(null);
  const [levels, setLevels] = useState<{ [key: string]: { [key: string]: number } }>({});
  const [appStates, setAppStates] = useState<{ [key: string]: 'enabled' | 'disabled' }>({});
  const [modifiedApps, setModifiedApps] = useState<{ [key: string]: boolean }>({});
  const [overlay, setOverlay] = useState(false);
  const [filterValue, setFilterValue] = useState("");
  const [sortDescriptor, setSortDescriptor] = useState<{
    column: string;
    direction: "ascending" | "descending";
  }>({
    column: "id",
    direction: "ascending",
  });

  useEffect(() => {
    if (isOpen) {
      axios.get(LOGGER_GET_API).then(response => {
        setLoggerData(response.data.logger_modules);
        // 初始化 levels 为默认值 4 (WARNING)
        const initialLevels: any = {};
        Object.keys(response.data.logger_modules).forEach(app => {
          initialLevels[app] = {};
          response.data.logger_modules[app].forEach((mod: any) => {
            initialLevels[app][mod.id] = 4;
          });
        });
        setLevels(initialLevels);
        // 初始化 appStates 为 'disabled'
        const initialStates: { [key: string]: 'enabled' | 'disabled' } = {};
        Object.keys(response.data.logger_modules).forEach(app => {
          initialStates[app] = 'disabled';
        });
        setAppStates(initialStates);
        // 初始化 modifiedApps 为 false
        const initialModified: { [key: string]: boolean } = {};
        Object.keys(response.data.logger_modules).forEach(app => {
          initialModified[app] = false;
        });
        setModifiedApps(initialModified);
      }).catch(error => {
        console.error("Error fetching logger data:", error);
      });
    }
  }, [isOpen]);

  const handleToggle = (action: 'enable' | 'disable' | 'overlay') => {
    const modifiedAppsList = Object.keys(modifiedApps).filter(modApp => modifiedApps[modApp]);
    // stop 操作不检查 modified 状态，直接发送请求
    // start 和 overlay 操作需要检查是否有修改，如果没有修改就不发送请求，也不切换状态
    if ((action === 'enable' || action === 'overlay') && modifiedAppsList.length === 0) {
      alert('Please modify at least one logger level before starting');
      return;
    }

    const newState: 'enabled' | 'disabled' = (action === 'enable' || action === 'overlay') ? 'enabled' : 'disabled';
    setAppStates(prev => {
      const updated: { [key: string]: 'enabled' | 'disabled' } = { ...prev };
      // 所有 app 状态同步到一致
      Object.keys(updated).forEach(key => {
        updated[key] = newState;
      });
      return updated;
    });

    const payload: {
      request_id: string;
      logger_modules: { [key: string]: any[] };
      action: string;
      request_parent: string;
    } = {
      request_id: requestId || "",
      logger_modules: {},
      action: action,
      request_parent: requestParent
    };

    // 如果是 stop 操作且没有修改，发送所有 app 的配置
    const appsToSend = action === 'disable' && modifiedAppsList.length === 0 
      ? Object.keys(loggerData) 
      : modifiedAppsList;

    appsToSend.forEach(modApp => {
      const modules = loggerData[modApp].map((mod: any) => ({
        ...mod,
        level: levels[modApp][mod.id]
      }));
      payload.logger_modules[modApp] = modules;
    });

    axios.post(LOGGER_SET_API, payload).then(() => {
      alert(`Logger ${action}d successfully`);
      // 重置 modifiedApps
      setModifiedApps(prev => {
        const reset = { ...prev };
        appsToSend.forEach(modApp => reset[modApp] = false);
        return reset;
      });
    }).catch(error => {
      console.error(`Error ${action}ing logger:`, error);
      // 如果失败，回滚状态
      setAppStates(prev => {
        const rolled = { ...prev };
        Object.keys(rolled).forEach(key => {
          rolled[key] = (action === 'enable' || action === 'overlay') ? 'disabled' : 'enabled';
        });
        return rolled;
      });
    });
  };

  const handleCheckLogs = () => {
    window.open(`/collector/logs?request_id=${requestId}`, '_blank');
  };

  const onSearchChange = useCallback((value: any) => {
    if (value) {
      setFilterValue(value);
    } else {
      setFilterValue("");
    }
  }, []);

  const onClear = useCallback(() => {
    setFilterValue("");
  }, []);

  const apps = Object.keys(loggerData || {});

  return (
    <Modal isOpen={isOpen} onOpenChange={onClose} size="5xl" scrollBehavior="inside">
      <ModalContent>
        <ModalHeader>Logger Configuration</ModalHeader>
        <ModalBody>
          {loggerData ? (
              <Tabs>
                {apps.map(app => (
                  <Tab key={app} title={app}>
                    <AppTable
                      app={app}
                      data={loggerData[app] || []}
                      levels={levels}
                      setLevels={setLevels}
                      filterValue={filterValue}
                      sortDescriptor={sortDescriptor}
                      setSortDescriptor={setSortDescriptor}
                      onSearchChange={onSearchChange}
                      onClear={onClear}
                      onModified={(app: string, modified: boolean) => setModifiedApps(prev => ({ ...prev, [app]: modified }))}
                      appStates={appStates}
                      handleToggle={handleToggle}
                      handleCheckLogs={handleCheckLogs}
                      overlay={overlay}
                      setOverlay={setOverlay}
                    />
                  </Tab>
                ))}
              </Tabs>
          ) : (
            <p>Loading...</p>
          )}
        </ModalBody>
      </ModalContent>
    </Modal>
  );
}