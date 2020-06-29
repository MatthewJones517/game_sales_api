package platform

import (
	"encoding/json"
	"log"
	"net/http"
)

func handlePlatforms(w http.ResponseWriter, r *http.Request) {
	platforms, err := getPlatforms()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(platforms)

	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(json)
	if err != nil {
		log.Fatal(err)
	}
}

// SetupRoutes prepares the platform package to handle its relevant routes
func SetupRoutes() {
	platformsHandler := http.HandlerFunc(handlePlatforms)

	http.Handle("/platforms", platformsHandler)
}
