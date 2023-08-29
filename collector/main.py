import os
from collector.article.article_bing_collector import ArticleBingCollector
from collector.article.article_grpc_service import ArticleGrpcService

from collector.messaging.main_consumer import MainConsumer
from collector.messaging.producer import Producer


if __name__ == '__main__':
    host = os.environ['KAFKA_HOST']
    listener_topic = os.environ['KAFKA_TOPIC']
    backend_url = os.environ['BACKEND_URL']
    api_key = os.environ['API_KEY']

    article_service = ArticleGrpcService(backend_url)
    article_collector = ArticleBingCollector(api_key)

    producer = Producer(article_service, article_collector)
    main_consumer = MainConsumer(host, listener_topic)

    main_consumer.attach(producer)

    while True:
        print('Starting consumption...')

        result = next(main_consumer)

        print(f'Collection run complete for {list(result.topics)}')
