from unittest.mock import MagicMock
from collector.article.article_service import ArticleService
from collector.messaging.ddg_producer import DDGProducer


def test_update():
    mock_article_service = ArticleService()
    mock_article_service.set_user_articles = MagicMock()

    subscriber = DDGProducer(mock_article_service)

    subscriber.update(['Metallica'], 1)

    mock_article_service.set_user_articles.assert_called_once()
