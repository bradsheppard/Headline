package api

import (
	"context"
	"headline/model"
	"log"

	pb "headline/proto/article"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
        errFormatString = "Error querying database: %v"
        errString = "Error querying database"
)

type ArticleServer struct {
	pb.UnimplementedArticleServiceServer
}

func (s *ArticleServer) GetArticles(ctx context.Context, in *pb.GetArticlesRequest) (*pb.ArticleResponse, error) {
        log.Printf("Received GetArticles")

        var articles []model.Article
        if err := db.Where(&model.Article{UserID: int(in.UserId)}).Find(&articles).Error; err != nil {
                log.Printf(errFormatString, err)
                return nil, status.Error(codes.Internal, errString)
        }

        grpcArticles := model.ToArticleProtos(articles)

	return &pb.ArticleResponse{
                Articles: grpcArticles,
        }, nil
}

func (s *ArticleServer) CreateArticle(ctx context.Context, in *pb.CreateArticleRequest) (*empty.Empty, error) {
        articles := model.FromArticleProtos(in.Articles)

        if err := db.Create(&articles).Error; err != nil {
                log.Printf(errFormatString, err)
                return nil, status.Error(codes.Internal, errString)
        }

        return &emptypb.Empty{}, nil
}

