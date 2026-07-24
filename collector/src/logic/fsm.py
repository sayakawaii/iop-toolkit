from components.kafka import KafkaHandler
from logic.router import HandlerRouter
import logging

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s][%(name)s] %(message)s",
    datefmt="%Y-%m-%d %H:%M:%S"
)
logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)

class FiniteStateMachine:
    def __init__(self):
        self.kafka_handler = KafkaHandler()
        self.router = HandlerRouter(self.kafka_handler)

    def _dispatch_message(self, message):
        """Dispatch message based on message type"""
        key = message.key
        value = message.value
        logger.info(f"Consumed message with key: {key}, value: {value}")

        if key:
            self.router.dispatch(key, value)
        #Add more message types for future extensions

    def run(self):
        """Main loop to consume messages and process them"""
        logger.info("Starting Finite State Machine")
        try:
            for message in self.kafka_handler.consume_message():
                self._dispatch_message(message)
        except KeyboardInterrupt:
            logger.info("Stopping FSM")
        finally:
            self.kafka_handler.close()
