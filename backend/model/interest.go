package model

import (
        "gorm.io/gorm"
	pb "headline/proto/interest"
)

type Interest struct {
        gorm.Model
        Name    string
        UserID  int
}

func ToInterestProtos(interests []*Interest) []*pb.Interest {
        var protoInterests []*pb.Interest

        for _, interest := range(interests) {
                protoInterest := &pb.Interest{
                        Id: uint64(interest.ID),
                        Name: interest.Name,
                        UserId: uint64(interest.UserID),
                }
                protoInterests = append(protoInterests, protoInterest)
        }

        return protoInterests
}

func FromInterestProtos(protoInterests []*pb.CreateInterest) []*Interest {
        var interests []*Interest

        for _, protoInterest := range(protoInterests) {
                interest := &Interest{
                        Name: protoInterest.Name,
                        UserID: int(protoInterest.UserId),
                }
                interests = append(interests, interest)
        }

        return interests
}

