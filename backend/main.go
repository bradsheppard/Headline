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

const dbPath = "test.db"

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedArticleServiceServer
}

func main() {
	api.InitDb(dbPath)

	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterArticleServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}
}
