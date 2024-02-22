package prod

import (
	"book-service/domain"
	"log"
	"strconv"

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
func (bp *bookProducer) WriteMessage(bookId int, data string) {
	id := strconv.Itoa(bookId)
	topicPartition := kafka.TopicPartition{
		Topic:     &bp.topic,
		Partition: kafka.PartitionAny,
	}
	log.Println("Processing to create event with bookId:", bookId, data)
	bp.producer.Produce(
		&kafka.Message{
			TopicPartition: topicPartition,
			Key:            []byte(id),
			Value:          []byte(data),
		},
		nil,
	)
}
func (bp *bookProducer) DeleteBook(bookId int) {
	data := "DELETE"
	bp.WriteMessage(bookId, data)
}
