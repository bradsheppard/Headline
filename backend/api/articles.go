package api

import (
	"context"
	"headline/model"
	"log"

	"headline/proto/article"
	pb "headline/proto/article"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

const (
	errDatabaseFormatString  = "Error querying database: %v"
	errDatabaseString        = "Error querying database"
	errMessagingFormatString = "Error writing to Kafka brokers: %v"
	errMessagingString       = "Error writing to Kafka brokers"
)

type ArticleServer struct {
	pb.UnimplementedArticleServiceServer
}

type TopicsWithArticles struct {
	topics   []*model.Topic
	articles [][]*model.Article
}

func Unwrap(m map[string]*article.Articles) TopicsWithArticles {
	topicsWithArticles := TopicsWithArticles{
		topics:   []*model.Topic{},
		articles: [][]*model.Article{},
	}
	for topic, articles := range m {
		topicsWithArticles.topics = append(topicsWithArticles.topics, &model.Topic{Name: topic})
		topicsWithArticles.articles = append(topicsWithArticles.articles, model.FromArticleProtos(articles.Articles))
	}
	return topicsWithArticles
}

func (s *ArticleServer) GetTopicArticles(ctx context.Context, in *pb.GetTopicArticlesRequest) (*pb.TopicArticles, error) {
	var articles []model.Article
	topics := []model.Topic{}

	for _, topic := range in.Topics {
		topics = append(topics, model.Topic{Name: topic})
	}

	if err := db.Model(&topics).Association("Articles").Find(&articles); err != nil {
		log.Printf(errDatabaseFormatString, err)
		return nil, status.Error(codes.Internal, errDatabaseString)
	}

	grpcArticles := model.ToConverter(articles)

	return grpcArticles, nil
}

func (s *ArticleServer) SetTopicArticles(ctx context.Context, in *pb.SetTopicArticlesRequest) (*empty.Empty, error) {
	unwrapped := Unwrap(in.TopicArticles)

	articles := []interface{}{}

	for _, article := range unwrapped.articles {
		articles = append(articles, article)
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := db.Unscoped().Model(unwrapped.topics).Association("Articles").Clear(); err != nil {
			return err
		}

		if err := db.Model(unwrapped.topics).Association("Articles").Append(articles...); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Printf(errDatabaseFormatString, err)
		return nil, status.Error(codes.Internal, errDatabaseString)
	}

	return &emptypb.Empty{}, nil
}
