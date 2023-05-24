package api

import (
	"context"
	"headline/model"
	"log"

	"headline/collection"

	pb "headline/proto/interest"

	"gorm.io/gorm"

	"github.com/golang/protobuf/ptypes/empty"
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

	if err != nil {
		log.Printf(errDatabaseFormatString, err)
		return nil, status.Error(codes.Internal, errDatabaseString)
	}

	interestsToCollect := []string{}

	for _, interest := range interests {
		interestsToCollect = append(interestsToCollect, interest.Name)
	}

	collection := collection.Collection{
		UserId:    in.UserId,
		Interests: interestsToCollect,
	}

	err = StartCollection(ctx, &collection)

	if err != nil {
		log.Printf(errMessagingFormatString, err)
		return nil, status.Error(codes.Internal, errMessagingString)
	}

	return &pb.InterestResponse{
		Interests: model.ToInterestProtos(interests),
	}, nil
}

func (interestServer *InterestServer) DeleteInterests(ctx context.Context, in *pb.DeleteInterestsRequest) (*empty.Empty, error) {
	if err := db.Delete(&model.Interest{}, in.Ids).Error; err != nil {
		log.Printf(errDatabaseFormatString, err)
		return nil, status.Error(codes.Internal, errDatabaseString)
	}

	return &emptypb.Empty{}, nil
}

func (interestServer *InterestServer) GetInterests(ctx context.Context, in *pb.GetInterestsRequest) (*pb.InterestResponse, error) {
	var interests []*model.Interest
	if err := db.Where(&model.Interest{UserID: int(in.UserId)}).Find(&interests).Error; err != nil {
		log.Printf(errDatabaseFormatString, err)
		return nil, status.Error(codes.Internal, errDatabaseString)
	}

	grpcInterests := model.ToInterestProtos(interests)

	return &pb.InterestResponse{
		Interests: grpcInterests,
	}, nil
}
