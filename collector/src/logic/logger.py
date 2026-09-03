from components.kafka import KafkaHandler
from models.msg import CollectorLoggerSettings
from models.host import OltHost
from logic.handler import BaseHandler
import logging

from components.collector import LogCollector
from components.logserver import LogServerSupervisor
from utils.env import Platform_Env

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s][%(name)s] %(message)s",
    datefmt="%Y-%m-%d %H:%M:%S"
)
logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)


lt_IP_mapping = {
    "Slot-Lt-1": "169.254.1.3",
    "Slot-Lt-2": "169.254.1.4",
    "Slot-Lt-3": "169.254.1.5",
    "Slot-Lt-4": "169.254.1.6",
    "Slot-Lt-5": "169.254.1.7",
    "Slot-Lt-6": "169.254.1.8",
    "Slot-Lt-7": "169.254.1.9",
    "Slot-Lt-8": "169.254.1.10",
}

def get_ip_by_lt_parent(lt_parent: str) -> str:
    return lt_IP_mapping.get(lt_parent, "169.254.1.3")

class LoggerHandler(BaseHandler):
    def __init__(self, kafka_handler: KafkaHandler):
        super().__init__(max_workers=8)
        self.kafka_handler = kafka_handler
        self.env = Platform_Env
        self.logs = None
        self.logserver_supervisor = None

    def _parse_olt_logger_settings(self, key: str, value: dict):
        if  'request_id' in value and 'oam_ip' in value and 'action' in value and 'request_parent' in value and 'logger_modules' in value:
            return CollectorLoggerSettings(
                request_id=value['request_id'],
                oam_ip=value['oam_ip'],
                action=value['action'],
                request_parent=value['request_parent'],
                logger_modules=value['logger_modules']
            )
        return None

    def handle(self, key, value):
        logger_settings = self._parse_olt_logger_settings(key, value)
        logger.info(f"Received logger_settings for {logger_settings.oam_ip}: {logger_settings.request_parent}: {logger_settings.action}")
        if logger_settings:
            self.process_logger_settings(logger_settings)

    def start_logging_server(self, request: CollectorLoggerSettings):
        jump = OltHost(request.oam_ip, "root", "2x2=4", 923)
        target = OltHost(get_ip_by_lt_parent(request.request_parent), "root", "2x2=4", 2222)
        collector = LogCollector(target, jump)
        self.logserver_supervisor = LogServerSupervisor(collector, self.env.ipaddress)
        self.logs = self.logserver_supervisor.start(request.logger_modules, request.action)

    def stop_logging_server(self, request: CollectorLoggerSettings):
        if self.logserver_supervisor:
            self.logs = self.logserver_supervisor.stop()
            logging.info(f"Log server for {request.oam_ip} stopped successfully")

    def process_logger_settings(self, request: CollectorLoggerSettings):
        try:
            logger.info(f"Processing logger_settings for {request.oam_ip}")
            if request.action == "enable" or request.action == "overlay":
                self.start_logging_server(request)
            else:
                self.stop_logging_server(request)

            # Send response to Kafka
            key = f"olt:{request.oam_ip}"
            value = {"request_id": request.request_id, "logs": self.logs}
            self.kafka_handler.send_response(key, value)

        except Exception as e:
            logger.error(f"Error processing logger_settings for {request.oam_ip}: {e}", exc_info=True)
            # Optionally send error response
            key = f"olt:{request.oam_ip}"
            value = {"request_id": request.request_id, "error": str(e)}
            self.kafka_handler.send_response(key, value)
