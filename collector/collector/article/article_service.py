from typing import List

from proto.article import article_pb2


class ArticleService:

    def set_user_articles(self, user_id: int, articles: List[article_pb2.Article]):
        pass
