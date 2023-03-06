package api

import (
	"context"
	"log"
        "headline/model"

	pb "headline/proto"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedArticleServiceServer
}

func (s *server) GetArticles(ctx context.Context, in *empty.Empty) (*pb.ArticleResponse, error) {
	log.Printf("Received a request")

        var articles []model.Article
        db.Find(&articles)

        var grpcArticles []*pb.Article
        for i := range(articles) {
                grpcArticle := model.ToProto(articles[i])
                grpcArticles = append(grpcArticles, grpcArticle)
        }

	return &pb.ArticleResponse{
                Articles: grpcArticles,
        }, nil
}

func (s *server) CreateArticle(ctx context.Context, in *pb.CreateArticleRequest) (*empty.Empty, error) {
        log.Printf("Creating article")

        var articles []model.Article
        for i := range(in.Articles) {
                article := model.FromProto(in.Articles[i])
                articles = append(articles, *article)
        }

        db.Create(&articles)

        return &emptypb.Empty{}, nil
}

