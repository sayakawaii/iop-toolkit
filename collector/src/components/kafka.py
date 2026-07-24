from kafka import KafkaConsumer, KafkaProducer
import json
import os
from typing import Optional

DEFAULT_KAFKA_BOOTSTRAP = os.environ.get("KAFKA_BOOTSTRAP", "localhost:9094")
TOPIC_REQUEST = os.environ.get("KAFKA_TOPIC_REQUEST", "collectorRequest")
TOPIC_RESPONSE = os.environ.get("KAFKA_TOPIC_RESPONSE", "collectorResponse")
GROUP_REQUEST = os.environ.get("KAFKA_GROUP_REQUEST", "group-collector-request")

class KafkaHandler:
    def __init__(self, bootstrap_servers: Optional[str] = None):
        self.bootstrap_servers = bootstrap_servers or DEFAULT_KAFKA_BOOTSTRAP
        self.topic_response = TOPIC_RESPONSE
        self.consumer = KafkaConsumer(
            TOPIC_REQUEST,
            bootstrap_servers=self.bootstrap_servers,
            value_deserializer=lambda m: json.loads(m.decode('utf-8')),
            key_deserializer=lambda m: m.decode('utf-8') if m else None,
            auto_offset_reset='earliest',
            enable_auto_commit=True,
            group_id=GROUP_REQUEST
        )
        self.producer = KafkaProducer(
            bootstrap_servers=self.bootstrap_servers,
            value_serializer=lambda v: json.dumps(v).encode('utf-8'),
            key_serializer=lambda k: k.encode('utf-8') if k else None,
        )

    def consume_message(self) -> Optional[any]:
        for message in self.consumer:
            yield message

    def send_response(self, key: str, value: dict):
        self.producer.send(self.topic_response, key=key, value=value)
        self.producer.flush()

    def close(self):
        self.consumer.close()
        self.producer.close()
