package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:19092", "localhost:29092", "localhost:39092"},
		Topic:    "myTopic",
		Balancer: &kafka.LeastBytes{},
	})

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)

		err := w.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte("Message"),
			Value: []byte(fmt.Sprintf("This is number %d message.", i)),
		})

		if err != nil {
			fmt.Printf("error occurs: %s\n", err.Error())
			return
		}

		fmt.Println("success~")
	}
}
