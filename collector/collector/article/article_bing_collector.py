from typing import List

import requests
from collector.article.article_collector import ArticleCollector
from proto.article import article_pb2


class ArticleBingCollector(ArticleCollector):
    _url = "https://bing-news-search1.p.rapidapi.com/news/search"

    def __init__(self, api_key: str):
        self._api_key = api_key

    def collect_articles(self, topic: str) -> List[article_pb2.Article]:
        articles = []

        querystring = {
                'textFormat': 'Raw', 
                'safeSearch': 'Off',
                'q': topic,
                'originalImg': True,
                'freshness': 'Day'
        }

        headers = {
            "X-BingApis-SDK": "true",
            "X-RapidAPI-Key": self._api_key,
            "X-RapidAPI-Host": "bing-news-search1.p.rapidapi.com"
        }

        response = requests.get(self._url, headers=headers, params=querystring, timeout=5)
        response_json = response.json()

        entries = response_json["value"]

        for entry in entries:
            article = article_pb2.Article(
                    description=entry["description"],
                    imageUrl=entry["image"]["contentUrl"] \
                        if "image" in entry else None,
                    title=entry["name"],
                    source=entry["provider"][0]["name"] if len(entry["provider"]) > 0 else None,
                    url=entry["url"]
            )

            articles.append(article)

        return articles
