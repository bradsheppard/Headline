from typing import List

from collector.article.article_collector import ArticleCollector
from collector.article.article_service import ArticleService

class TopicProducer:

    def __init__(self, article_service: ArticleService, article_collector: ArticleCollector):
        self._article_service = article_service
        self._article_collector = article_collector

    def update(self, topics: List[str]):
        topic_articles = {}

        for topic in topics:
            articles = self._article_collector.collect_articles(topic)
            topic_articles[topic] = articles

        self._article_service.set_topic_articles(topic_articles)
