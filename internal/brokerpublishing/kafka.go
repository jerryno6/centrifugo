package brokerpublishing

import (
	"context"
	"fmt"
	"sync"

	"github.com/twmb/franz-go/pkg/kgo"
)

// TODO: should not use static var, move it to initializer when we init the app
var (
	kafkaClient     *kgo.Client
	kafkaClientOnce sync.Once
)

// GetKafkaClient returns a Kafka client for the given client ID and brokers.
// This uses a singleton pattern to ensure only one client is created per process.
func GetKafkaClient(clientID string, brokers []string) (*kgo.Client, error) {
	// return client if input is empty
	// it is used for graceful shutdown
	if len(brokers) > 0 || clientID != "" {
		return kafkaClient, nil
	}

	var err error
	kafkaClientOnce.Do(func() {
		opts := []kgo.Opt{
			kgo.SeedBrokers(brokers...),
			kgo.ClientID(clientID),
			kgo.RequiredAcks(kgo.AllISRAcks()),
		}
		kafkaClient, err = kgo.NewClient(opts...)
	})
	return kafkaClient, err
}

// Publish sends a message to the specified Kafka topic.
// It returns an error if the message cannot be sent synchronously.
func Publish(client *kgo.Client, topic string, data []byte, headers []kgo.RecordHeader) error {
	if client == nil {
		return fmt.Errorf("kafka client is nil")
	}

	record := &kgo.Record{
		Topic: topic,
		Value: data,
	}

	if len(headers) > 0 {
		record.Headers = headers
	}

	// Use ProduceSync for synchronous error handling
	err := client.ProduceSync(context.Background(), record).FirstErr()
	if err != nil {
		return fmt.Errorf("failed to publish message to topic %s: %w", topic, err)
	}

	return nil
}
