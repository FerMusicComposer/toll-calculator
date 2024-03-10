package main

import (
	"fmt"
	"math/rand"
	"time"
)

const sendInterval = 60

type OBUData struct {
	OBUID     int     `json:"obuID"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func generateCoordinate() float64 {
	intPart := float64(rand.Intn(100) + 1)
	decimalPart := rand.Float64()

	return intPart + decimalPart
}

func main() {
	for {
		fmt.Println("Sending OBU data...")
		fmt.Println(generateCoordinate())
		time.Sleep(sendInterval * time.Second)

	}
}
