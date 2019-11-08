package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "myTopic"
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:19092", topic, 0)
	//conn, err := kafka.Dial("tcp", "localhost:19092")
	defer conn.Close()
	if err != nil {
		fmt.Printf("Connect error: %s\n", err.Error())
		return
	}

	message := kafka.Message{
		//Topic: "myTopic",
		//Key:   []byte("first"),
		Value: []byte("first message"),
		Time:  time.Now(),
	}

	n, err := conn.WriteMessages(message)
	if err != nil {
		fmt.Printf("Write message error: %s\n", err.Error())
		return
	}

	fmt.Printf("We have send %d bytes messages.", n)
}
