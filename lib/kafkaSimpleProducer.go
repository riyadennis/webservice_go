package lib

import (
	"log"
	"github.com/Shopify/sarama"
)

func SendMessageSynchronously(msg string) {
	brokers := []string{"localhost:9092"}

	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := producer.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	topic := "test"

	message := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(msg),
	}
	producer.SendMessage(message)
}
