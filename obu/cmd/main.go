package main

import (
	"log"

	"github.com/FerMusicComposer/toll-calculator/obu/pkg/obu"
	"github.com/gorilla/websocket"
)

func main() {
	obuIDs := obu.GenerateOBUIDs(20)
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		obu.GenerateData(obuIDs, conn)
	}

}
