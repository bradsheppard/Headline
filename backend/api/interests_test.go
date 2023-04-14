package api

import (
	"context"
	"testing"

	"headline/model"

	pb "headline/proto/interest"

	"gorm.io/gorm"
)

type SetupInterestResult struct {
	client *pb.InterestServiceClient
	closer func()
}

type interestExpectation struct {
	out []*pb.Interest
	err error
}

func equalInterests(interest1 *pb.Interest, interest2 *pb.Interest) bool {
	return interest1.Name == interest2.Name && interest1.UserId == interest2.UserId
}

func setupInterest(ctx context.Context) (*SetupInterestResult, error) {
	db.AutoMigrate(&model.Interest{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Interest{})

	conn, err := GenerateTestServer(ctx)

	if err != nil {
		return nil, err
	}

	client := pb.NewInterestServiceClient(conn.clientConn)
	pb.RegisterInterestServiceServer(conn.server, &InterestServer{})

	go conn.startup()

	return &SetupInterestResult{
		client: &client,
		closer: conn.closer,
	}, nil
}

func TestInterest_GetInterests_Empty(t *testing.T) {
	ctx := context.Background()
	setup, err := setupInterest(ctx)

	if err != nil {
		t.Errorf("Setup error: %v", err)
		t.FailNow()
	}

	defer setup.closer()

	client := *setup.client

	req := &pb.GetInterestsRequest{UserId: 999}
	interests, err := client.GetInterests(ctx, req)

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	expected := interestExpectation{
		out: []*pb.Interest{},
		err: nil,
	}

	if len(expected.out) != len(interests.Interests) {
		t.Errorf("Inequal length for articles")
		t.Errorf("Expected Length: %d", len(expected.out))
		t.Errorf("Actual length: %d", len(interests.Interests))
		t.FailNow()
	}

	for i := range expected.out {
		exp := expected.out[i]
		actual := interests.Interests[i]

		if !equalInterests(exp, actual) {
			t.Errorf("Expected -> %q\nGot: %q", exp, actual)
			t.FailNow()
		}
	}
}

func TestInterest_GetInterests_NotEmpty(t *testing.T) {
	ctx := context.Background()
	setup, err := setupInterest(ctx)

	defer setup.closer()

	if err != nil {
		t.Errorf("Setup error: %v", err)
		t.FailNow()
	}

	client := *setup.client

	_, err = client.AddInterests(ctx, &pb.AddInterestsRequest{
		Interests: []*pb.CreateInterest{
			&pb.CreateInterest{
				Name:   "Interest 1",
				UserId: 1,
			},
			&pb.CreateInterest{
				Name:   "Interest 2",
				UserId: 2,
			},
		},
	})

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	interests, err := client.GetInterests(ctx, &pb.GetInterestsRequest{UserId: 1})

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	expected := interestExpectation{
		out: []*pb.Interest{
			&pb.Interest{
				Name:   "Interest 1",
				UserId: 1,
			},
		},
		err: nil,
	}

	if len(expected.out) != len(interests.Interests) {
		t.Errorf("Inequal length for articles")
		t.Errorf("Expected Length: %d", len(expected.out))
		t.Errorf("Actual length: %d", len(interests.Interests))
		t.FailNow()
	}

	for i := range expected.out {
		exp := expected.out[i]
		actual := interests.Interests[i]

		if !equalInterests(exp, actual) {
			t.Errorf("Expected -> %q\nGot: %q", exp, actual)
			t.FailNow()
		}
	}
}

func TestInterest_DeleteInterests_NotEmpty(t *testing.T) {
	ctx := context.Background()
	setup, err := setupInterest(ctx)

	defer setup.closer()

	if err != nil {
		t.Errorf("Setup error: %v", err)
		t.FailNow()
	}

	client := *setup.client

	res, err := client.AddInterests(ctx, &pb.AddInterestsRequest{
		Interests: []*pb.CreateInterest{
			&pb.CreateInterest{
				Name:   "CreateInterest 1",
				UserId: 1,
			},
			&pb.CreateInterest{
				Name:   "CreateInterest 2",
				UserId: 2,
			},
		},
	})

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	client.DeleteInterests(ctx, &pb.DeleteInterestsRequest{Ids: []uint64{res.Interests[0].Id}})
	interests, err := client.GetInterests(ctx, &pb.GetInterestsRequest{UserId: res.Interests[1].UserId})

	if err != nil {
		t.Errorf("Error: %v", err)
		t.FailNow()
	}

	expected := interestExpectation{
		out: []*pb.Interest{res.Interests[1]},
		err: nil,
	}

	if len(expected.out) != len(interests.Interests) {
		t.Errorf("Inequal length for articles")
		t.Errorf("Expected Length: %d", len(expected.out))
		t.Errorf("Actual length: %d", len(interests.Interests))
		t.FailNow()
	}

	for i := range expected.out {
		exp := expected.out[i]
		actual := interests.Interests[i]

		if !equalInterests(exp, actual) {
			t.Errorf("Expected -> %q\nGot: %q", exp, actual)
			t.FailNow()
		}
	}
}
