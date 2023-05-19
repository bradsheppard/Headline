from typing import List
from duckduckgo_search import DDGS
from kafka import KafkaProducer

from collector.article import article_pb2
from collector.messaging.producer import Producer

class DDGProducer(Producer):

    def __init__(self, host: str, topic: str):
        self._kafka_producer = KafkaProducer(bootstrap_servers=host)
        self._topic = topic
        self._ddgs = DDGS()

    def update(self, interests: List[str], user_id: int):
        articles = []

        for interest in interests:
            responses = self._ddgs.news(interest)

            if responses is None:
                continue

            for response in responses:
                article = article_pb2.Article(
                    title=response['title'],
                    description=response['body'],
                    url=response['url'],
                    imageUrl=response['image'],
                    source=response['source'],
                    interest=interest
                )

                articles.append(article)

        request = article_pb2.UserArticles(
            articles=articles,
            userId=user_id
        ).SerializeToString()

        future = self._kafka_producer.send(self._topic, request)
        future.get(timeout=10)
