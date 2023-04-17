import os
import grpc
import time

from duckduckgo_search import ddg_news
from model.result import Result
from article import article_pb2_grpc, article_pb2
from interest import interest_pb2_grpc, interest_pb2

backend_service_url = os.environ['BACKEND_URL']

print("Starting collection")

with grpc.insecure_channel(backend_service_url) as channel:
    stub = interest_pb2_grpc.InterestServiceStub(channel)
    
    request = interest_pb2.GetInterestsRequest(
            userId=1
    )

    interest_response: interest_pb2.InterestResponse = stub.GetInterests(request)
    interests = interest_response.interests

print(f"Interests: {interests}")

for interest in interests:
    print(f"Crawling for interest: {interest}")
    
    search_term = interest.name
    responses = ddg_news(search_term)

    results = []

    if responses is None:
        continue

    with grpc.insecure_channel(backend_service_url) as channel:
        articles = []

        for response in responses:
            article = article_pb2.Article(
                title=response['title'],
                description=response['body'],
                url=response['url'],
                imageUrl=response['image'],
                source=response['source']
            )

            articles.append(article)
            
        stub = article_pb2_grpc.ArticleServiceStub(channel)

        request = article_pb2.UserArticles(
                articles=articles,
                userId=1
        )

        stub.SetUserArticles(request)

