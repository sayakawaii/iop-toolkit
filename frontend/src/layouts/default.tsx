import { Link } from "@heroui/link";

import { Navbar } from "@/components/navbar";
import {
  HeartFilledIcon,
} from "@/components/icons";

export default function DefaultLayout({
  children,
  fullWidth = false,
}: {
  children: React.ReactNode;
  fullWidth?: boolean;
}) {
  return (
    <div className="relative flex flex-col h-screen">
      <Navbar />
      {/* <main className="container mx-auto max-w-7xl px-6 flex-grow pt-16"> */}
      <main
        className={
          fullWidth
            ? "w-full px-6 flex-grow pt-16"
            : "container mx-auto max-w-7xl px-6 flex-grow pt-16"
        }
      >
        {children}
      </main>
      <footer className="w-full flex items-center justify-center py-3">
        <Link
          isExternal
          className="flex items-center gap-1 text-current"
          href="https://confluence.ext.net.nokia.com/display/Transpt2/Troubleshooting+-+OMCI+Analyzer"
          title="IOP Toolkit homepage"
        >
          <span className="text-default-600">Powered by</span>
            <HeartFilledIcon className="text-danger" />
          <p className="text-primary">Transport2 eONU</p>
        </Link>
      </footer>
    </div>
  );
}
