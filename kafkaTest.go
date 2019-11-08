package main

import (
	"context"
	"fmt"

	connector "kafka-go-test/kafka"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "my-topic"
	clientId := "my-client-id"
	brokerUrls := []string{
		"localhost:19092",
		"localhost:29092",
		"localhost:39092",
	}

	writer, _ := connector.Configure(brokerUrls, clientId, topic)
	err := writer.WriteMessages(context.Background(), kafka.Message{Value: []byte("One")})
	if err != nil {
		fmt.Printf("write message error: %s", err.Error())
	}

	fmt.Println("write success!")
}
