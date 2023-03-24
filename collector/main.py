import grpc

from duckduckgo_search import ddg_news
from model.result import Result
from article import article_pb2_grpc, article_pb2

interests = ['Metallica', 'Donald Trump']


for interest in interests:
    responses = ddg_news(interest)

    results = []

    if responses is None:
        continue

    with grpc.insecure_channel('localhost:50051') as channel:
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

