package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/LandvibeDev/gofka/collector/config"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func GetProducer(kafkaConfig config.KafkaConfiguration) (*ProducerConnector, error) {
	return NewProducerConnector(kafkaConfig)
}

func EnsureTopic(kafkaConfig config.KafkaConfiguration) ([]kafka.TopicResult, error) {
	admin, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": kafkaConfig.Hosts})
	if err != nil {
		return nil, err
	}
	defer admin.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timeout, err := time.ParseDuration("60s")
	if err != nil {
		return nil, err
	}

	results, err := admin.CreateTopics(
		ctx,
		[]kafka.TopicSpecification{{
			Topic:             kafkaConfig.Topic.Name,
			NumPartitions:     kafkaConfig.Topic.NumPartitions,
			ReplicationFactor: kafkaConfig.Topic.ReplicationFactor}},
		kafka.SetAdminOperationTimeout(timeout),
	)

	if err != nil {
		fmt.Printf("Failed to create topic: $d\n", kafkaConfig.Topic.Name, err)
		return nil, err
	}

	return results, nil
}
