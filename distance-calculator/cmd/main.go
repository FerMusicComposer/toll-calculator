package main

import (
	"log"

	dc "github.com/FerMusicComposer/toll-calculator/distance-calculator/pkg"
)

func main() {
	svc := dc.NewDistanceCalculator()
	svc = dc.NewLoggerMiddleware(svc)

	kafkaConsumer, err := dc.NewKafkaConsumer(svc)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
