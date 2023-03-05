package main

import (
        "context"
        "log"

	pb "headline/proto"
)

func (s *server) GetArticles(ctx context.Context, in *pb.AllArticlesRequest) (*pb.ArticleResponse, error) {
	log.Printf("Received a request")

	return &pb.ArticleResponse{
                Articles: []*pb.Article{
                        &pb.Article{
                                Title: "Test Article 1",
                                Summary: "Test Summary 1",
                        },
                        &pb.Article{
                                Title: "Test Article 2",
                                Summary: "Test Summary 2",
                        },
                },
        }, nil
}

