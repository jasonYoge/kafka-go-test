package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "myTopic"
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:19092", "localhost:29092", "localhost:39092"},
		GroupID:  "my-group",
		Topic:    topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
