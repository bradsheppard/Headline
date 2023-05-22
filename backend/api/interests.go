package api

import (
	"context"
	"fmt"
	"headline/model"
	"log"

	"headline/collection"

	pb "headline/proto/interest"

	"gorm.io/gorm"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type InterestServer struct {
	pb.UnimplementedInterestServiceServer
}

func (interestServer *InterestServer) AddInterests(ctx context.Context, in *pb.AddInterestsRequest) (*pb.InterestResponse, error) {
	interests := model.FromInterestProtos(in.Interests)

	err := db.Transaction(func(tx *gorm.DB) error {
                if err := db.Create(&interests).Error; err != nil {
                        return err
                }

                if err := db.Where(&model.Interest{UserID: int(in.UserId)}).Find(&interests).Error; err != nil {
                        return err
	        }

                return nil
	})

        interestsToCollect := []string{}

        for _, interest := range interests{
                interestsToCollect = append(interestsToCollect, interest.Name)
        }

        collection := collection.Collection{
                UserId: in.UserId,
                Interests: interestsToCollect,
        }

        bytes, err := proto.Marshal(&collection)

        if err != nil {
                return nil, status.Error(codes.Internal, "Error marshalling bytes")
        }

        err = writer.WriteMessages(ctx,
                kafka.Message{
                        Value: bytes,
                },
        )

        if err != nil {
                fmt.Printf("Error: %v", err)
                return nil, status.Error(codes.Internal, "Error writing to Kafka brokers")
        }

	return &pb.InterestResponse{
		Interests: model.ToInterestProtos(interests),
	}, nil
}

func (interestServer *InterestServer) DeleteInterests(ctx context.Context, in *pb.DeleteInterestsRequest) (*empty.Empty, error) {
	if err := db.Delete(&model.Interest{}, in.Ids).Error; err != nil {
		log.Printf(errFormatString, err)
		return nil, status.Error(codes.Internal, errString)
	}

	return &emptypb.Empty{}, nil
}

func (interestServer *InterestServer) GetInterests(ctx context.Context, in *pb.GetInterestsRequest) (*pb.InterestResponse, error) {
	var interests []*model.Interest
	if err := db.Where(&model.Interest{UserID: int(in.UserId)}).Find(&interests).Error; err != nil {
		log.Printf(errFormatString, err)
		return nil, status.Error(codes.Internal, errString)
	}

	grpcInterests := model.ToInterestProtos(interests)

	return &pb.InterestResponse{
		Interests: grpcInterests,
	}, nil
}
