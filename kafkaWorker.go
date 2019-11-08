package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

var (
	kafkaBrokerUrl     string
	kafkaVerbose       bool
	kafkaTopic         string
	kafkaConsumerGroup string
	kafkaClientId      string
)

func init() {
	flag.StringVar(
		&kafkaBrokerUrl,
		"kafka-brokers",
		"localhost:19092,localhost:29092,localhost:39092",
		"Kafka brokers in comma separated value")
	flag.BoolVar(&kafkaVerbose, "kafka-verbose", true, "Kafka verbose logging")
	flag.StringVar(&kafkaTopic, "kafka-topic", "foo", "Kafka topic. Only one topic per worker.")
	flag.StringVar(&kafkaConsumerGroup, "kafka-consumer-group", "consumer-group", "Kafka consumer group")
	flag.StringVar(&kafkaClientId, "kafka-client-id", "my-client-id", "Kafka client id")

	flag.Parse()
}

func main() {
	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	brokers := strings.Split(kafkaBrokerUrl, ",")

	config := kafka.ReaderConfig{
		Brokers:         brokers,
		GroupID:         kafkaClientId,
		Topic:           kafkaTopic,
		MinBytes:        10e3,
		MaxBytes:        10e6,
		MaxWait:         1 * time.Second,
		ReadLagInterval: -1,
	}

	reader := kafka.NewReader(config)
	defer reader.Close()

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("some error occurs: %s", err.Error())
			continue
		}

		value := m.Value
		if m.CompressionCodec == snappy.NewCompressionCodec() {
			_, err = snappy.NewCompressionCodec().Decoder(value, m.Value)
		}

		if err != nil {
			fmt.Printf("some error occurs: %s", err.Error())
			continue
		}

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s\n", m.Topic, m.Partition, m.Offset, string(value))
	}
}
