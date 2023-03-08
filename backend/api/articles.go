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

func (s *Server) GetArticles(ctx context.Context, in *pb.GetArticlesRequest) (*pb.ArticleResponse, error) {
	log.Printf("Received a request")

        var articles []model.Article
        if err := db.Where(&model.Article{UserID: int(in.UserId)}).Find(&articles).Error; err != nil {
                log.Printf(errFormatString, err)
                return nil, status.Error(codes.Internal, errString)
        }

        grpcArticles := model.ToProtos(articles)

	return &pb.ArticleResponse{
                Articles: grpcArticles,
        }, nil
}

func (s *Server) CreateArticle(ctx context.Context, in *pb.CreateArticleRequest) (*empty.Empty, error) {
        log.Printf("Creating article")

        articles := model.FromProtos(in.Articles)

        if err := db.Create(&articles).Error; err != nil {
                log.Printf(errFormatString, err)
                return nil, status.Error(codes.Internal, errString)
        }

        return &emptypb.Empty{}, nil
}

