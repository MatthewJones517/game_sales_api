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

	processResults(w, results, err)
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
	processResults(w, results, err)
}

func handleGameSearch(w http.ResponseWriter, r *http.Request) {
	// Get query parameters
	name := getSearchURLParameter(r, "name")
	platform := getSearchURLParameter(r, "platform")
	genre := getSearchURLParameter(r, "genre")
	publisher := getSearchURLParameter(r, "publisher")
	minSales := getSearchURLParameter(r, "minSales")
	maxSales := getSearchURLParameter(r, "maxSales")

	if name == "" && platform == "" && genre == "" && publisher == "" && minSales == "" && maxSales == "" {
		handleAllGames(w, r)
	} else {
		results, err := searchForGames(name, platform, genre, publisher, minSales, maxSales)

		processResults(w, results, err)
	}
}

func getSearchURLParameter(r *http.Request, paramName string) string {
	value := r.URL.Query()[paramName]

	if len(value) == 0 {
		return ""
	}

	return value[0]
}

func processResults(w http.ResponseWriter, results []Game, err error) {
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
	searchGamesHandler := http.HandlerFunc(handleGameSearch)

	http.Handle("/games/all", allGamesHandler)
	http.Handle("/games/search", searchGamesHandler)
	http.Handle("/games/", singleGameHander)
}
