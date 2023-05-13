import os

from collector.messaging.ddg_subscriber import DDGSubscriber
from collector.messaging.message_bus import MessageBus


if __name__ == '__main__':
    host = os.environ['KAFKA_HOST']
    listener_topic = os.environ['KAFKA_LISTENER_TOPIC']
    publisher_topic = os.environ['KAFKA_PUBLISHER_TOPIC']

    message_bus = MessageBus(host, listener_topic)
    ddg_subscriber = DDGSubscriber(host, publisher_topic)

    message_bus.attach(ddg_subscriber)

    message_bus.consume()
