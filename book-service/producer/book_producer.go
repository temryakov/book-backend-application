package prod

import (
	"book-service/domain"
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/segmentio/kafka-go"
)

type bookProducer struct {
	producer *kafka.Producer
	topic    string
}

func NewBookProducer(producer *kafka.Producer) domain.BookProducer {
	return &bookProducer{
		producer: producer,
		topic:    "books",
	}
}

func (p *bookProducer) DeleteBook(ctx context.Context, id string) {
	p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.topic, Partition: kafka.PartitionAny},
		Key:            []byte(id),
		Value:          []byte("delete"),
	}, nil)
}
