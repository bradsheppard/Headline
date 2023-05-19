import os

from collector.messaging.ddg_producer import DDGProducer
from collector.messaging.main_consumer import MainConsumer


if __name__ == '__main__':
    host = os.environ['KAFKA_HOST']
    listener_topic = os.environ['KAFKA_LISTENER_TOPIC']
    publisher_topic = os.environ['KAFKA_PUBLISHER_TOPIC']

    ddg_producer = DDGProducer(host, publisher_topic)
    main_consumer = MainConsumer(host, listener_topic)

    main_consumer.attach(ddg_producer)

    while True:
        result = next(main_consumer)

        print(f'Collection run for {result.userId}')
