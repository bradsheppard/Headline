from collector.article.article_collector import ArticleCollector
from collector.article.article_service import ArticleService

class TrendingProducer:

    def __init__(self, article_service: ArticleService, article_collector: ArticleCollector):
        self._article_service = article_service
        self._article_collector = article_collector

    def update(self):
        articles = self._article_collector.collect_trending_articles()

        self._article_service.set_trending_articles(articles)
