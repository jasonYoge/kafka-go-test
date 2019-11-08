package main

import (
	"fmt"
	"io"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, err := kafka.Dial("tcp", "localhost:19092")
	defer conn.Close()
	if err != nil {
		fmt.Printf("Connect error: %s\n", err.Error())
		return
	}
	//fmt.Printf("%#v", conn)

	topicConfig := kafka.TopicConfig{
		Topic:             "myTopic",
		NumPartitions:     4,
		ReplicationFactor: 3,
	}

	err = conn.CreateTopics(topicConfig)
	if kafkaErr, ok := err.(kafka.Error); ok {
		fmt.Printf("error description   %s\n", kafkaErr.Description())
		fmt.Printf("error title         %s\n", kafkaErr.Title())
		fmt.Printf("error is temporary: %#v\n", kafkaErr.Temporary())
		fmt.Printf("error is timeout:   %#v\n", kafkaErr.Timeout())
	}

	if err == io.EOF {
		fmt.Println("eof error")
		return
	}

	fmt.Println("success")
}
