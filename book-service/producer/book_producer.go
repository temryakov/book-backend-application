package prod

import (
	"book-service/domain"

	"github.com/confluentinc/confluent-kafka-go/kafka"
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

func (bp *bookProducer) WriteMessage(id string, data string) {

	topicPartition := kafka.TopicPartition{
		Topic:     &bp.topic,
		Partition: kafka.PartitionAny,
	}

	bp.producer.Produce(
		&kafka.Message{
			TopicPartition: topicPartition,
			Key:            []byte(id),
			Value:          []byte(data),
		},
		nil,
	)
}

func (bp *bookProducer) DeleteBook(bookId string, data string) {
	bp.WriteMessage(bookId, data)
}
