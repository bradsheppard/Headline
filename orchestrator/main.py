import os
import datetime

from google.protobuf.timestamp_pb2 import Timestamp

import grpc
from kafka import KafkaProducer

from topic import topic_pb2, topic_pb2_grpc
from collection import collection_pb2, collection_pb2_grpc


if __name__ == '__main__':
    backend_url = os.environ['BACKEND_URL']
    kafka_host = os.environ['KAFKA_HOST']
    kafka_topic_topic = os.environ['KAFKA_TOPIC_TOPIC']
    kafka_trending_topic = os.environ['KAFKA_TRENDING_TOPIC']

    hour_offset = int(os.environ['HOUR_OFFSET'])

    search_time = datetime.datetime.now() + datetime.timedelta(hours=hour_offset)
    search_time_proto = Timestamp()
    search_time_proto.FromDatetime(dt=search_time)

    kafka_producer = KafkaProducer(bootstrap_servers=kafka_host)

    with grpc.insecure_channel(backend_url) as channel:
        stub = topic_pb2_grpc.TopicServiceStub(channel)

        request = topic_pb2.GetPendingTopicsRequest(
                lastUpdated=search_time_proto
        )

        print('Requesting pending topics...')

        topicResponse: topic_pb2.TopicResponse
        for topicResponse in stub.GetPendingTopics(request):
            topics = []

            for topic in topicResponse.topics:
                topics.append(topic.name)

            print(f'Orchestrating topics: {topics}')
            collection = collection_pb2.Collection(
                    topics=topics
            )

            message = collection.SerializeToString()

            future = kafka_producer.send(kafka_topic_topic, value=message)
            future.get(timeout=60)

        print('Orchestrating trending articles')

        future = kafka_producer.send(kafka_trending_topic, value=b'N/A')
        future.get(timeout=60)

