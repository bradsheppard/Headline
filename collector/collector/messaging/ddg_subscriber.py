from duckduckgo_search import ddg_news
from kafka import KafkaProducer

from collector.article import article_pb2
from collector.messaging.subscriber import Subscriber

class DDGSubscriber(Subscriber):

    def __init__(self, host: str, topic: str):
        self._kafka_producer = KafkaProducer(bootstrap_servers=host)
        self._topic = topic

    def update(self, message_val: str):
        responses = ddg_news(message_val)

        articles = []

        if responses is None:
            return

        for response in responses:
            # pylint: disable=no-member
            article = article_pb2.Article(
                title=response['title'],
                description=response['body'],
                url=response['url'],
                imageUrl=response['image'],
                source=response['source'],
                interest=message_val
            )

            articles.append(article)

        # pylint: disable=no-member
        request = article_pb2.UserArticles(
                articles=articles,
                userId=1
        ).SerializeToString()

        future = self._kafka_producer.send(self._topic, request)
        future.get(timeout=10)
