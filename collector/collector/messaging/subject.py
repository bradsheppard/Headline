from abc import ABC, abstractmethod
from messaging.subscriber import Subscriber


class Subject(ABC):

    @abstractmethod
    def attach(self, subscriber: Subscriber) -> None:
        pass

    @abstractmethod
    def detach(self, subscriber: Subscriber) -> None:
        pass

    @abstractmethod
    def notify(self, message: str) -> None:
        pass
