package main

import (
	"log"
	"net/http"

	dr "github.com/FerMusicComposer/toll-calculator/data-receiver/pkg"
)

func main() {

	dataReceiver, err := dr.NewDataReceiver()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/ws", dataReceiver.WebsocketHandler)
	http.ListenAndServe(":8080", nil)
}
