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
	"gorm.io/gorm"
)

const (
        errFormatString = "Error querying database: %v"
        errString = "Error querying database"
)

type ArticleServer struct {
	pb.UnimplementedArticleServiceServer
}

func (s *ArticleServer) GetArticles(ctx context.Context, in *pb.User) (*pb.UserArticles, error) {
        log.Printf("Received GetArticles")

        var articles []model.Article
        if err := db.Where(&model.Article{UserID: int(in.UserId)}).Find(&articles).Error; err != nil {
                log.Printf(errFormatString, err)
                return nil, status.Error(codes.Internal, errString)
        }

        grpcArticles := model.ToArticleProtos(articles)

	return &pb.UserArticles{
                Articles: grpcArticles,
                UserId: in.UserId,
        }, nil
}

func (s *ArticleServer) SetUserArticles(ctx context.Context, in *pb.UserArticles) (*empty.Empty, error) {
        articles := model.FromArticleProtos(in.Articles, in.UserId)

        err := db.Transaction(func(tx *gorm.DB) error {
                if err := db.Where("user_id = ?", in.UserId).Delete(&pb.Article{}).Error; err != nil {
                        return err
                }

                if err := db.Create(&articles).Error; err != nil {
                        return err
                }

                return nil
        })

        if err != nil {
                log.Printf(errFormatString, err)
                return nil, status.Error(codes.Internal, errString)
        }

        return &emptypb.Empty{}, nil
}

