from typing import List
from duckduckgo_search import DDGS

from proto.article import article_pb2
from collector.article.article_service import ArticleService
from collector.messaging.producer import Producer

class DDGProducer(Producer):

    def __init__(self, article_service: ArticleService):
        self._article_service = article_service

    def update(self, interests: List[str], user_id: int):
        articles = []

        for interest in interests:
            with DDGS() as ddgs:
                responses = ddgs.news(interest)

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

        self._article_service.set_user_articles(user_id, articles)
