from logic.mgnt import ManagementHandler
from logic.logger import LoggerHandler

class HandlerRouter:
    def __init__(self, kafka_handler):
        self.mgmt_handler = ManagementHandler(kafka_handler)
        self.logger_handler = LoggerHandler(kafka_handler)

    def dispatch(self, key, value):
        msg_type = value.get("msg_type")
        if msg_type == "connect_request":
            self.mgmt_handler.submit(key, value)
        elif msg_type == "logger_settings":
            self.logger_handler.submit(key, value)

    def shutdown(self):
        self.mgmt_handler.shutdown()
