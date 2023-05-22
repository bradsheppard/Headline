package api

import (
       "github.com/segmentio/kafka-go" 
)

var (
	writer *kafka.Writer
)

func InitMessaging(broker string, topic string) {
        writer = &kafka.Writer{
                Addr: kafka.TCP(broker),
                Topic: topic,
                AllowAutoTopicCreation: true,
        }
}

