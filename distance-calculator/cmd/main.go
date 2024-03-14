package main

import (
	"log"

	da "github.com/FerMusicComposer/toll-calculator/distance-aggregator/pkg"
	dc "github.com/FerMusicComposer/toll-calculator/distance-calculator/pkg"
)

const aggregatorEndpoint = "http://localhost:5050/aggregate"

func main() {
	svc := dc.NewDistanceCalculator()
	svc = dc.NewLoggerMiddleware(svc)

	kafkaConsumer, err := dc.NewKafkaConsumer(svc, da.NewClient(aggregatorEndpoint))
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
