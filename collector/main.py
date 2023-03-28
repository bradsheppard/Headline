import os
import grpc

from duckduckgo_search import ddg_news
from model.result import Result
from article import article_pb2_grpc, article_pb2
from interest import interest_pb2_grpc, interest_pb2

backend_service_url = os.environ['BACKEND_URL']

with grpc.insecure_channel(backend_service_url) as channel:
    stub = interest_pb2_grpc.InterestServiceStub(channel)
    
    request = interest_pb2.GetInterestsRequest(
            userId=1
    )

    interest_response: interest_pb2.InterestResponse = stub.GetInterests(request)
    interests = interest_response.interests


for interest in interests:
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
                summary=response['body'],
                link=response['url'],
                userId=1
            )

            articles.append(article)
            
        stub = article_pb2_grpc.ArticleServiceStub(channel)

        request = article_pb2.CreateArticleRequest(
                articles=articles
        )

        stub.CreateArticle(request)

