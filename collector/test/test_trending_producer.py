from unittest.mock import MagicMock
from collector.article.article_collector import ArticleCollector
from collector.article.article_service import ArticleService
from collector.messaging.trending.trending_producer import TrendingProducer
from proto.article import article_pb2


def test_update():
    mock_article_service = ArticleService()
    mock_article_service.set_trending_articles = MagicMock()

    mock_article_collector = ArticleCollector()

    new_article = article_pb2.Article(
            description='New Description',
            imageUrl='testimg.jpg',
            source='www.wikipedia.com',
            title='New title',
            url='www.testurl.com'
    )

    mock_article_collector.collect_trending_articles = MagicMock(return_value=[new_article])

    subscriber = TrendingProducer(mock_article_service, mock_article_collector)
    subscriber.update()

    expected = [new_article]

    mock_article_service.set_trending_articles.assert_called_once_with(expected)
