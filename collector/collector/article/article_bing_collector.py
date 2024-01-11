from typing import List

import requests
from collector.article.article_collector import ArticleCollector
from proto.article import article_pb2


class ArticleBingCollector(ArticleCollector):
    _url = "https://api.bing.microsoft.com/v7.0/news"

    def __init__(self, api_key: str):
        self._api_key = api_key

    def collect_articles(self, topic: str) -> List[article_pb2.Article]:
        articles = []

        querystring = {
                'textFormat': 'Raw', 
                'q': topic,
                'originalImg': True,
                'freshness': 'Day'
        }

        headers = {
            "Ocp-Apim-Subscription-Key": self._api_key,
        }

        response = requests.get(self._url + '/search', headers=headers, params=querystring, timeout=5)
        response_json = response.json()

        if "value" not in response_json:
            raise Exception("Value not in response: ", response_json)

        entries = response_json["value"]

        for entry in entries:
            article = article_pb2.Article(
                    description=entry["description"],
                    imageUrl=entry["image"]["contentUrl"] \
                        if "image" in entry and "contentUrl" in entry["image"] else None,
                    title=entry["name"],
                    source=entry["provider"][0]["name"] if len(entry["provider"]) > 0 else None,
                    url=entry["url"]
            )

            articles.append(article)

        return articles

    def collect_trending_articles(self) -> List[article_pb2.Article]:
        articles = []

        querystring = {
                'textFormat': 'Raw', 
                'safeSearch': 'Off',
                'originalImg': True,
                'freshness': 'Day'
        }

        headers = {
            "Ocp-Apim-Subscription-Key": self._api_key,
        }

        response = requests.get(self._url, headers=headers, params=querystring, timeout=5)
        response_json = response.json()

        if "value" not in response_json:
            raise Exception("Value not in response: ", response_json)

        entries = response_json["value"]

        for entry in entries:
            article = article_pb2.Article(
                    description=entry["description"],
                    imageUrl=entry["image"]["contentUrl"] \
                        if "image" in entry and "contentUrl" in entry["image"] else None,
                    title=entry["name"],
                    source=entry["provider"][0]["name"] if len(entry["provider"]) > 0 else None,
                    url=entry["url"]
            )

            articles.append(article)

        return articles
