package datareceiver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FerMusicComposer/toll-calculator/models"
	"github.com/gorilla/websocket"
)

type DataReceiver struct {
	msgch chan models.OBUData
	conn  *websocket.Conn
}

func NewDataReceiver() *DataReceiver {
	return &DataReceiver{
		msgch: make(chan models.OBUData, 128),
	}
}

func (dataRec *DataReceiver) websocketReceiveData() {
	fmt.Println("New OBU connected. Receiving data...")
	for {
		var data models.OBUData
		if err := dataRec.conn.ReadJSON(&data); err != nil {
			log.Println("Error reading JSON:", err)
			continue
		}
		fmt.Printf("Received OBU data from [%d] :: <lat: %.2f, long: %.2f>\n", data.OBUID, data.Latitude, data.Longitude)
		dataRec.msgch <- data

	}
}

func (dataRec *DataReceiver) WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	dataRec.conn = conn

	go dataRec.websocketReceiveData()
}
