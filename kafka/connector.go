package kafka

import (
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
	"time"
)

var writer *kafka.Writer

func Configure(kafkaBrokerUrls []string, clientId string, topic string) (w *kafka.Writer, err error) {
	dialer := &kafka.Dialer{
		ClientID:        clientId,
		Timeout:         10 * time.Second,
	}

	config := kafka.WriterConfig{
		Brokers:           kafkaBrokerUrls,
		Topic:             topic,
		Dialer:            dialer,
		Balancer:          &kafka.LeastBytes{},
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		CompressionCodec:  snappy.NewCompressionCodec(),
	}
	w = kafka.NewWriter(config)
	writer = w
	return w, nil
}
