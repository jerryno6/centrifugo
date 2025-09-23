package brokerpublishing

import (
	"context"
	"fmt"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/twmb/franz-go/pkg/kgo"
)

// TODO: should not use static var, move it to initializer when we init the app
var (
	kafkaClient     *kgo.Client
	kafkaClientOnce sync.Once
)

const clientID = "centrifugo"

func InitKafkaClient(brokers []string) error {
	var err error
	kafkaClientOnce.Do(func() {
		opts := []kgo.Opt{
			kgo.SeedBrokers(brokers...),
			kgo.ClientID(clientID),
			kgo.RequiredAcks(kgo.AllISRAcks()),
		}
		kafkaClient, err = kgo.NewClient(opts...)

		log.Info().Msg("Kafka client created with brokers: " + fmt.Sprintf("%v", brokers))
	})

	return err
}

func GetKafkaClient() *kgo.Client {
	return kafkaClient
}

// Publish sends a message to the specified Kafka topic.
// It returns an error if the message cannot be sent synchronously.
func Publish(client *kgo.Client, topic *string, key []byte, data []byte, headers []kgo.RecordHeader) {
	if client == nil {
		log.Error().Msg("Publish(). Kafka client is nil")
	}

	record := &kgo.Record{
		Topic: *topic,
		Key:   key,
		Value: data,
	}

	if len(headers) > 0 {
		record.Headers = headers
	}

	// Use ProduceSync for synchronous error handling
	errs := client.ProduceSync(context.Background(), record)
	if len(errs) > 0 {

		// loop through errs and log each error
		for _, err := range errs {
			// check err and err.Err for nil
			if err.Err != nil {
				log.Error().Err(err.Err).Msg("Publish(). Failed to publish message to Kafka")
			}
		}
	}
}
