package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"headline/api"

	pb "headline/proto"

	"google.golang.org/grpc"
)

const dbPath = "host=headline-postgresql user=postgres password=pass123"

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	err := api.InitDb(dbPath)

	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
		return
	}

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterArticleServiceServer(s, &api.Server{})

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
		return
	}
}
