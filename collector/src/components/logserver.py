import os
import socketserver
import socket
import logging
import threading
from typing import Any, Dict, List

from components.collector import LogCollector
from components.minio import MinIOLogHandler
from models.host import LogServer
from utils.remotelog import Remotelog
from utils.env import Platform_Env

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s][%(name)s] %(message)s",
    datefmt="%Y-%m-%d %H:%M:%S"
)
logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)

class LogTCPHandler(socketserver.BaseRequestHandler):
    def handle(self):
        client_ip = self.client_address[0]
        file_path = getattr(self.server, "logfile")
        logger.info(f"[+] Logging from {client_ip} -> {file_path}")
        with open(file_path, "ab+") as f:
            while True:
                data = self.request.recv(4096)
                if not data:
                    break
                f.write(data)
                f.flush()

class ThreadedTCPServer(socketserver.ThreadingMixIn, socketserver.TCPServer):
    daemon_threads = True
    allow_reuse_address = True

class LogServerManager:
    def __init__(self, host_ip : any, appname : str, logfile : str, base_port=10000):
        self.base_port = base_port
        self.logfile = logfile
        self.appname = appname
        self.host_ip = host_ip
        self.server = None

    # ---------- Find free port ----------
    def find_free_port(self):
        port = self.base_port
        while True:
            with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
                if s.connect_ex(("127.0.0.1", port)) != 0:
                    return self.host_ip, port
            port += 1

    def start_tcp_server(self, port=10000):
        self.server = ThreadedTCPServer(("0.0.0.0", port), LogTCPHandler)
        # 将 LogServerManager 的 logfile 传给 handler 使用
        self.server.logfile = self.logfile
        logger.info(f"[*] TCP Log Server running on 0.0.0.0:{port}")
        self.server.serve_forever()

    def stop_server(self):
        if self.server:
            logger.info("[*] Stopping Log Server...")
            self.server.shutdown()
            self.server.server_close()
            logger.info("[*] Log Server Stopped")


class LogServerSupervisor:
    """Manage multiple log servers and their log files for multiple apps."""

    def __init__(
        self,
        collector: LogCollector,
        host_ip: str,
        base_port: int = 10000,
        log_root: str = None,
        bucket_name: str = "olt-logs",
    ):
        self.collector = collector
        self.host_ip = host_ip
        self.base_port = base_port
        # Set log_root based on platform
        if log_root is None:
            if Platform_Env.platform == "linux":
                self.log_root = "/app/logs"
            else:
                self.log_root = "/MyData/Project/collector/logs/"
        else:
            self.log_root = log_root
        self.bucket_name = bucket_name
        self.managers: Dict[str, LogServerManager] = {}
        self.threads: Dict[str, threading.Thread] = {}
        self.minio_objects: Dict[str, str] = {}
        self.logfiles: Dict[str, str] = {}
        self.app_configs: Dict[str, Any] = {}
        self.logenable_type: str = "enable"  # or "overlay"

    def _extract_apps(self, logger_modules: Any) -> List[str]:
        """Extract app names from logger_modules payload."""
        if isinstance(logger_modules, dict):
            return list(logger_modules.keys())
        if isinstance(logger_modules, list):
            apps: List[str] = []
            for item in logger_modules:
                if isinstance(item, str):
                    apps.append(item)
                elif isinstance(item, dict):
                    apps.extend(item.keys())
            return apps
        if isinstance(logger_modules, str):
            return [logger_modules]
        return []

    def _extract_app_configs(self, logger_modules: Any) -> Dict[str, Any]:
        """Extract app configs keyed by app name."""
        if isinstance(logger_modules, dict):
            return logger_modules
        if isinstance(logger_modules, list):
            result: Dict[str, Any] = {}
            for item in logger_modules:
                if isinstance(item, dict):
                    result.update(item)
            return result
        if isinstance(logger_modules, str):
            return {logger_modules: []}
        return {}

    def start(self, logger_modules: Any, action: str) -> List[Dict[str, Any]]:
        """Start log servers for each app derived from logger_modules."""
        apps = self._extract_apps(logger_modules)
        self.app_configs = self._extract_app_configs(logger_modules)
        if not apps:
            raise ValueError("No app found in logger_modules")

        logger.info(f"Starting log servers for apps: {apps}")
        objects, files = Remotelog(self.collector.jump.host, self.collector.target.host, self.log_root).new(apps)

        for app, object_name, logfile in zip(apps, objects, files):
            manager = LogServerManager(self.host_ip, app, logfile, base_port=self.base_port)
            host_ip, port = manager.find_free_port()
            log_server = LogServer(host_ip, port)

            thread = threading.Thread(target=manager.start_tcp_server, args=(port,), daemon=True)
            thread.start()

            self.managers[app] = manager
            self.threads[app] = thread
            self.minio_objects[app] = object_name
            self.logfiles[app] = logfile

            # Enable remote logging for the app on the target device
            if action == "enable":
                self.collector.enable_remote(log_server, app, self.app_configs.get(app, []))
                logger.info(f"[*] started_tcp_server for app={app} on {host_ip}:{port}")
            elif action == "overlay":
                self.logenable_type = "overlay"
                self.collector.enable_overlay(log_server, app, self.app_configs.get(app, []))
                logger.info(f"[*] overlay enabled for app={app} on {host_ip}:{port}")

        return [{"app": app, "minio": object_name, "size": 0} for app, object_name in self.minio_objects.items()]

    def stop(self):
        """Stop all running log servers."""
        for app, manager in self.managers.items():
            logger.info(f"Stopping log server for app={app}")
            manager.stop_server()
            if self.logenable_type == "overlay":
                self.collector.disable_overlay(app)
        for app, thread in self.threads.items():
            if thread.is_alive():
                thread.join()
                logger.info(f"Thread for app={app} joined")
        # Upload collected logs to MinIO and collect results
        minio = MinIOLogHandler()
        results: List[Dict[str, Any]] = []
        for app, object_name in self.minio_objects.items():
            logfile = self.logfiles.get(app)
            if not logfile:
                continue
            try:
                size_mb = round(os.path.getsize(logfile) / (1024 * 1024), 2) if os.path.exists(logfile) else 0.0
                minio.upload_file(self.bucket_name, object_name, logfile)
                logger.info(f"Uploaded logfile for app={app} -> {object_name}")
                results.append({"app": app, "minio": object_name, "size": size_mb})
            except Exception as exc:  # pylint: disable=broad-except
                logger.error(f"Failed to upload logfile for app={app}: {exc}")
        self.managers.clear()
        self.threads.clear()
        self.minio_objects.clear()
        self.logfiles.clear()
        return results

    def stop_app(self, app: str):
        """Stop a single app's log server if running."""
        manager = self.managers.get(app)
        thread = self.threads.get(app)
        if manager:
            logger.info(f"Stopping log server for app={app}")
            manager.stop_server()
        if thread and thread.is_alive():
            thread.join()
            logger.info(f"Thread for app={app} joined")
        # Upload collected log for this app to MinIO
        object_name = self.minio_objects.get(app)
        logfile = self.logfiles.get(app)
        if object_name and logfile:
            try:
                MinIOLogHandler().upload_file(self.bucket_name, object_name, logfile)
                logger.info(f"Uploaded logfile for app={app} -> {object_name}")
            except Exception as exc:  # pylint: disable=broad-except
                logger.error(f"Failed to upload logfile for app={app}: {exc}")
        self.managers.pop(app, None)
        self.threads.pop(app, None)
        self.minio_objects.pop(app, None)
        self.logfiles.pop(app, None)
