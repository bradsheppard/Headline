import os
from collector.article.article_grpc_service import ArticleGrpcService

from collector.messaging.ddg_producer import DDGProducer
from collector.messaging.main_consumer import MainConsumer


if __name__ == '__main__':
    host = os.environ['KAFKA_HOST']
    listener_topic = os.environ['KAFKA_TOPIC']
    backend_url = os.environ['BACKEND_URL']

    article_service = ArticleGrpcService(backend_url)
    ddg_producer = DDGProducer(article_service)
    main_consumer = MainConsumer(host, listener_topic)

    main_consumer.attach(ddg_producer)

    while True:
        print('Starting consumption...')

        result = next(main_consumer)

        print(f'Collection run complete for {result.userId}')
