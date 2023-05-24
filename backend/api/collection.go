package api

import (
	"context"
	"errors"
	"headline/collection"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/segmentio/kafka-go"
)

var (
	writer *kafka.Writer
)

const retries = 3

func InitCollectionWriter(broker string, topic string) {
	writer = &kafka.Writer{
		Addr:                   kafka.TCP(broker),
		Topic:                  topic,
		AllowAutoTopicCreation: true,
	}
}

func StartCollection(ctx context.Context, col *collection.Collection) error {
	bytes, err := proto.Marshal(col)

	if err != nil {
		return err
	}

	for i := 0; i < retries; i++ {
		err = writer.WriteMessages(ctx,
			kafka.Message{
				Value: bytes,
			},
		)

		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(time.Millisecond * 250)
			continue
		}

		if err != nil {
			return err
		}

		break
	}

	if err != nil {
		return err
	}

	return nil
}
