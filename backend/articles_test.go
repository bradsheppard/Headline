package main

import (
	"context"
	"log"
	"net"
	"testing"

	pb "headline/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func testserver(ctx context.Context) (pb.ArticleServiceClient, func()) {
        buffer := 101024 * 1024
        lis := bufconn.Listen(buffer)
        
        s := grpc.NewServer()
        pb.RegisterArticleServiceServer(s, &server{})

        go func() {
                if err := s.Serve(lis); err != nil {
                        log.Printf("Error serving server: %v", err)
                }
        }()

        conn, err := grpc.DialContext(ctx, "", 
                grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
                        return lis.Dial()
                }), grpc.WithTransportCredentials(insecure.NewCredentials()))

        if err != nil {
                log.Printf("Error connecting to server: %v", err)
        }

        closer := func() {
                err := lis.Close()

                if err != nil {
                        log.Printf("Error closing listener: %v", err)
                }

                s.Stop()
        }

        client := pb.NewArticleServiceClient(conn)

        return client, closer
}

type expectation struct {
        out []*pb.Article
        err error
}

func TestArticle_GetArticles(t *testing.T) {
        ctx := context.Background()

        client, closer := testserver(ctx)
        defer closer()

        in := &pb.AllArticlesRequest{}
        articles, err := client.GetArticles(ctx, in)

        if err != nil {
                t.Errorf("Error: %v", err)
        }

        expected := expectation{
                out: []*pb.Article{
                        &pb.Article{
                                Title: "Test Article 1",
                                Summary: "Test Summary 1",
                        },
                        &pb.Article{
                                Title: "Test Article 2",
                                Summary: "Test Summary 2",
                        },
                },
                err: nil,
        }

        if len(expected.out) != len(articles.Articles) {
                t.Errorf("Inequal length for articles")
        }

        for i := range expected.out {
                exp := expected.out[i]
                actual := articles.Articles[i]

                if exp.Title != actual.Title || exp.Summary != actual.Summary {
                        t.Errorf("Expected -> %q\nGot: %q", exp, actual)
                }
        }
}
