package bootstrap

import (
	"log"

	"github.com/IBM/sarama"
)

func setupProducer(cfg *Config) sarama.SyncProducer {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{cfg.KafkaServerAddress},
		config)
	if err != nil {
		log.Fatal("failed to initialize producer: %v", err)
	}
	return producer
}
