from abc import ABC, abstractmethod
from collections.abc import Iterator
from typing import List
from collector.collection import collection_pb2

from collector.messaging.topics.topic_producer import TopicProducer


class TopicConsumer(ABC):

    @property
    @abstractmethod
    def consumer(self) -> Iterator:
        pass

    def __init__(self) -> None:
        self._producer = None

    def attach(self, producer: TopicProducer):
        self._producer = producer

    def __next__(self):
        message = next(self.consumer)
        message_val = message.value

        collection = collection_pb2.Collection.FromString(message_val)

        self.notify(list(collection.topics))

        return collection

    @abstractmethod
    def notify(self, topics: List[str]):
        pass
