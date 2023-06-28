package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	article_pb "headline/proto/article"
	topic_pb "headline/proto/topic"

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
	interest_client := topic_pb.NewTopicServiceClient(conn)

	if arg == "GetArticles" {
		topic := os.Args[2]
		runGetArticles(article_client, topic)
	} else if arg == "GetTopics" {
		runGetTopics(interest_client)
	} else if arg == "CreateTopic" {
		interest := os.Args[2]
		runCreateTopic(interest_client, interest)
	} else {
		log.Println("Invalid command")
	}
}

func runGetArticles(client article_pb.ArticleServiceClient, topic string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &article_pb.GetTopicArticlesRequest{
		Topics: []string{
			topic,
		},
	}

	res, err := client.GetTopicArticles(ctx, req)

	if err != nil {
		log.Printf("Error getting contact: %v", err)
		return
	}

	fmt.Printf("Got articles: %+v\n", res)
}

func runGetTopics(client topic_pb.TopicServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &topic_pb.GetTopicsRequest{UserId: 1}

	res, err := client.GetTopics(ctx, req)

	if err != nil {
		log.Printf("Error getting interests: %v", err)
	}

	fmt.Printf("Got interests: %v\n", res)
}

func runCreateTopic(client topic_pb.TopicServiceClient, interest string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &topic_pb.AddTopicsRequest{
		Topics: []*topic_pb.Topic{
			&topic_pb.Topic{
				Name: interest,
			},
		},
		UserId: 1,
	}

	_, err := client.AddTopics(ctx, req)

	if err != nil {
		log.Printf("Error creating interests: %v", err)
	}
}
