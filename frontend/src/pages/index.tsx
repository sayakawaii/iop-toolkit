import { useEffect, useState } from "react";
import axios from "axios";
import { button as buttonStyles } from "@heroui/theme";
import { Button} from "@heroui/react";
import { title} from "@/components/primitives";
import DefaultLayout from "@/layouts/default";
import { API_BASE } from "@/config/api";
import {
  PieChart,
  Pie,
  Cell,
  Tooltip as ReTooltip,
  ResponsiveContainer,
  BarChart,
  Bar,
  XAxis,
  YAxis,
  Legend,
} from "recharts";

const baseUrl = API_BASE;
const SEARCH_API = baseUrl + "/api/library/search";
const COUNTERS_API = baseUrl + "/api/omcianalyzer/counters";

export default function IndexPage() {
  const [rows, setRows] = useState<any[]>([]);
  const [counters, setCounters] = useState<{ total: number; today: number } | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function load() {
      try {
        setLoading(true);
        // 1. fetch table data
        const r = await axios.post(SEARCH_API, {}); // empty params per spec
        const data = Array.isArray(r.data) ? r.data : r.data?.rows ?? [];
        setRows(data);

        // 4. fetch counters
        const c = await axios.get(COUNTERS_API);
        setCounters(c.data);
      } catch (e) {
        console.error(e);
      } finally {
        setLoading(false);
      }
    }
    load();
  }, []);

  // helpers to aggregate
  const countBy = (key: string) => {
    const map = new Map<string, number>();
    for (const r of rows) {
      const v = (r as any)[key] ?? "Unknown";
      map.set(v, (map.get(v) ?? 0) + 1);
    }
    return Array.from(map.entries()).map(([name, value]) => ({ name, value }));
  };

  // colors for pies
  const COLORS = [
    "#4F46E5",
    "#06B6D4",
    "#F59E0B",
    "#10B981",
    "#EF4444",
    "#8B5CF6",
    "#F97316",
    "#0EA5E9",
    "#A78BFA",
    "#34D399",
  ];

  // 2. percentage charts for Customer, Vendor, PonType, RGMode
  const customerData = countBy("Customer");
  const vendorData = countBy("Vendor");
  const ponTypeData = countBy("PonType");
  const rgModeData = countBy("RGMode");

  // 3. LSRelease timeline (count per LSRelease)
  const lsData = countBy("LSRelease").sort((a, b) => {
    // try numeric or lexical sort
    const na = Number(a.name);
    const nb = Number(b.name);
    if (!isNaN(na) && !isNaN(nb)) return na - nb;
    return String(a.name).localeCompare(String(b.name));
  });

  return (
    <DefaultLayout>
      <section className="flex flex-col items-center justify-center gap-4 py-8 md:py-10 w-full">
        <header className="flex items-center justify-between w-full">
          <span className={title({ size: "sm" })}>Dashboard&nbsp;</span>
          <div className="flex items-center gap-2">
            <Button className={buttonStyles()} onClick={() => window.location.reload()}>
              Refresh
            </Button>
          </div>
        </header>

        {loading ? (
          <div>Loading...</div>
        ) : (
          <>
            <div className="w-full grid grid-cols-1 md:grid-cols-3 gap-6">
              <div className="col-span-1 md:col-span-2 p-4 bg-white rounded shadow">
                <h3 className="font-medium mb-2">LSRelease timeline</h3>
                <div style={{ width: "100%", height: 240 }}>
                  <ResponsiveContainer>
                    <BarChart data={lsData}>
                      <XAxis dataKey="name" />
                      <YAxis />
                      <ReTooltip />
                      <Legend />
                      <Bar dataKey="value" fill={COLORS[0]} />
                    </BarChart>
                  </ResponsiveContainer>
                </div>
              </div>

              <div className="col-span-1 p-4 bg-white rounded shadow">
                <h3 className="font-medium mb-2">Counters</h3>
                <div className="text-lg">
                  Total parsed: <strong>{counters?.total ?? "—"}</strong>
                </div>
                <div className="text-lg mt-2">
                  Today: <strong>{counters?.today ?? "—"}</strong>
                </div>
              </div>
            </div>

            <div className="w-full grid grid-cols-1 md:grid-cols-4 gap-6 mt-6">
              {[
                { title: "Customer", data: customerData },
                { title: "Vendor", data: vendorData },
                { title: "PonType", data: ponTypeData },
                { title: "RGMode", data: rgModeData },
              ].map((d, _) => (
                <div key={d.title} className="p-4 bg-white rounded shadow">
                  <h4 className="font-medium mb-2">{d.title}</h4>
                  <div style={{ width: "100%", height: 200 }}>
                    <ResponsiveContainer>
                      <PieChart>
                        <Pie data={d.data} dataKey="value" nameKey="name" innerRadius={30} outerRadius={70} label>
                          {d.data.map((_, idx) => (
                            <Cell key={idx} fill={COLORS[idx % COLORS.length]} />
                          ))}
                        </Pie>
                        <ReTooltip />
                      </PieChart>
                    </ResponsiveContainer>
                  </div>
                </div>
              ))}
            </div>
          </>
        )}
      </section>
    </DefaultLayout>
  );
}
