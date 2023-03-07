package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "headline/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	grpcAddr := net.JoinHostPort("localhost", "50051")
	conn, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Error connecting to grpc server: %v", err)
	}

	client := pb.NewArticleServiceClient(conn)

	runGetArticles(client)
}


func runGetArticles(client pb.ArticleServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.GetArticles(ctx, &emptypb.Empty{})

	if err != nil {
		log.Printf("Error getting contact: %v", err)
		return
	}

	fmt.Printf("Got articles: %+v\n", res)
}

