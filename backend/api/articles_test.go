package api

import (
	"context"
	"log"
	"os"
	"reflect"
	"testing"

	"headline/model"

	article_pb "headline/proto/article"
	topic_pb "headline/proto/topic"

	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	err := InitDb("host=headline-postgresql user=postgres password=pass123")

	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
		return
	}

	InitCollectionWriter("queue-kafka-bootstrap:9092", "start-collection")

	exitVal := m.Run()
	os.Exit(exitVal)
}

type topicArticlesExpectation struct {
	out *article_pb.TopicArticles
	err error
}

type SetupArticleResult struct {
	articleClient *article_pb.ArticleServiceClient
	topicClient   *topic_pb.TopicServiceClient
	closer        func()
}

func setupArticles(ctx context.Context) (*SetupArticleResult, error) {
	err := db.AutoMigrate(&model.Topic{}, &model.Article{})

	if err != nil {
		return nil, err
	}

	err = db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Article{}, &model.Topic{}).Error

	if err != nil {
		return nil, err
	}

	conn, err := GenerateTestServer(ctx)

	if err != nil {
		return nil, err
	}

	article_client := article_pb.NewArticleServiceClient(conn.clientConn)
	article_pb.RegisterArticleServiceServer(conn.server, &ArticleServer{})

	topic_client := topic_pb.NewTopicServiceClient(conn.clientConn)
	topic_pb.RegisterTopicServiceServer(conn.server, &TopicServer{})

	go conn.startup()

	return &SetupArticleResult{
		articleClient: &article_client,
		topicClient:   &topic_client,
		closer:        conn.closer,
	}, nil
}

func TestArticle_GetTopicArticles_Empty(t *testing.T) {
    ctx := getContext()
	setup, err := setupArticles(ctx)

	if err != nil {
		t.Errorf("Setup error: %v", err)
		t.FailNow()
	}

	defer setup.closer()

	client := *setup.articleClient

	req := &article_pb.TopicNames{
		Topics: []string{"Topic 1", "Topic 2"},
	}

	topicArticles, err := client.GetTopicArticles(ctx, req)

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	expected := topicArticlesExpectation{
		out: &article_pb.TopicArticles{},
		err: nil,
	}

	if len(expected.out.TopicArticles) != len(topicArticles.TopicArticles) {
		t.Errorf("Inequal length for articles")
		t.Errorf("Expected Length: %d", len(expected.out.TopicArticles))
		t.Errorf("Actual length: %d", len(topicArticles.TopicArticles))
		t.FailNow()
	}
}

func TestArticle_SetTrendingArticles_GetTrendingArticles_NotEmpty(t *testing.T) {
    ctx := getContext()
	setup, err := setupArticles(ctx)

	defer setup.closer()

	if err != nil {
		t.Errorf("Setup error: %v", err)
		t.FailNow()
	}

	articleClient := *setup.articleClient
	topicClient := *setup.topicClient

	_, err = articleClient.SetTrendingArticles(ctx, &article_pb.Articles{
		Articles: []*article_pb.Article{
			&article_pb.Article{
				Title:       "Trending Title 1",
				Description: "Trending Description 1",
				Url:         "Trending Url 1",
				ImageUrl:    "Trending Image Url 1",
				Source:      "Trending Source 1",
			},
			&article_pb.Article{
				Title:       "Trending Title 2",
				Description: "Trending Description 2",
				Url:         "Trending Url 2",
				ImageUrl:    "Trending Image Url 2",
				Source:      "Trending Source 2",
			},
		},
	})

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	_, err = topicClient.AddTopics(ctx, &topic_pb.AddTopicsRequest{
		Topics: []*topic_pb.Topic{
			&topic_pb.Topic{
				Name: "Unrelated Topic",
			},
		},
	})

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	_, err = articleClient.SetTopicArticles(ctx, &article_pb.TopicArticles{
		TopicArticles: map[string]*article_pb.Articles{
			"Unrelated Topic": &article_pb.Articles{
				Articles: []*article_pb.Article{
					&article_pb.Article{
						Title:       "New Title 1",
						Description: "New Description 1",
						Url:         "New Url 1",
						ImageUrl:    "New Image Url 1",
						Source:      "New Source 1",
					},
					&article_pb.Article{
						Title:       "New Title 2",
						Description: "New Description 2",
						Url:         "New Url 2",
						ImageUrl:    "New Image Url 2",
						Source:      "New Source 2",
					},
				},
			},
		},
	})

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	articles, err := articleClient.GetTrendingArticles(ctx, &emptypb.Empty{})

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	expected := []*article_pb.Article{
		&article_pb.Article{
			Title:       "Trending Title 1",
			Description: "Trending Description 1",
			Url:         "Trending Url 1",
			ImageUrl:    "Trending Image Url 1",
			Source:      "Trending Source 1",
		},
		&article_pb.Article{
			Title:       "Trending Title 2",
			Description: "Trending Description 2",
			Url:         "Trending Url 2",
			ImageUrl:    "Trending Image Url 2",
			Source:      "Trending Source 2",
		},
	}

	if len(expected) != len(articles.Articles) {
		t.Errorf("Inequal length for articles")
		t.Errorf("Expected Length: %d", len(expected))
		t.Errorf("Actual length: %d", len(articles.Articles))
		t.FailNow()
	}

	for i := range expected {
		exp := expected[i]
		actual := articles.Articles[i]

		if !reflect.DeepEqual(exp, actual) {
			t.Errorf("Expected -> %q\nGot: %q", exp, actual)
			t.FailNow()
		}
	}
}

func TestArticle_SetTopicArticles_GetTopicArticles_NotEmpty(t *testing.T) {
    ctx := getContext()
	setup, err := setupArticles(ctx)

	defer setup.closer()

	if err != nil {
		t.Errorf("Setup error: %v", err)
		t.FailNow()
	}

	articleClient := *setup.articleClient
	topicClient := *setup.topicClient

	_, err = topicClient.AddTopics(ctx, &topic_pb.AddTopicsRequest{
		Topics: []*topic_pb.Topic{
			&topic_pb.Topic{
				Name: "Topic 3",
			},
			&topic_pb.Topic{
				Name: "Topic 4",
			},
		},
	})

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	_, err = articleClient.SetTopicArticles(ctx, &article_pb.TopicArticles{
		TopicArticles: map[string]*article_pb.Articles{
			"Topic 3": &article_pb.Articles{
				Articles: []*article_pb.Article{
					&article_pb.Article{
						Title:       "New Title 1",
						Description: "New Description 1",
						Url:         "New Url 1",
						ImageUrl:    "New Image Url 1",
						Source:      "New Source 1",
					},
					&article_pb.Article{
						Title:       "New Title 2",
						Description: "New Description 2",
						Url:         "New Url 2",
						ImageUrl:    "New Image Url 2",
						Source:      "New Source 2",
					},
				},
			},
			"Topic 4": &article_pb.Articles{
				Articles: []*article_pb.Article{
					&article_pb.Article{
						Title:       "New Title 3",
						Description: "New Description 3",
						Url:         "New Url 3",
						ImageUrl:    "New Image Url 3",
						Source:      "New Source 3",
					},
					&article_pb.Article{
						Title:       "New Title 4",
						Description: "New Description 4",
						Url:         "New Url 4",
						ImageUrl:    "New Image Url 4",
						Source:      "New Source 4",
					},
				},
			},
		},
	})

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	articles, err := articleClient.GetTopicArticles(ctx, &article_pb.TopicNames{Topics: []string{"Topic 3"}})

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	expected := topicArticlesExpectation{
		out: &article_pb.TopicArticles{
			TopicArticles: map[string]*article_pb.Articles{
				"Topic 3": &article_pb.Articles{
					Articles: []*article_pb.Article{
						&article_pb.Article{
							Title:       "New Title 1",
							Description: "New Description 1",
							Url:         "New Url 1",
							ImageUrl:    "New Image Url 1",
							Source:      "New Source 1",
						},
						&article_pb.Article{
							Title:       "New Title 2",
							Description: "New Description 2",
							Url:         "New Url 2",
							ImageUrl:    "New Image Url 2",
							Source:      "New Source 2",
						},
					},
				},
			},
		},
		err: nil,
	}

	if len(expected.out.TopicArticles) != len(articles.TopicArticles) {
		t.Errorf("Inequal length for articles")
		t.Errorf("Expected Length: %d", len(expected.out.TopicArticles))
		t.Errorf("Actual length: %d", len(articles.TopicArticles))
		t.FailNow()
	}

	for i := range expected.out.TopicArticles {
		exp := expected.out.TopicArticles[i]
		actual := articles.TopicArticles[i]

		if !reflect.DeepEqual(exp, actual) {
			t.Errorf("Expected -> %q\nGot: %q", exp, actual)
			t.FailNow()
		}
	}
}
