from dataclasses import dataclass, asdict

@dataclass
class OltHost:
    host: str
    user: str
    pwd: str
    port: int = 22

@dataclass
class LogServer:
    host: str
    port: int

@dataclass
class Netconf:
    host: str
    port: int = 830
    username: str = "admin"
    password: str = "admin"
    hostkey_verify: bool = False
    timeout: int = 30
    def as_params(self):
        return asdict(self)