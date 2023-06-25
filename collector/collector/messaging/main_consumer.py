from collections.abc import Iterator
from typing import List
from kafka import KafkaConsumer
from collector.messaging.consumer import Consumer


class MainConsumer(Consumer):

    @property
    def consumer(self) -> Iterator:
        return self._consumer

    def __init__(self, host: str, topic: str):
        super().__init__()
        self._consumer = KafkaConsumer(topic, bootstrap_servers=host, group_id='collector')
        self._consumer.poll()

    def notify(self, topics: List[str]):
        if not self._producer:
            return

        self._producer.update(topics)
