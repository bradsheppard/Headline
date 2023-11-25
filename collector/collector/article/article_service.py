from typing import Dict, List

from proto.article import article_pb2


class ArticleService:

    def set_topic_articles(self, topic_articles: Dict[str, List[article_pb2.Article]]):
        pass

    def get_topic_articles(self, topic_name: str) -> List[article_pb2.Article]:
        pass

    def set_trending_articles(self, trending_articles: List[article_pb2.Article]):
        pass

    def get_treding_articles(self) -> List[article_pb2.Article]:
        pass
