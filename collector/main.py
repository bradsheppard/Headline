import os

from concurrent.futures import ThreadPoolExecutor
from concurrent.futures import wait
from concurrent.futures import FIRST_EXCEPTION

from collector.article.article_bing_collector import ArticleBingCollector
from collector.article.article_grpc_service import ArticleGrpcService

from collector.messaging.topics.kafka_topic_consumer import KafkaTopicConsumer
from collector.messaging.topics.topic_consumer import TopicConsumer
from collector.messaging.topics.topic_producer import TopicProducer
from collector.messaging.trending.kafka_trending_consumer import KafkaTrendingConsumer
from collector.messaging.trending.trending_consumer import TrendingConsumer
from collector.messaging.trending.trending_producer import TrendingProducer


def trending_loop(consumer: TrendingConsumer):
    while True:
        print('Starting trending consumption...', flush=True)

        next(consumer)

        print('Collection run complete for trending', flush=True)


def topic_loop(consumer: TopicConsumer):
    while True:
        print('Starting topic consumption...', flush=True)

        result = next(consumer)

        print(f'Collection run complete for {list(result.topics)}', flush=True)



if __name__ == '__main__':
    host = os.environ['KAFKA_HOST']

    topic_topic = os.environ['KAFKA_TOPIC_TOPIC']
    trending_topic = os.environ['KAFKA_TRENDING_TOPIC']

    backend_url = os.environ['BACKEND_URL']
    api_key = os.environ['API_KEY']

    article_service = ArticleGrpcService(backend_url)
    article_collector = ArticleBingCollector(api_key)

    topic_producer = TopicProducer(article_service, article_collector)
    topic_consumer = KafkaTopicConsumer(host, topic_topic)
    topic_consumer.attach(topic_producer)

    trending_producer = TrendingProducer(article_service, article_collector)
    trending_consumer = KafkaTrendingConsumer(host, trending_topic)
    trending_consumer.attach(trending_producer)

    with ThreadPoolExecutor(2) as executor:
        trending_future = executor.submit(trending_loop, trending_consumer)
        topic_future = executor.submit(topic_loop, topic_consumer)

        futures = [trending_future, topic_future]

        done, not_done = wait(futures, return_when=FIRST_EXCEPTION)
