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
	return article1.Title == article2.Title && article1.Description == article2.Description &&
		article1.Url == article2.Url && article1.ImageUrl == article2.ImageUrl &&
		article1.Source == article2.Source && article1.Interest == article2.Interest
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

	req := &pb.User{UserId: 999}
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

	_, err = client.SetUserArticles(ctx, &pb.UserArticles{
		Articles: []*pb.Article{
			&pb.Article{
				Title:       "New Title 1",
				Description: "New Description 1",
				Url:         "New Url 1",
				ImageUrl:    "New Image Url 1",
				Source:      "New Source 1",
				Interest:    "New Interest 1",
			},
		},
		UserId: 1,
	})

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	_, err = client.SetUserArticles(ctx, &pb.UserArticles{
		Articles: []*pb.Article{
			&pb.Article{
				Title:       "New Title 2",
				Description: "New Description 2",
				Url:         "New Url 2",
				ImageUrl:    "New Image Url 2",
				Source:      "New Source 2",
				Interest:    "New Interest 2",
			},
		},
		UserId: 2,
	})

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	articles, err := client.GetArticles(ctx, &pb.User{UserId: 1})

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	expected := articleExpectation{
		out: []*pb.Article{
			&pb.Article{
				Title:       "New Title 1",
				Description: "New Description 1",
				Url:         "New Url 1",
				ImageUrl:    "New Image Url 1",
				Source:      "New Source 1",
				Interest:    "New Interest 1",
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
