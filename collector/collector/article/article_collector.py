from typing import List
from proto.article import article_pb2


class ArticleCollector:

    def collect_articles(self, interest: str) -> List[article_pb2.Article]:
        pass
