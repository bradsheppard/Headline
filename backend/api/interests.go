package api

import (
	"context"
	"headline/model"
	"log"

	pb "headline/proto/interest"

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

	if err := db.Create(&interests).Error; err != nil {
		log.Printf(errFormatString, err)
		return nil, status.Error(codes.Internal, errString)
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
