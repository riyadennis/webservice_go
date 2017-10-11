package lib

import (
	"github.com/Shopify/sarama"
	"log"
)

type Message struct {
	Jsonmsg string
}

func (msg *Message) Save() {
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
	topic := "products"
	message := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: 0,
		Value:     sarama.StringEncoder(msg.Jsonmsg),
	}
	producer.SendMessage(message)
}
