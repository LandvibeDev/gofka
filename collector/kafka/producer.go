package kafka

import (
	"github.com/LandvibeDev/gofka/collector/config"
	. "github.com/LandvibeDev/gofka/collector/kafka/message"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func NewProducerConnector(kafkaConfig config.KafkaConfiguration) (*ProducerConnector, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaConfig.Hosts})
	if err != nil {
		return nil, err
	}
	return &ProducerConnector{
		Producer: p,
	}, nil
}

type ProducerConnector struct {
	Producer *kafka.Producer
}

func (p *ProducerConnector) Send(topic string, msg Message, deliveryChan chan kafka.Event) error {
	return p.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msg.GetMessage(),
	}, deliveryChan)
}

func (p *ProducerConnector) SubscribeEvents() chan kafka.Event {
	return p.Producer.Events()
}
func (p *ProducerConnector) Close() {
	p.Producer.Close()
}
