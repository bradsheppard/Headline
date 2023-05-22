from kafka import KafkaConsumer, TopicPartition
from collector.article import article_pb2
from collector.messaging.ddg_producer import DDGProducer


def test_update():
    host = 'queue-kafka-bootstrap:9092'
    topic = 'test-topic'

    consumer = KafkaConsumer(topic, bootstrap_servers=host, group_id='test')
    consumer.poll()

    topic_partition = TopicPartition(topic, 0)

    offset = consumer.position(topic_partition)
    consumer.seek(topic_partition, max([offset - 2, 0]))

    subscriber = DDGProducer(host, topic)

    subscriber.update(['Metallica'], 1)
    message = next(consumer).value

    # pylint: disable=no-member
    user_articles = article_pb2.UserArticles.FromString(message)
    assert len(user_articles.articles) > 0

    for article in user_articles.articles:
        assert article.description
        assert article.interest
        assert article.source
        assert article.title
        assert article.url
