from collections.abc import Iterator
from kafka import KafkaConsumer
from collector.messaging.trending.trending_consumer import TrendingConsumer


class KafkaTrendingConsumer(TrendingConsumer):

    @property
    def consumer(self) -> Iterator:
        return self._consumer

    def __init__(self, host: str, topic: str):
        super().__init__()
        self._consumer = KafkaConsumer(topic, bootstrap_servers=host, group_id='collector')
        self._consumer.poll()

    def notify(self):
        if not self._producer:
            return

        self._producer.update()
