import time
import logging

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s][%(name)s] %(message)s",
    datefmt="%Y-%m-%d %H:%M:%S"
)
logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)

class AppClient:
    def __init__(self, ssh):
        self.ssh = ssh
        self.chan = None

    def open_shell(self):
        self.chan = self.ssh.invoke_shell()
        self.chan.settimeout(2)
        time.sleep(1)
        self._drain()

    def _drain(self):
        """clear buffer"""
        output = ""
        while True:
            try:
                chunk = self.chan.recv(9999).decode("utf-8", "ignore")
                if not chunk:
                    break
                output += chunk
                time.sleep(0.1)
            except:
                break
        return output

    def send_and_wait(self, cmd, prompt=">> ", timeout=10):
        """send command and wait until prompt appears"""

        self.chan.send(cmd + "\n")
        output = ""
        start = time.time()

        while time.time() - start < timeout:
            try:
                chunk = self.chan.recv(9999).decode("utf-8", "ignore")
                output += chunk

                if prompt in output:
                    return output
            except:
                pass

            time.sleep(0.2)

        raise TimeoutError(f"wait prompt {prompt} timeout, received content:\n{output}")

    def enter_app(self, app:str):
        logger.info(f"cd enter {app} directory")
        self.send_and_wait(f"cd /isam/slot_default/{app}/run/", "#")

        logger.info("start calamares")
        self.send_and_wait("calamares -cnt tty1_ext", "[calamares]")
        out = self.send_and_wait("\r", ">> ")
        logger.info(f"enter {app} successfully")
        return out

    def enter_glob_app(self, isbrugal: bool):
        logger.info("cd enter glob directory")
        if isbrugal:
            self.send_and_wait("cd /isam/slot_default/glob_app/run/", "#")
        else:
            self.send_and_wait("cd /isam/slot_default/run/", "#")

        logger.info("start calamares")
        self.send_and_wait("calamares -cnt tty_dev22_ext", "[calamares]")
        out = self.send_and_wait("\r", "---> ")
        logger.info("enter glob successfully")
        return out

    def enter_confd_cli(self):
        logger.info("start confd cli")
        out = self.send_and_wait("confd_cli -u techsupport", "#")
        logger.info("enter confd cli successfully")
        return out

    def run_app_cmd(self, cmd):
        logger.info(f"[APP CMD] {cmd}")
        return self.send_and_wait(cmd, ">> ")
    def run_glob_cmd(self, cmd):
        logger.info(f"[GLOB CMD] {cmd}")
        return self.send_and_wait(cmd, "---> ")

    def run_console_cmd(self, cmd):
        logger.info(f"[CONSOLE CMD] {cmd}")
        return self.send_and_wait(cmd, "~ #")

    def run_confd_cli_cmd(self, cmd):
        logger.info(f"[CONFD CLI CMD] {cmd}")
        return self.send_and_wait(cmd, "#")
