package main

import (
	"net/http"

	dataRec "github.com/FerMusicComposer/toll-calculator/data-receiver/pkg"
)

func main() {
	dataReceiver := dataRec.NewDataReceiver()

	http.HandleFunc("/ws", dataReceiver.WebsocketHandler)
	http.ListenAndServe(":8080", nil)
}
