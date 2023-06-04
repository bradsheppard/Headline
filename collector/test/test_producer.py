from unittest.mock import MagicMock
from collector.article.article_collector import ArticleCollector
from collector.article.article_service import ArticleService
from collector.messaging.producer import Producer
from proto.article import article_pb2


def test_update():
    unrelated_article = article_pb2.Article(
            description='Unrelated Description',
            imageUrl='testimg.jpg',
            interest='Megadeth',
            source='www.wikipedia.com',
            title='Unrelated title',
            url='www.testurl.com'
    )
    related_article = article_pb2.Article(
            description='Related Description',
            imageUrl='testimg.jpg',
            interest='Metallica',
            source='www.wikipedia.com',
            title='Related title',
            url='www.testurl.com'
    )

    mock_article_service = ArticleService()
    mock_article_service.set_user_articles = MagicMock()
    mock_article_service.get_user_articles = MagicMock(
            return_value=[unrelated_article, related_article])

    mock_article_collector = ArticleCollector()

    new_article = article_pb2.Article(
            description='New Description',
            imageUrl='testimg.jpg',
            interest='Metallica',
            source='www.wikipedia.com',
            title='New title',
            url='www.testurl.com'
    )

    mock_article_collector.collect_articles = MagicMock(return_value=[new_article])

    subscriber = Producer(mock_article_service, mock_article_collector)

    subscriber.update(['Metallica'], 1)

    mock_article_service.set_user_articles.assert_called_once_with(
            1, [new_article, unrelated_article])
