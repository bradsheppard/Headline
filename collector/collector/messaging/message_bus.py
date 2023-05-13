from kafka import KafkaConsumer

from messaging.subject import Subject
from messaging.subscriber import Subscriber


class MessageBus(Subject):

    def __init__(self, host: str, topic: str):
        self._consumer = KafkaConsumer(topic, bootstrap_servers=host, group_id='collector')
        self._subscribers = []

    def attach(self, subscriber):
        self._subscribers.append(subscriber)

    def detach(self, subscriber: Subscriber) -> None:
        self._subscribers.remove(subscriber)

    def notify(self, message: str):
        for subscriber in self._subscribers:
            subscriber.update(message)

    def consume(self):
        for message in self._consumer:
            message_val = message.value

            self.notify(message_val)
