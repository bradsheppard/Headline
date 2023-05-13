from kafka import KafkaConsumer
from collector.article import article_pb2
from collector.messaging.ddg_subscriber import DDGSubscriber


def test_update():
    host = 'queue-kafka-bootstrap:9092'
    topic = 'test-topic'

    consumer = KafkaConsumer(topic, bootstrap_servers=host, group_id='test')
    consumer.poll()

    subscriber = DDGSubscriber(host, topic)

    subscriber.update('Metallica')
    message = next(consumer).value

    # pylint: disable=no-member
    user_articles = article_pb2.UserArticles.FromString(message)

    for article in user_articles.articles:
        assert article.description
        assert article.interest
        assert article.source
        assert article.title
        assert article.url
