from typing import Dict, List
import grpc
import google.protobuf.empty_pb2

from proto.article import article_pb2, article_pb2_grpc

from collector.article.article_service import ArticleService


class ArticleGrpcService(ArticleService):

    def __init__(self, backend_url: str) -> None:
        super().__init__()
        self._backend_url = backend_url

    def set_topic_articles(self, topic_articles: Dict[str, List[article_pb2.Article]]):
        topic_articles_proto = {}

        for topic, articles in topic_articles.items():
            topic_articles_proto[topic] = article_pb2.Articles(
                    articles=articles
            )

        request = article_pb2.TopicArticles(
                topicArticles=topic_articles_proto
        )

        with grpc.insecure_channel(self._backend_url) as channel:
            stub = article_pb2_grpc.ArticleServiceStub(channel)
            stub.SetTopicArticles(request)

    def get_topic_articles(self, topic_name: str) -> List[article_pb2.Article]:
        request = article_pb2.TopicNames(
                topics=[topic_name]
        )

        with grpc.insecure_channel(self._backend_url) as channel:
            stub = article_pb2_grpc.ArticleServiceStub(channel)
            topic_articles: article_pb2.TopicArticles = stub.GetTopicArticles(request)

            return list(topic_articles.topicArticles[topic_name].articles)

    def set_trending_articles(self, trending_articles: List[article_pb2.Article]):
        request = article_pb2.Articles(
                articles=trending_articles
        )

        with grpc.insecure_channel(self._backend_url) as channel:
            stub = article_pb2_grpc.ArticleServiceStub(channel)
            stub.SetTrendingArticles(request)

    def get_trending_articles(self):

        with grpc.insecure_channel(self._backend_url) as channel:
            stub = article_pb2_grpc.ArticleServiceStub(channel)
            articles: article_pb2.Articles = stub.GetTrendingArticles(google.protobuf.empty_pb2)

            return list(articles.articles)
