from typing import List

import requests

from collector.article.article_collector import ArticleCollector
from proto.article import article_pb2


class ArticleContextualWebCollector(ArticleCollector):

    _url = "https://contextualwebsearch-websearch-v1.p.rapidapi.com/api/search/NewsSearchAPI"

    def __init__(self, api_key: str) -> None:
        self._api_key = api_key

    def collect_articles(self, interest: str) -> List[article_pb2.Article]:
        articles = []
        querystring = {
                "q": interest,
                "pageNumber":"1",
                "pageSize":"10",
                "autoCorrect":"true",
                "fromPublishedDate":"null",
                "toPublishedDate":"null"
        }

        headers = {
                "X-RapidAPI-Key": self._api_key,
                "X-RapidAPI-Host": "contextualwebsearch-websearch-v1.p.rapidapi.com"
        }

        response = requests.get(self._url, headers=headers, params=querystring, timeout=5)
        response_json = response.json()
        entries = response_json["value"]

        for entry in entries:
            article = article_pb2.Article(
                    description=entry["description"],
                    imageUrl=entry["image"]["url"],
                    title=entry["title"],
                    interest=interest,
                    source=entry["provider"]["name"],
                    url=entry["url"]
            )

            articles.append(article)

        return articles
