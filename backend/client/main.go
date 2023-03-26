package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "headline/proto/article"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	grpcAddr := net.JoinHostPort("localhost", "8088")
	conn, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Error connecting to grpc server: %v", err)
                return
	}

	client := pb.NewArticleServiceClient(conn)

	runGetArticles(client)
}


func runGetArticles(client pb.ArticleServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

        req := &pb.GetArticlesRequest{UserId: 1}

	res, err := client.GetArticles(ctx, req)

	if err != nil {
		log.Printf("Error getting contact: %v", err)
		return
	}

	fmt.Printf("Got articles: %+v\n", res)
}

