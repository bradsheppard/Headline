from abc import ABC, abstractmethod
from collections.abc import Iterator
from typing import List
from proto.collection import collection_pb2

from collector.messaging.producer import Producer


class Consumer(ABC):

    @property
    @abstractmethod
    def consumer(self) -> Iterator:
        pass

    def __init__(self) -> None:
        self._producer = None

    def attach(self, producer: Producer):
        self._producer = producer

    def __next__(self):
        message = next(self.consumer)
        message_val = message.value

        collection = collection_pb2.Collection.FromString(message_val)

        self.notify(collection.userId, list(collection.interests))

        return collection

    @abstractmethod
    def notify(self, user_id: int, interests: List[str]):
        pass
