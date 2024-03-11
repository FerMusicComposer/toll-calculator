package distancecalculator

import "github.com/confluentinc/confluent-kafka-go/kafka"

type DataConsumer interface {
}

type kafkaconsumer struct {
	consumer *kafka.Consumer
}
