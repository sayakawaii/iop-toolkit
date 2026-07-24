from models.host import OltHost
import paramiko
import socket
import logging

logger = logging.getLogger(__name__)

class SSHConnection:
    def __init__(self, host: OltHost, jump: OltHost | None = None):
        self.host = host
        self.jump = jump
        self.ssh: paramiko.SSHClient | None = None

        self._transport: paramiko.Transport | None = None
        self._jump_transport: paramiko.Transport | None = None
        self._jump_client: paramiko.SSHClient | None = None

    def connect(self):
        try:
            if not self.jump:
                logger.info("not jump")
                self._transport = self._open_transport(
                    sock=self._direct_socket()
                )
            else:
                logger.info("jump")
                self._transport = self._open_transport(
                    sock=self._jump_socket()
                )

            self.ssh = paramiko.SSHClient()
            self.ssh._transport = self._transport
            return self.ssh

        except paramiko.AuthenticationException:
            logger.error("Authentication failed")
        except paramiko.SSHException as e:
            logger.error(f"SSH error: {e}")
        except Exception:
            logger.exception("Unexpected SSH error")

        return None

    def _open_transport(self, sock) -> paramiko.Transport:
        transport = paramiko.Transport(sock)
        self._allow_legacy_rsa(transport)

        transport.start_client()
        transport.auth_password(self.host.user, self.host.pwd)
        return transport

    def _direct_socket(self):
        return socket.create_connection(
            (self.host.host, self.host.port),
            timeout=30
        )

    def _jump_socket(self):
        jump_sock = socket.create_connection(
            (self.jump.host, self.jump.port),
            timeout=30
        )

        self._jump_transport = paramiko.Transport(jump_sock)
        self._allow_legacy_rsa(self._jump_transport)

        self._jump_transport.start_client()
        self._jump_transport.auth_password(
            self.jump.user,
            self.jump.pwd
        )

        chan = self._jump_transport.open_channel(
            "direct-tcpip",
            (self.host.host, self.host.port),
            ("127.0.0.1", 0)
        )

        return chan

    @staticmethod
    def _allow_legacy_rsa(transport: paramiko.Transport):
        sec = transport.get_security_options()
        sec.key_types = [
            "ssh-rsa",
            "rsa-sha2-256",
            "rsa-sha2-512",
        ]

    def close(self):
        if self.ssh:
            self.ssh.close()
            self.ssh = None

        if self._transport:
            self._transport.close()
            self._transport = None

        if self._jump_transport:
            self._jump_transport.close()
            self._jump_transport = None
