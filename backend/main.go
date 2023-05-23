package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"headline/api"

	article_pb "headline/proto/article"
	interest_pb "headline/proto/interest"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	database := os.Getenv("DATABASE")

        kafkaHost := os.Getenv("KAFKA_HOST")
        kafkaTopic := os.Getenv("KAFKA_TOPIC")

	dbPath := fmt.Sprintf("host=%s user=%s password=%s database=%s", host, user, password, database)
	err := api.InitDb(dbPath)

	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
		return
	}

        api.InitMessaging(kafkaHost, kafkaTopic)

	err = api.AutoMigrate()

	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
		return
	}

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()

	article_pb.RegisterArticleServiceServer(s, &api.ArticleServer{})
	interest_pb.RegisterInterestServiceServer(s, &api.InterestServer{})

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
		return
	}
}
