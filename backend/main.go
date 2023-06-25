package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"headline/api"

	article_pb "headline/proto/article"
	topic_pb "headline/proto/topic"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	s, lis, err := initGrpc()

	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	initMessaging()

	log.Printf("Server listening at %v", (*lis).Addr())

	if err := s.Serve(*lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
		return
	}
}

func initMessaging() {
	kafkaHost := os.Getenv("KAFKA_HOST")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")

	api.InitCollectionWriter(kafkaHost, kafkaTopic)
}

func initGrpc() (*grpc.Server, *net.Listener, error) {
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	database := os.Getenv("DATABASE")

	dbPath := fmt.Sprintf("host=%s user=%s password=%s database=%s", host, user, password, database)

	err := api.InitDb(dbPath)

	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
		return nil, nil, err
	}

	err = api.AutoMigrate()

	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
		return nil, nil, err
	}

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return nil, nil, err
	}

	s := grpc.NewServer()

	article_pb.RegisterArticleServiceServer(s, &api.ArticleServer{})
	topic_pb.RegisterTopicServiceServer(s, &api.TopicServer{})

	return s, &lis, nil
}
