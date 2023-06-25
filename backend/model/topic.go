package model

import (
	pb "headline/proto/topic"
	"time"

	"gorm.io/gorm"
)

type Topic struct {
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Name            string `gorm:"primaryKey"`
        Articles        []Article `gorm:"foreignKey:TopicName;references:Name"`
        Users           []User `gorm:"many2many:user_topics;"`
}

func ToTopicProtos(topics []*Topic) []*pb.Topic {
	var protoTopics []*pb.Topic

	for _, topic := range topics {
		protoTopic := &pb.Topic{
			Name:   topic.Name,
		}
		protoTopics = append(protoTopics, protoTopic)
	}

	return protoTopics
}

func FromTopicProtos(protoTopics []*pb.Topic) []*Topic {
	var topics []*Topic

	for _, protoTopic := range protoTopics {
		topic := &Topic{
			Name:   protoTopic.Name,
		}
		topics = append(topics, topic)
	}

	return topics
}
