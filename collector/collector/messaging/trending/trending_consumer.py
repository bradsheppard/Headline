from abc import ABC, abstractmethod
from collections.abc import Iterator

from collector.messaging.trending.trending_producer import TrendingProducer


class TrendingConsumer(ABC):

    @property
    @abstractmethod
    def consumer(self) -> Iterator:
        pass

    def __init__(self) -> None:
        self._producer = None

    def attach(self, producer: TrendingProducer):
        self._producer = producer

    def __next__(self):
        next(self.consumer)
        self.notify()

    @abstractmethod
    def notify(self):
        pass
