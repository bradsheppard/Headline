package api

import (
	"context"
	"log"
	"os"
	"testing"

	"headline/model"

	pb "headline/proto/article"

	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
        err := InitDb("host=headline-postgresql user=postgres password=pass123")

        if err != nil {
                log.Fatalf("Failed to connect to DB: %v", err)
                return
        }

        exitVal := m.Run()
        os.Exit(exitVal)
}

func equal(article1 *pb.Article, article2 *pb.Article) bool {
        return article1.Title == article2.Title && article1.Summary == article2.Summary &&
                article1.Link == article2.Link
}

type articleExpectation struct {
        out []*pb.Article
        err error
}

type SetupArticleResult struct {
        client *pb.ArticleServiceClient
        closer func()
}

func setupArticles(ctx context.Context) (*SetupArticleResult, error) {
        db.AutoMigrate(&model.Article{})
        db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Article{})

        conn, err := GenerateTestServer(ctx)

        if err != nil {
                return nil, err
        }
        
        client := pb.NewArticleServiceClient(conn.clientConn)
        pb.RegisterArticleServiceServer(conn.server, &ArticleServer{})

        go conn.startup()

        return &SetupArticleResult{
                client: &client,
                closer: conn.closer,
        }, nil
}

func TestArticle_GetArticles_Empty(t *testing.T) {
        ctx := context.Background()
        setup, err := setupArticles(ctx)

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

        expected := articleExpectation{
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
        setup, err := setupArticles(ctx)

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

        expected := articleExpectation{
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
