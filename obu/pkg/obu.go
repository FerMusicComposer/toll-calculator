package obu

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

const sendInterval = 2

type OBUData struct {
	OBUID     int     `json:"obuID"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func GenerateCoordinate() float64 {
	intPart := float64(rand.Intn(100) + 1)
	decimalPart := rand.Float64()

	return intPart + decimalPart
}

func GenerateLocation() (float64, float64) {
	return GenerateCoordinate(), GenerateCoordinate()
}

func GenerateOBUIDs(numOfIDs int) []int {
	ids := make([]int, numOfIDs)
	for i := 0; i < numOfIDs; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}

	return ids
}

func GenerateData(obuIDs []int, conn *websocket.Conn) {
	for i := 0; i < len(obuIDs); i++ {
		lat, long := GenerateLocation()
		obu := OBUData{
			OBUID:     obuIDs[i],
			Latitude:  lat,
			Longitude: long,
		}

		fmt.Printf("%+v\n", obu)

		if err := conn.WriteJSON(obu); err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Duration(sendInterval) * time.Second)
	}
}
