package distancecalculator

import (
	"encoding/json"
	"fmt"
	"time"

	da "github.com/FerMusicComposer/toll-calculator/distance-aggregator/pkg"
	"github.com/FerMusicComposer/toll-calculator/models"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const kafkaTopic = "obu-data"

type kafkaconsumer struct {
	consumer          *kafka.Consumer
	isRunning         bool
	calculatorService DistCalculator
	aggClient         *da.Client
}

func NewKafkaConsumer(svc DistCalculator, aggClient *da.Client) (*kafkaconsumer, error) {
	topic := kafkaTopic

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	consumer.SubscribeTopics([]string{topic}, nil)

	return &kafkaconsumer{
		consumer:          consumer,
		calculatorService: svc,
		aggClient:         aggClient,
	}, nil
}

func (kafkaCons *kafkaconsumer) ConsumeData() error {
	for kafkaCons.isRunning {
		msg, err := kafkaCons.consumer.ReadMessage(-1)
		if err != nil {
			if err.(kafka.Error).IsTimeout() {
				fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			}

			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			continue
		}

		var data models.OBUData
		if err := json.Unmarshal(msg.Value, &data); err != nil {
			fmt.Printf("Error unmarshalling data: %v\n", err)
			continue
		}

		distance, err := kafkaCons.calculatorService.CalculateDistance(data)
		if err != nil {
			fmt.Printf("Error calculating distance: %v\n", err)
			continue
		}

		req := models.Distance{
			Value: distance,
			OBUID: data.OBUID,
			Unix:  time.Now().UnixNano(),
		}

		if err := kafkaCons.aggClient.AggregateDistance(req); err != nil {
			fmt.Printf("Error aggregating distance: %v\n", err)
			continue
		}

		fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		fmt.Printf("Distance: %.2f\n", distance)
	}

	kafkaCons.consumer.Close()
	return nil
}

func (kafkaCons *kafkaconsumer) Start() {
	kafkaCons.isRunning = true
	kafkaCons.ConsumeData()
}
