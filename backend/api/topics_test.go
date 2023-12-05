package api

import (
	"context"
	"testing"
	"time"

	"headline/model"

	pb "headline/proto/topic"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type SetupTopicResult struct {
	client *pb.TopicServiceClient
	closer func()
}

type topicExpectation struct {
	out []*pb.Topic
	err error
}

var (
    config = &oauth2.Config{
        ClientID: "769516741930-aqo3qtl4vjf1odrgng2j4ndq18ddllut.apps.googleusercontent.com",
        ClientSecret: "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",
        Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile"},
        Endpoint: google.Endpoint,
    }
)

func getToken(ctx context.Context, t *testing.T) (*string, error) {
    // Step 1: Get device code and user code
    deviceCode, err := config.DeviceAuth(ctx)

    if err != nil {
        t.Errorf("Failed to start device flow: %v", err)
        return nil, err
    }

    t.Logf("Visit %s and enter code %s to grant access\n", deviceCode.VerificationURI, deviceCode.UserCode)

	// Step 2: Poll for token using device code
	for {
		token, err := config.Exchange(ctx, deviceCode.DeviceCode)
		if err == nil {
			t.Logf("Access Token: %s\n", token.AccessToken)
			t.Logf("Refresh Token: %s\n", token.RefreshToken)
			t.Logf("Token Type: %s\n", token.TokenType)
			t.Logf("Expiry: %s\n", token.Expiry)

            return &token.AccessToken, nil
		}

		if e, ok := err.(*oauth2.RetrieveError); ok {
			if e.Response.StatusCode == 400 && e.Response.Header.Get("Retry-After") != "" {
				retryAfter, _ := time.ParseDuration(e.Response.Header.Get("Retry-After"))
				time.Sleep(retryAfter)
				continue
			}
		}

		t.Errorf("Failed to exchange device code for token: %v", err)
	}
}

func setupTopic(ctx context.Context) (*SetupTopicResult, error) {
	db.AutoMigrate(&model.Topic{}, &model.Article{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Topic{}, &model.Article{})

	conn, err := GenerateTestServer(ctx)

	if err != nil {
		return nil, err
	}

	client := pb.NewTopicServiceClient(conn.clientConn)
	pb.RegisterTopicServiceServer(conn.server, &TopicServer{})

	go conn.startup()

	return &SetupTopicResult{
		client: &client,
		closer: conn.closer,
	}, nil
}

func equalTopics(topic1 *pb.Topic, topic2 *pb.Topic) bool {
	return topic1.Name == topic2.Name
}

func checkEqual(t *testing.T, expected *topicExpectation, actual *pb.TopicResponse) {
	if len(expected.out) != len(actual.Topics) {
		t.Errorf("Inequal length for articles")
		t.Errorf("Expected Length: %d", len(expected.out))
		t.Errorf("Actual length: %d", len(actual.Topics))
		t.FailNow()
	}

	for i := range expected.out {
		exp := expected.out[i]
		actual := actual.Topics[i]

		if !equalTopics(exp, actual) {
			t.Errorf("Expected -> %q\nGot: %q", exp, actual)
			t.FailNow()
		}
	}
}

func TestTopic_GetTopics_Empty(t *testing.T) {
	ctx := context.Background()
	setup, err := setupTopic(ctx)

	if err != nil {
		t.Errorf("Setup error: %v", err)
		t.FailNow()
	}

    _, err = getToken(ctx, t)

    if err != nil {
        t.Errorf("Token error: %v", err)
        t.FailNow()
    }

	defer setup.closer()

	client := *setup.client

	req := &pb.GetTopicsRequest{UserId: 999}
	topics, err := client.GetTopics(ctx, req)

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	expected := topicExpectation{
		out: []*pb.Topic{},
		err: nil,
	}

	checkEqual(t, &expected, topics)
}

func TestTopic_GetTopics_NotEmpty(t *testing.T) {
	ctx := context.Background()
	setup, err := setupTopic(ctx)

	defer setup.closer()

	if err != nil {
		t.Errorf("Setup error: %v", err)
		t.FailNow()
	}

	client := *setup.client

	_, err = client.AddTopics(ctx, &pb.AddTopicsRequest{
		Topics: []*pb.Topic{
			&pb.Topic{
				Name: "Topic 1",
			},
		},
		UserId: 10,
	})

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	_, err = client.AddTopics(ctx, &pb.AddTopicsRequest{
		Topics: []*pb.Topic{
			&pb.Topic{
				Name: "Topic 2",
			},
		},
		UserId: 20,
	})

	topics, err := client.GetTopics(ctx, &pb.GetTopicsRequest{UserId: 10})

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	expected := topicExpectation{
		out: []*pb.Topic{
			&pb.Topic{
				Name: "Topic 1",
			},
		},
		err: nil,
	}

	checkEqual(t, &expected, topics)
}

func TestTopic_DeleteTopics(t *testing.T) {
	ctx := context.Background()
	setup, err := setupTopic(ctx)

	defer setup.closer()

	if err != nil {
		t.Errorf("Setup error: %v", err)
		t.FailNow()
	}

	client := *setup.client

	_, err = client.AddTopics(ctx, &pb.AddTopicsRequest{
		Topics: []*pb.Topic{
			&pb.Topic{
				Name: "Topic 1",
			},
			&pb.Topic{
				Name: "Topic 2",
			},
		},
		UserId: 20,
	})

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	_, err = client.RemoveTopics(ctx, &pb.RemoveTopicsRequest{
		TopicNames: []string{"Topic 1"},
		UserId:     20,
	})

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	topics, err := client.GetTopics(ctx, &pb.GetTopicsRequest{UserId: 20})

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	expected := topicExpectation{
		out: []*pb.Topic{
			&pb.Topic{
				Name: "Topic 2",
			},
		},
		err: nil,
	}

	checkEqual(t, &expected, topics)
}

func TestTopic_GetPendingTopics(t *testing.T) {
	ctx := context.Background()
	setup, err := setupTopic(ctx)

	defer setup.closer()

	if err != nil {
		t.Errorf("Setup error: %v", err)
		t.FailNow()
	}

	client := *setup.client

	_, err = client.AddTopics(ctx, &pb.AddTopicsRequest{
		Topics: []*pb.Topic{
			&pb.Topic{
				Name: "Topic 1",
			},
			&pb.Topic{
				Name: "Topic 2",
			},
		},
		UserId: 21,
	})

	currentTime := time.Now()
	streamClient, err := client.GetPendingTopics(ctx, &pb.GetPendingTopicsRequest{LastUpdated: timestamppb.New(currentTime)})

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	expected := topicExpectation{
		out: []*pb.Topic{
			&pb.Topic{
				Name: "Topic 1",
			},
			&pb.Topic{
				Name: "Topic 2",
			},
		},
		err: nil,
	}

	topics, err := streamClient.Recv()

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	checkEqual(t, &expected, topics)
}
