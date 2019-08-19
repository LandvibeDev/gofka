package service

import (
	"fmt"

	"github.com/LandvibeDev/gofka/collector/kafka"
	"github.com/LandvibeDev/gofka/collector/kafka/message"
	ckafka "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

const (
	LogTopic string = "log"
)

type LogServiceInterface interface {
	Send(string, message.LogMessage) error
}

type LogService struct {
	producer *kafka.ProducerConnector
}

func NewLogService(producer *kafka.ProducerConnector) *LogService {
	return &LogService{producer: producer}
}

func (ls *LogService) Send(topic string, msg message.LogMessage) error {
	var deliveryChan = make(chan ckafka.Event)
	defer close(deliveryChan)

	err := ls.producer.Send(topic, msg, deliveryChan)
	if err != nil {
		return err
	}

	e := <-deliveryChan // blocking
	m := e.(*ckafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		return m.TopicPartition.Error
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		return nil
	}
}
