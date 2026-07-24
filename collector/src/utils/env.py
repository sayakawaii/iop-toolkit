
import os
import subprocess
import json
import socket
import platform
import logging

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s][%(name)s] %(message)s",
    datefmt="%Y-%m-%d %H:%M:%S"
)
logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)

class Environment:
    def __init__(self, platform_name : str = None):
        # 自动检测平台
        if platform_name is None:
            system = platform.system().lower()  # "Windows" / "Linux"
            if "windows" in system:
                self.platform = "windows"
            else:
                self.platform = "linux"
        else:
            self.platform = platform_name.lower()

        # Resolve the host IP that OLTs will push logs back to.
        # Priority: COLLECTOR_HOST_IP env var > platform auto-detection.
        env_ip = os.environ.get("COLLECTOR_HOST_IP", "").strip()
        if env_ip:
            self.ipaddress = env_ip
        elif self.platform == "windows":
            self.ipaddress = self.get_valid_ipv4_windows()
        else:
            self.ipaddress = self.get_valid_ipv4_linux()

        logger.info(f"platform={self.platform}, ip={self.ipaddress}")
    # ---------- Get Windows host IP ----------
    def get_valid_ipv4_windows(self):
        # 执行 powershell 获取所有 IPv4 地址（以 JSON 格式）
        ps_cmd = [
            "powershell.exe",
            "-Command",
            "Get-NetIPAddress -AddressFamily IPv4 | ConvertTo-Json"
        ]
        output = subprocess.check_output(ps_cmd, encoding="utf-8")
        entries = json.loads(output)

        if isinstance(entries, dict):
            entries = [entries]

        def is_valid(entry):
            ip = entry["IPAddress"]
            alias = entry["InterfaceAlias"]
            state = entry["AddressState"]

            # ---- 过滤规则 ----
            # 1. 必须是 Preferred
            if state != 4:
                return False

            # 2. 忽略 loopback
            if ip.startswith("127."):
                return False

            # 3. 忽略 link-local 169.254.x.x
            if ip.startswith("169.254."):
                return False

            # 4. 忽略 WSL / vEthernet / Virtual 接口
            bad_alias = ["WSL", "vEthernet", "Virtual", "Hyper-V", "Docker"]
            if any(bad in alias for bad in bad_alias):
                return False

            return True

        # 找到所有有效 IP
        valid = [e["IPAddress"] for e in entries if is_valid(e)]

        if not valid:
            raise RuntimeError("No valid IPv4 found")

        return valid[0]   # 通常第一个就是 Wi-Fi
    # ---------- Get Linux host IP ----------
    def get_valid_ipv4_linux(self):
        # Auto-detect the primary outbound IPv4 address (no external traffic sent).
        s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        try:
            s.connect(("8.8.8.8", 80))
            return s.getsockname()[0]
        except Exception:
            return "127.0.0.1"
        finally:
            s.close()

Platform_Env = Environment()