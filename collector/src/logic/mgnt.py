from components.kafka import KafkaHandler
from models.msg import CollectorRequest
from components.netconf import NetconfClient
from components.oltinfo import OltInfoManager
from models.host import Netconf
from dataclasses import asdict
from logic.handler import BaseHandler
import logging

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s][%(name)s] %(message)s",
    datefmt="%Y-%m-%d %H:%M:%S"
)
logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)

class ManagementHandler(BaseHandler):
    def __init__(self, kafka_handler: KafkaHandler):
        super().__init__(max_workers=8)
        self.kafka_handler = kafka_handler

    def _parse_olt_request(self, key: str, value: dict):
        """Parse OLT collector request from message"""
        if 'request_id' in value and 'action' in value and 'oam_ip' in value and 'username' in value and 'password' in value:
            return CollectorRequest(
                request_id=value['request_id'],
                action=value['action'],
                oam_ip=value['oam_ip'],
                username=value['username'],
                password=value['password']
            )
        return None

    def handle(self, key, value):
        request = self._parse_olt_request(key, value)
        if request:
            self.process_request(request)

    def process_request(self, request: CollectorRequest):
        try:
            logger.info(f"Processing request for {request.oam_ip}")
            # Create Netconf connection with context manager to ensure proper cleanup
            # Use longer timeout (120s) for topology queries that may take time
            netconf = Netconf(request.oam_ip, 832, request.username, request.password, timeout=30)
            
            with NetconfClient(netconf) as netconf_client:
                # Get OLT topology
                olt_info_manager = OltInfoManager(netconf_client)
                topology = olt_info_manager.get_olt_topology(request.action)

                # Send response to Kafka
                key = f"olt:{request.oam_ip}"
                value = {"request_id": request.request_id, "ntinfo": asdict(topology)}
                self.kafka_handler.send_response(key, value)
                logger.info(f"value sent to Kafka: {value}")
                logger.info(f"Processed request for {request.oam_ip} and sent response to collectorResponse topic")

        except Exception as e:
            logger.error(f"Error processing request for {request.oam_ip}: {e}", exc_info=True)
            # Optionally send error response
            key = f"olt:{request.oam_ip}"
            value = {"request_id": request.request_id, "error": str(e)}
            self.kafka_handler.send_response(key, value)
