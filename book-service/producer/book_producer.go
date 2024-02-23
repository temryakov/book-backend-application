package prod

import (
	"book-service/domain"
	"log"
	"strconv"

	"github.com/IBM/sarama"
)

type bookProducer struct {
	producer sarama.SyncProducer
	topic    string
}

func NewBookProducer(producer sarama.SyncProducer) domain.BookProducer {
	return &bookProducer{
		producer: producer,
		topic:    "books",
	}
}
func (bp *bookProducer) WriteMessage(bookId int, data string) error {
	msg := &sarama.ProducerMessage{
		Topic: bp.topic,
		Key:   sarama.StringEncoder(strconv.Itoa(bookId)),
		Value: sarama.StringEncoder(data),
	}

	p, offset, err := bp.producer.SendMessage(msg)
	log.Println("Partition: %w, offset: %w", p, offset)
	return err
}
func (bp *bookProducer) DeleteBook(bookId int) {
	data := "DELETE"
	bp.WriteMessage(bookId, data)
}
