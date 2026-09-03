from dataclasses import dataclass

@dataclass
class CollectorRequest:
    request_id: str
    action: str
    oam_ip: str
    username: str
    password: str

@dataclass
class CollectorLoggerSettings:
    request_id: str
    oam_ip: str
    action: str
    request_parent: str
    logger_modules: dict