from typing import List

from collector.article.article_collector import ArticleCollector
from collector.article.article_service import ArticleService

class Producer:

    def __init__(self, article_service: ArticleService, article_collector: ArticleCollector):
        self._article_service = article_service
        self._article_collector = article_collector

    def update(self, interests: List[str], user_id: int):
        articles = []

        for interest in interests:
            collected_articles = self._article_collector.collect_articles(interest)
            articles.extend(collected_articles)

        user_articles = self._article_service.get_user_articles(user_id)
        unrelated_articles = list(
                filter(lambda article: article.interest not in interests, user_articles))

        articles.extend(unrelated_articles)

        self._article_service.set_user_articles(user_id, articles)
