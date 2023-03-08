package api

import (
	"context"
	"log"
	"net"
	"os"
	"testing"

	"headline/model"

	pb "headline/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
        err := InitDb("host=headline-postgresql user=postgres password=pass123")

        if err != nil {
                log.Fatalf("Failed to connect to DB: %v", err)
                return
        }

        db.AutoMigrate(&model.Article{})
        db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Article{})

        exitVal := m.Run()
        os.Exit(exitVal)
}

func testserver(ctx context.Context) (pb.ArticleServiceClient, func()) {
        buffer := 101024 * 1024
        lis := bufconn.Listen(buffer)
        
        s := grpc.NewServer()
        pb.RegisterArticleServiceServer(s, &Server{})

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

func equal(article1 *pb.Article, article2 *pb.Article) bool {
        return article1.Title == article2.Title && article1.Summary == article2.Summary &&
                article1.Link == article2.Link
}

type expectation struct {
        out []*pb.Article
        err error
}

func TestArticle_GetArticles_Empty(t *testing.T) {
        ctx := context.Background()
        client, closer := testserver(ctx)
        defer closer()

        req := &pb.GetArticlesRequest{UserId: 999}
        articles, err := client.GetArticles(ctx, req)

        if err != nil {
                t.Errorf("Error: %v", err)
        }

        expected := expectation{
                out: []*pb.Article{},
                err: nil,
        }

        if len(expected.out) != len(articles.Articles) {
                t.Errorf("Inequal length for articles")
                t.Errorf("Expected Length: %d", len(expected.out))
                t.Errorf("Actual length: %d", len(articles.Articles))
                t.FailNow()
        }

        for i := range expected.out {
                exp := expected.out[i]
                actual := articles.Articles[i]

                if !equal(exp, actual) {
                        t.Errorf("Expected -> %q\nGot: %q", exp, actual)
                        t.FailNow()
                }
        }
}

func TestArticle_GetArticles_NotEmpty(t *testing.T) {
        ctx := context.Background()
        client, closer := testserver(ctx)
        defer closer()

        _, err := client.CreateArticle(ctx, &pb.CreateArticleRequest{
                Articles: []*pb.Article{
                        &pb.Article{
                                Title: "New Title 1",
                                Summary: "New Summary 1",
                                Link: "New Link 1",
                                UserId: 1,
                        },
                        &pb.Article{
                                Title: "New Title 2",
                                Summary: "New Summary 2",
                                Link: "New Link 2",
                                UserId: 2,
                        },
                },
        })

        if err != nil {
                t.Errorf("Error: %v", err)
                t.FailNow()
        }

        articles, err := client.GetArticles(ctx, &pb.GetArticlesRequest{UserId: 1})

        if err != nil {
                t.Errorf("Error: %v", err)
                t.FailNow()
        }

        expected := expectation{
                out: []*pb.Article{
                        &pb.Article{
                                Title: "New Title 1",
                                Summary: "New Summary 1",
                                Link: "New Link 1",
                                UserId: 1,
                        },
                },
                err: nil,
        }

        if len(expected.out) != len(articles.Articles) {
                t.Errorf("Inequal length for articles")
                t.Errorf("Expected Length: %d", len(expected.out))
                t.Errorf("Actual length: %d", len(articles.Articles))
                t.FailNow()
        }

        for i := range expected.out {
                exp := expected.out[i]
                actual := articles.Articles[i]

                if !equal(exp, actual) {
                        t.Errorf("Expected -> %q\nGot: %q", exp, actual)
                        t.FailNow()
                }
        }
}
