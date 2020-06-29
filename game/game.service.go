package game

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func handleAllGames(w http.ResponseWriter, r *http.Request) {
	// Get URL parameters
	pageNumParam := r.URL.Query()["page"]
	resultsPerPageParam := r.URL.Query()["resultsPerPage"]

	var pageNum int
	var resultsPerPage int
	var err error

	// Check that a valid page number was passed in.
	if len(pageNumParam) == 0 || pageNumParam == nil {
		pageNum = 1
	} else {
		pageNum, err = strconv.Atoi(pageNumParam[0])

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	// Check that a valid results per page was passed in
	if len(resultsPerPageParam) == 0 || resultsPerPageParam == nil {
		resultsPerPage = 20
	} else {
		resultsPerPage, err = strconv.Atoi(resultsPerPageParam[0])

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if resultsPerPage > 50 {
		resultsPerPage = 50
	}

	results, err := getAllGames(resultsPerPage, pageNum)

	handleResults(w, results, err)
}

func handleSingleGame(w http.ResponseWriter, r *http.Request) {
	// Get rank from url
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", "games"))

	if len(urlPathSegments[1:]) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rank, err := strconv.Atoi(urlPathSegments[1])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Make database request
	results, err := getSingleGameByRank(rank)

	// Display to user
	handleResults(w, results, err)
}

func handleResults(w http.ResponseWriter, results []Game, err error) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(results)

	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(json)
	if err != nil {
		log.Fatal(err)
	}
}

// SetupRoutes prepares the games package to handle its relevant routes
func SetupRoutes() {
	allGamesHandler := http.HandlerFunc(handleAllGames)
	singleGameHander := http.HandlerFunc(handleSingleGame)

	http.Handle("/games/all", allGamesHandler)
	http.Handle("/games/", singleGameHander)
}
