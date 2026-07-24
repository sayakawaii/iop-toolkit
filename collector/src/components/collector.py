from typing import Any, List
from functools import wraps

from paramiko import SSHClient
from models.host import OltHost, LogServer
from components.connection import SSHConnection
from components.applauncher import AppClient
import logging

from utils.olttype import is_brugal_platform

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s][%(name)s] %(message)s",
    datefmt="%Y-%m-%d %H:%M:%S"
)
logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)

#normal
BASE_REMOTE_LOG_COMMANDS = [
    "ykraken 'fwd {app} log set output loglevel LOGFILE YLOG_DISABLED'",
    "ykraken 'fwd {app} log set output loglevel SYSLOG YLOG_DISABLED'",
    "ykraken 'fwd {app} log set output loglevel CONSOLE YLOG_DISABLED'",
    "ykraken 'fwd {app} log set output loglevel TNDD YLOG_DISABLED'",
    "ykraken 'fwd {app} log set output loglevel REMOTE YLOG_DEBUG'",
    "ykraken 'fwd {app} log set remote disconnect'",
]

REMOTE_TCP_COMMAND = "ykraken 'fwd {app} log set remote connect tcp:{ip}:{port}'"
MODULE_LOGLEVEL_COMMAND = "ykraken 'fwd {app} log set module loglevel {module} {level}'"


def get_enable_remote_commands(target: LogServer, app: str, app_config: List[dict]) -> str:
    commands: List[str] = [cmd.format(app=app) for cmd in BASE_REMOTE_LOG_COMMANDS]

    for item in app_config or []:
        module_name = item.get("module_name") or item.get("module")
        level = item.get("level", "")
        if module_name:
            commands.append(MODULE_LOGLEVEL_COMMAND.format(app=app, module=module_name, level=level))

    commands.append(REMOTE_TCP_COMMAND.format(app=app, ip=target.host, port=target.port))

    commands_str = ";".join(commands)
    logger.info("Enable remote log commands:")
    logger.info(commands_str)
    return commands_str

ONU_ENGINE_BASE_REMOTE_LOG_COMMANDS = [
"curl -X POST 127.0.0.1:8801/v_onu_mgnt/v1/updateLoggingLevel/*/trace",
"curl -i -X POST -d '{{}}' http://127.0.0.1:8801/v_onu_mgnt/v1/updateDefaultLoggingLevel/trace",
"curl -X PATCH http://127.0.0.1:8880/api/v2/tnd/onus/tracers -d '{{\"enabled\": true, \"protected\": false}}'",
"curl -X PATCH http://127.0.0.1:8880/api/v2/tnd/yang2omci/tracer -d '{{\"enabled\": true}}'",
]

ONU_ENGINE_DELETE_REMOTE_LOG_COMMAND = "curl -X DELETE http://127.0.0.1:8880/api/v2/tnd/loggers/outputs/{app}"
ONU_ENGINE_REMOTE_TCP_COMMAND = "curl -X POST http://127.0.0.1:8880/api/v2/tnd/loggers/outputs -d '{{\"name\":\"{app}\", \"uri\":\"tcp://{ip}:{port}\"}}'"
ONU_ENGINE_MODULE_LOGLEVEL_COMMAND = "curl -X PATCH http://127.0.0.1:8880/api/v2/tnd/{module}/tracer -d '{{\"enabled\": true}}'"

def onu_engine_get_enable_remote_commands(target: LogServer, app: str, app_config: List[dict]) -> str:
    commands: List[str] = [cmd.format(app=app) for cmd in ONU_ENGINE_BASE_REMOTE_LOG_COMMANDS]

    for item in app_config or []:
        module_name = item.get("module_name") or item.get("module")
        level = item.get("level", "")
        if module_name:
            commands.append(ONU_ENGINE_MODULE_LOGLEVEL_COMMAND.format(app=app, module=module_name, level=level))

    commands.append(ONU_ENGINE_DELETE_REMOTE_LOG_COMMAND.format(app=app))
    commands.append(ONU_ENGINE_REMOTE_TCP_COMMAND.format(app=app, ip=target.host, port=target.port))

    commands_str = ";".join(commands)
    logger.info("Enable remote log commands for onu_engine:")
    logger.info(commands_str)
    return commands_str

#overlay
BASE_OVERLAY_LOG_COMMANDS = [
    "mkdir -p /mnt/persistent/rootfs-overlay/isam/slot_default/{app}/run/",
    "cp /isam/slot_default/{app}/run/extra_isam_params  /mnt/persistent/rootfs-overlay/isam/slot_default/{app}/run/extra_isam_params",
    "echo --cmd-file=debug.cmd >> /mnt/persistent/rootfs-overlay/isam/slot_default/{app}/run/extra_isam_params",
    "echo log set output loglevel LOGFILE YLOG_DISABLED  >> /mnt/persistent/rootfs-overlay/isam/slot_default/{app}/run/debug.cmd",
    "echo log set output loglevel SYSLOG YLOG_DISABLED  >> /mnt/persistent/rootfs-overlay/isam/slot_default/{app}/run/debug.cmd",
    "echo log set output loglevel CONSOLE YLOG_DISABLED  >> /mnt/persistent/rootfs-overlay/isam/slot_default/{app}/run/debug.cmd",
    "echo log set output loglevel TNDD YLOG_DISABLED  >> /mnt/persistent/rootfs-overlay/isam/slot_default/{app}/run/debug.cmd",
    "echo log set output loglevel REMOTE YLOG_DEBUG  >> /mnt/persistent/rootfs-overlay/isam/slot_default/{app}/run/debug.cmd",
    "echo log set remote disconnect  >> /mnt/persistent/rootfs-overlay/isam/slot_default/{app}/run/debug.cmd",
]

OVERLAY_TCP_COMMAND = "echo log set remote connect tcp:{ip}:{port} >> /mnt/persistent/rootfs-overlay/isam/slot_default/{app}/run/debug.cmd"
OVERLAY_MODULE_LOGLEVEL_COMMAND = "echo log set module loglevel {module} {level} >> /mnt/persistent/rootfs-overlay/isam/slot_default/{app}/run/debug.cmd"

OVERLAY_GLOB_ENGINE_COMMAND = '''
globtty=/isam/slot_default/run/tty_dev22_ext
echo "[INFO] Waiting for device $globtty to appear..."
while [ ! -c "$globtty" ]; do
    sleep 5
done
echo "[OK] Device $globtty detected. check engine"
cat <<-EOF >> $globtty
log logRemote ip {ip}:{port}
log tndToCloud 1
EOF
'''


OVERLAY_ENGINE_ENGINE_COMMAND = '''
wait_for_port() {{
    for i in $(seq 1 30); do
        nc -z 127.0.0.1 8801 && return 0
        sleep 1
    done
    return 1
}}

wait_for_port || {{
    echo "engine API not ready"
    exit 1
}}

curl -X POST 127.0.0.1:8801/v_onu_mgnt/v1/updateLoggingLevel/*/trace
curl -i -X POST -d '{{}}' http://127.0.0.1:8801/v_onu_mgnt/v1/updateDefaultLoggingLevel/trace
curl -X PATCH http://127.0.0.1:8880/api/v2/tnd/onus/tracers -d '{{\"enabled\": true, \"protected\": false}}'
curl -X PATCH http://127.0.0.1:8880/api/v2/tnd/yang2omci/tracer -d '{{\"enabled\": true}}'
curl -X DELETE http://127.0.0.1:8880/api/v2/tnd/loggers/outputs/{app}
curl -X POST http://127.0.0.1:8880/api/v2/tnd/loggers/outputs -d '{{\"name\":\"{app}\", \"uri\":\"tcp://{ip}:{port}\"}}'
'''

def get_overlay_remote_commands(target: LogServer, app: str, app_config: List[dict]) -> str:
    commands: List[str] = [cmd.format(app=app) for cmd in BASE_OVERLAY_LOG_COMMANDS]

    for item in app_config or []:
        module_name = item.get("module_name") or item.get("module")
        level = item.get("level", "")
        if module_name:
            commands.append(OVERLAY_MODULE_LOGLEVEL_COMMAND.format(app=app, module=module_name, level=level))

    commands.append(OVERLAY_TCP_COMMAND.format(app=app, ip=target.host, port=target.port))

    commands_str = ";".join(commands)
    logger.info("Enable remote log commands for overlay:")
    logger.info(commands_str)
    return commands_str

def get_overlay_glob_commands(target: LogServer) -> str:
    commands_str = OVERLAY_GLOB_ENGINE_COMMAND.format(ip=target.host, port=target.port)
    logger.info("Enable remote log commands for overlay glob:")
    logger.info(commands_str)
    return commands_str

def get_overlay_engine_commands(target: LogServer, app: str) -> str:
    commands_str = OVERLAY_ENGINE_ENGINE_COMMAND.format(app=app, ip=target.host, port=target.port)
    logger.info("Enable remote log commands for overlay engine:")
    logger.info(commands_str)
    return commands_str

def _check_and_reconnect(self):
    """Check SSH connection and reconnect if needed"""
    if self.ssh is None or not self.ssh.get_transport() or not self.ssh.get_transport().is_active():
        logger.info("SSH connection lost, reconnecting...")
        self.ssh = self._conn.connect()
        if self.ssh is None:
            raise RuntimeError("SSH reconnection failed")
        logger.info("SSH reconnected successfully")

def _is_last_attempt(attempt, max_retries):
    """Check if this is the last retry attempt"""
    return attempt >= max_retries - 1

def ensure_connected(func):
    """Decorator to ensure SSH connection is alive before executing method"""
    @wraps(func)
    def wrapper(self, *args, **kwargs):
        max_retries = 2
        for attempt in range(max_retries):
            try:
                _check_and_reconnect(self)
                return func(self, *args, **kwargs)
                
            except (EOFError, OSError) as e:
                if _is_last_attempt(attempt, max_retries):
                    logger.error(f"Max retries reached. Final error: {type(e).__name__}: {e}")
                    raise
                
                logger.warning(f"Connection error during execution (attempt {attempt + 1}/{max_retries}): {type(e).__name__}: {e}")
                logger.info("Attempting to reconnect...")
                try:
                    _check_and_reconnect(self)
                    logger.info("SSH reconnected successfully, retrying operation...")
                except Exception as reconnect_error:
                    logger.error(f"Reconnection failed: {reconnect_error}")
                    raise
        
        raise RuntimeError("Failed to execute after maximum retries")
    return wrapper

class LogCollector:
    def __init__(self, target: OltHost, jump: OltHost | None = None):
        self.target = target
        self.jump = jump
        self._conn = SSHConnection(self.target, self.jump)
        self.ssh = self._conn.connect()
        if self.ssh is None:
            raise RuntimeError("SSH connection failed")

    def close(self):
        if self._conn:
            self._conn.close()
            self.ssh = None

    def __enter__(self):
        return self

    def __exit__(self, exc_type, exc_val, exc_tb):
        self.close()
        return False

    @ensure_connected
    def get_console_logs(self, cmd: str):
        stdin, stdout, stderr = self.ssh.exec_command(cmd, get_pty=True)
        return stdout.read().decode()

    @ensure_connected
    def get_confd_cli_logs(self, cmd: str):
        client = AppClient(self.ssh)
        client.open_shell()
        client.enter_confd_cli()
        return client.run_confd_cli_cmd(cmd)

    @ensure_connected
    def run_app_debug_command(self, app: str, cmd: str):
        client = AppClient(self.ssh)
        client.open_shell()
        client.enter_app(app)
        return client.run_app_cmd(cmd)
    
    @ensure_connected
    def run_glob_debug_command(self, isbrugal: bool, cmd: str):
        client = AppClient(self.ssh)
        client.open_shell()
        client.enter_glob_app(isbrugal)
        return client.run_glob_cmd(cmd)

    def enable_glob_remote(self, target: LogServer):
        boardname = self.get_console_logs("cat /tmp/boardname")
        logger.info(f"Detected boardname: {boardname.strip()}")
        cmd = f"log logRemote disconnect\rlog logRemote ip {target.host}:{target.port}\rlog tndToCloud 1\r"
        if is_brugal_platform(boardname.strip().lower()[0:6]):
            return self.run_glob_debug_command(True, cmd)
        return self.run_glob_debug_command(False, cmd)

    @ensure_connected
    def enable_glob_remote_overlay(self, target: LogServer):
        cmd = get_overlay_glob_commands(target)
        
        # Combine all commands into one execution to save SSH connection time
        # Use heredoc to preserve quotes and special characters
        commands = "mkdir -p /mnt/persistent/rootfs-overlay/etc/init.d; "
        commands += "cp /etc/init.d/S72resource_monitor /mnt/persistent/rootfs-overlay/etc/init.d/S72resource_monitor; "
        commands += "sed -i '/start)/a\\      /mnt/persistent/glob_remote_overlay.sh &' /mnt/persistent/rootfs-overlay/etc/init.d/S72resource_monitor; "
        commands += f"cat >> /mnt/persistent/glob_remote_overlay.sh << 'OVERLAY_EOF'\n{cmd}\nOVERLAY_EOF\n"
        commands += "chmod +x /mnt/persistent/glob_remote_overlay.sh"
        
        stdin, stdout, stderr = self.ssh.exec_command(commands, get_pty=True)
        return stdout.read().decode()

    @ensure_connected
    def enable_engine_remote_overlay(self, target: LogServer, app: str):
        cmd = get_overlay_engine_commands(target, app)
        
        # Combine all commands into one execution to save SSH connection time
        # Use heredoc to preserve quotes and special characters
        commands = "mkdir -p /mnt/persistent/rootfs-overlay/etc/init.d; "
        commands += "cp /etc/init.d/S89endofinit_hook /mnt/persistent/rootfs-overlay/etc/init.d/S89endofinit_hook; "
        commands += "sed -i '/dump_pre_reboot_action_list/a\\      /mnt/persistent/engine_remote_overlay.sh &' /mnt/persistent/rootfs-overlay/etc/init.d/S89endofinit_hook; "
        commands += f"cat >> /mnt/persistent/engine_remote_overlay.sh << 'OVERLAY_EOF'\n{cmd}\nOVERLAY_EOF\n"
        commands += "chmod +x /mnt/persistent/engine_remote_overlay.sh"
        
        stdin, stdout, stderr = self.ssh.exec_command(commands, get_pty=True)
        return stdout.read().decode()

    @ensure_connected
    def enable_remote(self, target: LogServer, app: str, app_config: List[dict]):
        if app == "glob_app":
            return self.enable_glob_remote(target)
        if app == "onu_engine":
            commands = onu_engine_get_enable_remote_commands(target, app, app_config)
        else:
            commands = get_enable_remote_commands(target, app, app_config)
        stdin, stdout, stderr = self.ssh.exec_command(commands, get_pty=True)
        return stdout.read().decode()

    @ensure_connected
    def enable_overlay(self, target: LogServer, app: str, app_config: List[dict]):
        logger.info(f"Enabling overlay for app={app}")
        if app == "glob_app":
            return self.enable_glob_remote_overlay(target)
        if app == "onu_engine":
            return self.enable_engine_remote_overlay(target, app)
        commands = get_overlay_remote_commands(target, app, app_config)
        stdin, stdout, stderr = self.ssh.exec_command(commands, get_pty=True)
        return stdout.read().decode()

    @ensure_connected
    def disable_overlay(self, app: str):
        logger.info(f"Disabling overlay for app={app}")
        if app == "glob_app":
            commands = "rm /mnt/persistent/glob_remote_overlay.sh; rm /mnt/persistent/rootfs-overlay/etc/init.d/S72resource_monitor"
        elif app == "onu_engine":
            commands = "rm /mnt/persistent/engine_remote_overlay.sh; rm /mnt/persistent/rootfs-overlay/etc/init.d/S89endofinit_hook"
        else:
            commands = f"rm /mnt/persistent/rootfs-overlay/isam/slot_default/{app}/run/debug.cmd"
        stdin, stdout, stderr = self.ssh.exec_command(commands, get_pty=True)
        return stdout.read().decode()