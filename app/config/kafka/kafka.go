package kafkaMs

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func GetKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func SendMessage(kafkaWriter *kafka.Writer, key, message string) {
	msg := kafka.Message{
		Key:   []byte(key),
		Value: []byte(message),
	}
	kafkaWriter.WriteMessages(context.Background(), msg)
}
