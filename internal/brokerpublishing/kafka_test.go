//abcgo:build integration

package brokerpublishing

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/centrifugal/centrifugo/v6/internal/api"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
)

const (
	testKafkaBrokerURL = "localhost:29092"
)

// MockDispatcher implements the Dispatcher interface for testing.
type MockDispatcher struct {
	onDispatchCommand     func(ctx context.Context, method string, data []byte) error
	onDispatchPublication func(ctx context.Context, channels []string, pub api.ConsumedPublication) error
}

func (m *MockDispatcher) DispatchCommand(ctx context.Context, method string, data []byte) error {
	return m.onDispatchCommand(ctx, method, data)
}

func (m *MockDispatcher) DispatchPublication(ctx context.Context, channels []string, pub api.ConsumedPublication) error {
	return m.onDispatchPublication(ctx, channels, pub)
}

func produceTestMessage(topic string, message []byte, headers []kgo.RecordHeader) error {
	// Create a new client
	client, err := kgo.NewClient(kgo.SeedBrokers(testKafkaBrokerURL))
	if err != nil {
		return fmt.Errorf("failed to create Kafka client: %w", err)
	}
	defer client.Close()

	// Produce a message
	err = client.ProduceSync(context.Background(), &kgo.Record{
		Topic:     topic,
		Partition: 0,
		Value:     message,
		Headers:   headers,
	}).FirstErr()
	if err != nil {
		return fmt.Errorf("failed to produce message: %w", err)
	}
	return nil
}

func produceManyRecords(records ...*kgo.Record) error {
	client, err := kgo.NewClient(kgo.SeedBrokers(testKafkaBrokerURL))
	if err != nil {
		return fmt.Errorf("failed to create Kafka client: %w", err)
	}
	defer client.Close()
	err = client.ProduceSync(context.Background(), records...).FirstErr()
	if err != nil {
		return fmt.Errorf("failed to produce message: %w", err)
	}
	return nil
}

func produceTestMessageToPartition(topic string, message []byte, partition int32) error {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(testKafkaBrokerURL),
		kgo.RecordPartitioner(kgo.ManualPartitioner()),
	)
	if err != nil {
		return fmt.Errorf("failed to create Kafka client: %w", err)
	}
	defer client.Close()

	res := client.ProduceSync(context.Background(), &kgo.Record{
		Topic: topic, Partition: partition, Value: message})
	if res.FirstErr() != nil {
		return fmt.Errorf("failed to produce message: %w", err)
	}
	producedPartition := res[0].Record.Partition
	if producedPartition != partition {
		return fmt.Errorf("failed to produce message to partition %d, produced to %d", partition, producedPartition)
	}
	return nil
}

func createTestTopic(ctx context.Context, topicName string, numPartitions int32, replicationFactor int16) error {
	cl, err := kgo.NewClient(kgo.SeedBrokers(testKafkaBrokerURL))
	if err != nil {
		return fmt.Errorf("failed to create Kafka client: %w", err)
	}
	defer cl.Close()

	client := kadm.NewClient(cl)
	defer client.Close()

	// Create the topic
	resp, err := client.CreateTopics(ctx, numPartitions, replicationFactor, nil, topicName)
	if err != nil {
		return fmt.Errorf("failed to create topic: %w", err)
	}

	for _, topic := range resp.Sorted() {
		if topic.Err != nil {
			if strings.Contains(topic.Err.Error(), "TOPIC_ALREADY_EXISTS") {
				continue
			}
			return fmt.Errorf("failed to create topic '%s': %v", topic.Topic, topic.Err)
		}
	}
	return nil
}

func waitCh(t *testing.T, ch chan struct{}, timeout time.Duration, failureMessage string) {
	t.Helper()
	select {
	case <-ch:
	case <-time.After(timeout):
		require.Fail(t, failureMessage)
	}
}

func waitAnyCh(t *testing.T, channels []chan struct{}, timeout time.Duration, failureMessage string) {
	t.Helper()

	anyCh := make(chan struct{}, 1)

	for _, ch := range channels {
		go func(c chan struct{}) {
			select {
			case <-c:
				select {
				case anyCh <- struct{}{}:
				default:
				}
			case <-time.After(timeout):
				// let main handle timeout
			}
		}(ch)
	}

	select {
	case <-anyCh:
		// One of the channels received a signal
	case <-time.After(timeout):
		require.Fail(t, failureMessage)
	}
}

func TestKafkaPublisher_ShouldPublishMessageSuccessfully(t *testing.T) {
	t.Parallel()
	// 1. Arrange

	kafkaTopic := "centrifugo_publisher_test_" + uuid.New().String()
	clientId := "testClientId"
	userId := "testUserId"

	// Create Kafka topic
	err := createTestTopic(context.Background(), kafkaTopic, 1, 1)
	require.NoError(t, err)

	// Create payload
	type Score struct {
		GameID     string `json:"gameId"`
		UserID     string `json:"userId"`
		Score      int    `json:"score"`
		TotalScore int    `json:"totalScore"`
		Timestamp  int64  `json:"timestamp"` // epoch time in UnixMicro()
	}

	type WSMessage struct {
		Type string `json:"type"`
		Data Score  `json:"data"`
	}

	sample := WSMessage{
		Type: "score_update",
		Data: Score{
			GameID:     "game1",
			UserID:     clientId,
			Score:      0,
			TotalScore: 0,
			Timestamp:  time.Now().UnixMicro(),
		},
	}
	payload, err := json.Marshal(sample)
	require.NoError(t, err)

	// 2. Act
	err = Publish(clientId, userId, kafkaTopic, payload, nil)

	// 3. Assert
	require.NoError(t, err)
}

func TestClientSend_ShouldPublishMessageSuccessfully(t *testing.T) {
	t.Parallel()
	// 1. Arrange

	kafkaTopic := "centrifugo_publisher_test_" + uuid.New().String()
	clientId := "testClientId"
	userId := "testUserId"

	// Create Kafka topic
	err := createTestTopic(context.Background(), kafkaTopic, 1, 1)
	require.NoError(t, err)

	// Create payload
	type Score struct {
		GameID     string `json:"gameId"`
		UserID     string `json:"userId"`
		Score      int    `json:"score"`
		TotalScore int    `json:"totalScore"`
		Timestamp  int64  `json:"timestamp"` // epoch time in UnixMicro()
	}

	type WSMessage struct {
		Type string `json:"type"`
		Data Score  `json:"data"`
	}

	sample := WSMessage{
		Type: "score_update",
		Data: Score{
			GameID:     "game1",
			UserID:     clientId,
			Score:      0,
			TotalScore: 0,
			Timestamp:  time.Now().UnixMicro(),
		},
	}
	payload, err := json.Marshal(sample)
	require.NoError(t, err)

	// 2. Act
	err = Publish(clientId, userId, kafkaTopic, payload, nil)

	// 3. Assert
	require.NoError(t, err)
}
