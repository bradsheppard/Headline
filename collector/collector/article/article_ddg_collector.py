from typing import List
from duckduckgo_search import DDGS
from collector.article.article_collector import ArticleCollector
from proto.article import article_pb2


class ArticleDDGCollector(ArticleCollector):

    def collect_articles(self, topic: str) -> List[article_pb2.Article]:
        articles = []

        with DDGS() as ddgs:
            responses = ddgs.news(topic)

            for response in responses:
                article = article_pb2.Article(
                    title=response['title'],
                    description=response['body'],
                    url=response['url'],
                    imageUrl=response['image'],
                    source=response['source'],
                )

                articles.append(article)

        return articles
