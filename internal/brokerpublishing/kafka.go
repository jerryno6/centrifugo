package brokerpublishing

import (
	"context"
	"fmt"

	"sync"

	"github.com/twmb/franz-go/pkg/kgo"
)

var (
	kafkaClient     *kgo.Client
	kafkaClientOnce sync.Once
)

func getKafkaClient(clientId string) (*kgo.Client, error) {
	var err error
	kafkaClientOnce.Do(func() {
		seeds := []string{"localhost:19092", "localhost:29092", "localhost:39092"}
		opts := []kgo.Opt{
			kgo.SeedBrokers(seeds...),
			kgo.ClientID(clientId),
			kgo.RequiredAcks(kgo.AllISRAcks()),
		}
		kafkaClient, err = kgo.NewClient(opts...)
	})
	return kafkaClient, err
}

func Publish(clientId string, userId string, topic string, data []byte, headers []kgo.RecordHeader) error {
	publisherClientId := userId // it's the user who call OnMessage() to publish to kafka
	client, err := getKafkaClient(publisherClientId)
	if err != nil {
		panic(err)
	}

	record := &kgo.Record{
		Topic: topic,
		Value: data,
	}
	if len(headers) > 0 {
		record.Headers = headers
	}

	client.Produce(context.Background(), record, func(record *kgo.Record, err error) {
		if err != nil {
			fmt.Printf("Error sending message: %v \n", err)
		} else {
			fmt.Printf("Message sent: topic: %s, offset: %d, value: %s \n",
				topic, record.Offset, record.Value)
		}
	})

	return nil
}
