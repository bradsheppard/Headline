from unittest.mock import MagicMock

from kafka import KafkaProducer
from collector.collection import collection_pb2
from collector.messaging.main_consumer import MainConsumer


def test_notify():
    host = 'queue-kafka-bootstrap:9092'
    topic = 'test-topic'

    main_consumer = MainConsumer(host, topic)
    kafka_producer = KafkaProducer(bootstrap_servers=host)

    main_consumer.notify = MagicMock()

    topics = ['Metallica', 'Software Engineering']

    collection = collection_pb2.Collection(
            topics=topics
    )
    collection_bytes = collection.SerializeToString()

    kafka_producer.send(topic, collection_bytes)
    next(main_consumer)

    main_consumer.notify.assert_called_with(topics)
