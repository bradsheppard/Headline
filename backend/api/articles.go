package api

import (
	"context"
	"headline/model"
	"log"

	pb "headline/proto"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
        errFormatString = "Error querying database: %v"
        errString = "Error querying database"
)

type Server struct {
	pb.UnimplementedArticleServiceServer
}

func (s *Server) GetArticles(ctx context.Context, in *empty.Empty) (*pb.ArticleResponse, error) {
	log.Printf("Received a request")

        var articles []model.Article
        if err := db.Find(&articles).Error; err != nil {
                log.Printf(errFormatString, err)
                return nil, status.Error(codes.Internal, errString)
        }

        var grpcArticles []*pb.Article
        for i := range(articles) {
                grpcArticle := model.ToProto(articles[i])
                grpcArticles = append(grpcArticles, grpcArticle)
        }

	return &pb.ArticleResponse{
                Articles: grpcArticles,
        }, nil
}

func (s *Server) CreateArticle(ctx context.Context, in *pb.CreateArticleRequest) (*empty.Empty, error) {
        log.Printf("Creating article")

        var articles []model.Article
        for i := range(in.Articles) {
                article := model.FromProto(in.Articles[i])
                articles = append(articles, *article)
        }

        if err := db.Create(&articles).Error; err != nil {
                log.Printf(errFormatString, err)
                return nil, status.Error(codes.Internal, errString)
        }

        return &emptypb.Empty{}, nil
}

