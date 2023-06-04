from typing import List
import grpc

from proto.article import article_pb2, article_pb2_grpc

from collector.article.article_service import ArticleService


class ArticleGrpcService(ArticleService):

    def __init__(self, backend_url: str) -> None:
        super().__init__()
        self._backend_url = backend_url

    def set_user_articles(self, user_id: int, articles: List[article_pb2.Article]):
        request = article_pb2.UserArticles(
            articles=articles,
            userId=user_id
        )

        with grpc.insecure_channel(self._backend_url) as channel:
            stub = article_pb2_grpc.ArticleServiceStub(channel)
            stub.SetUserArticles(request)

    def get_user_articles(self, user_id: int) -> List[article_pb2.Article]:
        request = article_pb2.User(
            userId=user_id
        )

        with grpc.insecure_channel(self._backend_url) as channel:
            stub = article_pb2_grpc.ArticleServiceStub(channel)
            user_articles: article_pb2.UserArticles = stub.GetArticles(request)

            return list(user_articles.articles)
