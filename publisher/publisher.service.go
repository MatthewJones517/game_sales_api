package publisher

import (
	"encoding/json"
	"log"
	"net/http"
)

func handlePublishers(w http.ResponseWriter, r *http.Request) {
	publishers, err := getPublishers()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(publishers)

	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(json)
	if err != nil {
		log.Fatal(err)
	}
}

// SetupRoutes prepares the publisher package to handle its relevant routes
func SetupRoutes() {
	publishersHandler := http.HandlerFunc(handlePublishers)

	http.Handle("/publishers", publishersHandler)
}
