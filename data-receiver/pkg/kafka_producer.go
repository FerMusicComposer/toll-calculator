package datareceiver

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/FerMusicComposer/toll-calculator/models"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const kafkaTopic string = "obu-data"

type DataProducer interface {
	ProduceData(data models.OBUData) error
}

type kafkaproducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer() (*kafkaproducer, error) {
	prod, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		return nil, err
	}

	// Starts new goroutine to check if data has been delivered
	go func() {
		for event := range prod.Events() {
			switch ev := event.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	return &kafkaproducer{producer: prod}, nil
}

func (kafkaProd *kafkaproducer) ProduceData(data models.OBUData) error {
	topic := kafkaTopic
	b, err := json.Marshal(data)
	if err != nil {
		log.Println("Error marshalling data:", err)
		return err
	}

	err = kafkaProd.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: int32(kafka.PartitionAny),
		},
		Value: b,
	}, nil)

	if err != nil {
		log.Println("Error producing data:", err)
		return err
	}

	return nil

}
