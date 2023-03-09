package api

import (
	"context"
	"log"
	"os"
	"testing"

	"headline/model"

	pb "headline/proto"

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

func equal(article1 *pb.Article, article2 *pb.Article) bool {
        return article1.Title == article2.Title && article1.Summary == article2.Summary &&
                article1.Link == article2.Link
}

type expectation struct {
        out []*pb.Article
        err error
}

type SetupResult struct {
        client *pb.ArticleServiceClient
        closer func()
}

func setup(ctx context.Context) (*SetupResult, error) {
        conn, err := GenerateTestServer(ctx)

        if err != nil {
                return nil, err
        }
        
        client := pb.NewArticleServiceClient(conn.clientConn)
        pb.RegisterArticleServiceServer(conn.server, &Server{})

        go conn.startup()

        return &SetupResult{
                client: &client,
                closer: conn.closer,
        }, nil
}

func TestArticle_GetArticles_Empty(t *testing.T) {
        ctx := context.Background()
        setup, err := setup(ctx)

        if err != nil {
                t.Errorf("Setup error: %v", err)
                t.FailNow()
        }

        defer setup.closer()

        client := *setup.client

        req := &pb.GetArticlesRequest{UserId: 999}
        articles, err := client.GetArticles(ctx, req)

        if err != nil {
                t.Errorf("Error: %v", err)
                t.FailNow()
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
        setup, err := setup(ctx)

        defer setup.closer()

        if err != nil {
                t.Errorf("Setup error: %v", err)
                t.FailNow()
        }

        client := *setup.client

        _, err = client.CreateArticle(ctx, &pb.CreateArticleRequest{
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
