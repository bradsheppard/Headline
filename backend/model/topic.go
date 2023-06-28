package model

import (
	pb "headline/proto/topic"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type Topic struct {
	CreatedAt time.Time
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"primaryKey"`
	Articles  []Article      `gorm:"foreignKey:TopicName;references:Name"`
	Users     []User         `gorm:"many2many:user_topics;"`
}

func ToTopicProtos(topics []*Topic) []*pb.Topic {
	var protoTopics []*pb.Topic

	for _, topic := range topics {
		protoTopic := &pb.Topic{
			Name:        topic.Name,
			LastUpdated: timestamppb.New(topic.UpdatedAt),
		}
		protoTopics = append(protoTopics, protoTopic)
	}

	return protoTopics
}

func FromTopicProtos(protoTopics []*pb.Topic) []*Topic {
	var topics []*Topic

	for _, protoTopic := range protoTopics {
		topic := &Topic{
			Name:      protoTopic.Name,
			UpdatedAt: protoTopic.LastUpdated.AsTime(),
		}
		topics = append(topics, topic)
	}

	return topics
}
