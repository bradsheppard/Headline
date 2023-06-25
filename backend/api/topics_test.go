package api

import (
	"context"
	"testing"

	"headline/model"

	pb "headline/proto/topic"

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

        checkEqual(t,&expected, topics)
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
				Name:   "Topic 1",
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
				Name:   "Topic 2",
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
				Name:   "Topic 1",
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
				Name:   "Topic 1",
			},
			&pb.Topic{
				Name:   "Topic 2",
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
                UserId: 20,
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
				Name:   "Topic 2",
			},
		},
		err: nil,

	}

        checkEqual(t, &expected, topics)
}
