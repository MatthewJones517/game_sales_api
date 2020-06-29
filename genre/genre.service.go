package genre

import (
	"encoding/json"
	"log"
	"net/http"
)

func handleGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := getPlatforms()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(genres)

	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(json)
	if err != nil {
		log.Fatal(err)
	}
}

// SetupRoutes prepares the genre package to handle its relevant routes
func SetupRoutes() {
	genresHandler := http.HandlerFunc(handleGenres)

	http.Handle("/genres", genresHandler)
}
