package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	article_pb "headline/proto/article"
	interest_pb "headline/proto/interest"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
        url := os.Getenv("BACKEND_URL")
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Error connecting to grpc server: %v", err)
                return
	}

        arg := os.Args[1]

	article_client := article_pb.NewArticleServiceClient(conn)
        interest_client := interest_pb.NewInterestServiceClient(conn)

        if arg == "GetArticles" {
	        runGetArticles(article_client)
        } else if arg == "GetInterests" {
                runGetInterests(interest_client)
        } else if arg == "CreateInterest"{
                interest := os.Args[2]
                runCreateInterest(interest_client, interest)
        } else {
                log.Println("Invalid command")
        }
}


func runGetArticles(client article_pb.ArticleServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

        req := &article_pb.User{UserId: 1}

	res, err := client.GetArticles(ctx, req)

	if err != nil {
		log.Printf("Error getting contact: %v", err)
		return
	}

	fmt.Printf("Got articles: %+v\n", res)
}

func runGetInterests(client interest_pb.InterestServiceClient) {
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()

        req := &interest_pb.GetInterestsRequest{UserId: 1}

        res, err := client.GetInterests(ctx, req)

        if err != nil {
                log.Printf("Error getting interests: %v", err)
        }

        fmt.Printf("Got interests: %v\n", res)
}

func runCreateInterest(client interest_pb.InterestServiceClient, interest string) {
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()

        req := &interest_pb.AddInterestsRequest{
                UserId: 1,
                Interests: []*interest_pb.CreateInterest{
                        &interest_pb.CreateInterest{
                                Name: interest,
                                UserId: 1,
                        },
                },
        }

        _, err := client.AddInterests(ctx, req)

        if err != nil {
                log.Printf("Error creating interests: %v", err)
        }
}
