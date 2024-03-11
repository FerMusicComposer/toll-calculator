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

	prod DataProducer
}

func NewDataReceiver() (*DataReceiver, error) {
	var (
		prod DataProducer
		err  error
	)

	prod, err = NewKafkaProducer()
	if err != nil {
		return nil, err
	}

	prod = NewLogMiddleware(prod)

	return &DataReceiver{
		msgch: make(chan models.OBUData, 128),
		prod:  prod,
	}, nil
}

func (dataRec *DataReceiver) produceData(data models.OBUData) error {
	return dataRec.prod.ProduceData(data)
}

func (dataRec *DataReceiver) websocketReceiveData() {
	fmt.Println("New OBU connected. Receiving data...")
	for {
		var data models.OBUData
		if err := dataRec.conn.ReadJSON(&data); err != nil {
			log.Println("Error reading JSON:", err)
			continue
		}

		err := dataRec.produceData(data)
		if err != nil {
			log.Println("Error producing data:", err)
		}
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
