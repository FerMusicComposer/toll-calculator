package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	da "github.com/FerMusicComposer/toll-calculator/distance-aggregator/pkg"
	"github.com/FerMusicComposer/toll-calculator/models"
)

func main() {
	listenAddr := flag.String("listen-addr", ":5050", "aggrgator service online")
	flag.Parse()

	store := da.NewMemoryStore()
	svc := da.NewDistanceAggregator(store)

	makeHTTPTransport(*listenAddr, svc)

}

func makeHTTPTransport(listenAddr string, svc da.Aggregator) {
	fmt.Println("HTTP Transport listening on port", listenAddr)
	http.HandleFunc("/aggregate", handleAggregate(svc))
	http.ListenAndServe(listenAddr, nil)
}

func handleAggregate(svc da.Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var distance models.Distance

		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		if err := svc.AggregateDistance(distance); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
	}
}

func writeJSON(rw http.ResponseWriter, status int, value any) error {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	return json.NewEncoder(rw).Encode(value)
}
