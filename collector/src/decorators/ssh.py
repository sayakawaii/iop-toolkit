from components.connection import SSHConnection

def with_ssh(func):
    def wrapper(self, *args, **kwargs):
        if hasattr(self, "_ssh") and self._ssh:
            ssh = self._ssh
        else:
            conn = SSHConnection(self.target, self.jump)
            ssh = conn.connect()
            if ssh is None:
                raise RuntimeError("SSH connection failed")
            self._ssh = ssh
            self._conn = conn
        try:
            return func(self, ssh, *args, **kwargs)
        finally:
            # conn.close()
            pass
    return wrapper