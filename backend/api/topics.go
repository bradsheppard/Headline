package api

import (
	"context"
	"headline/model"
	"log"

	"headline/collection"

	pb "headline/proto/topic"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/golang/protobuf/ptypes/empty"

	"google.golang.org/api/idtoken"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type TopicServer struct {
	pb.UnimplementedTopicServiceServer
}

func repeated(str *string, n int) []interface{} {
	arr := make([]interface{}, n)

	for i := 0; i < n; i++ {
		arr[i] = &model.User{
			Id: *str,
		}
	}

	return arr
}

func getUserId(ctx context.Context) (*string, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, status.Errorf(codes.DataLoss, "Failed to get metadata")
	}

	token, ok := md["id_token"]
	token_string := ""

	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to get ID token")
	}

	for _, e := range token {
		token_string += e
	}

	payload, err := idtoken.Validate(context.Background(), token_string, "")

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to authenticate token")
	}

	user_id, ok := payload.Claims["sub"].(string)

	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "cannot retrieve subject from token")
	}

	return &user_id, nil
}

func (topicServer *TopicServer) AddTopics(ctx context.Context, in *pb.AddTopicsRequest) (*pb.TopicResponse, error) {
	topics := model.FromTopicProtos(in.Topics)

	user_id, err := getUserId(ctx)

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		if err := db.Clauses(clause.OnConflict{UpdateAll: true, Columns: []clause.Column{{Name: "name"}}}).Create(topics).Error; err != nil {
			return err
		}

		if err := db.Model(topics).Association("Users").Append(repeated(user_id, len(topics))...); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Printf(errDatabaseFormatString, err)
		return nil, status.Error(codes.Internal, errDatabaseString)
	}

	topicsToCollect := []string{}

	for _, topic := range topics {
		topicsToCollect = append(topicsToCollect, topic.Name)
	}

	collection := collection.Collection{
		Topics: topicsToCollect,
	}

	err = StartCollection(ctx, &collection)

	if err != nil {
		log.Printf(errMessagingFormatString, err)
		return nil, status.Error(codes.Internal, errMessagingString)
	}

	return &pb.TopicResponse{
		Topics: model.ToTopicProtos(topics),
	}, nil
}

func (topicServer *TopicServer) GetTopics(ctx context.Context, in *empty.Empty) (*pb.TopicResponse, error) {
	var topics []*model.Topic

	user_id, err := getUserId(ctx)

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	if err := db.Model(&model.Topic{}).Joins("join user_topics on topics.name = user_topics.topic_name").Joins("join users on user_topics.user_id = users.id").Where("users.id = ?", user_id).Select("topics.*").Scan(&topics).Error; err != nil {
		log.Printf(errDatabaseFormatString, err)
		return nil, status.Error(codes.Internal, errDatabaseString)
	}

	grpcTopics := model.ToTopicProtos(topics)

	return &pb.TopicResponse{
		Topics: grpcTopics,
	}, nil
}

func (topicServer *TopicServer) RemoveTopics(ctx context.Context, in *pb.RemoveTopicsRequest) (*empty.Empty, error) {
	topics := []*model.Topic{}

	for _, topicName := range in.TopicNames {
		topic := &model.Topic{
			Name: topicName,
		}

		topics = append(topics, topic)
	}

	user_id, err := getUserId(ctx)

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	err = db.Model(topics).Association("Users").Delete(repeated(user_id, len(topics))...)

	if err != nil {
		log.Printf(errMessagingFormatString, err)
		return nil, status.Error(codes.Internal, errMessagingString)
	}

	return &empty.Empty{}, nil
}

func (topicServer *TopicServer) GetPendingTopics(in *pb.GetPendingTopicsRequest, stream pb.TopicService_GetPendingTopicsServer) error {
	topics := []*model.Topic{}

	result := db.Where("updated_at < ?", in.LastUpdated.AsTime()).FindInBatches(&topics, 1000, func(tx *gorm.DB, batch int) error {
		grpcTopics := model.ToTopicProtos(topics)
		stream.Send(&pb.TopicResponse{
			Topics: grpcTopics,
		})

		return nil
	})

	if result.Error != nil {
		log.Printf(errDatabaseFormatString, result.Error)
		return status.Error(codes.Internal, errDatabaseString)
	}

	return nil
}
