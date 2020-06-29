package game

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func handleAllGames(w http.ResponseWriter, r *http.Request) {
	// Get URL parameters
	pageNumParam := r.URL.Query()["page"]
	resultsPerPageParam := r.URL.Query()["resultsPerPage"]
	orderByParam := r.URL.Query()["orderby"]

	var pageNum int
	var resultsPerPage int
	var orderBy string
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

	// Check that a valid orderBy was passed in.
	if len(orderByParam) == 0 || orderByParam == nil {
		orderBy = "asc"
	} else {
		orderBy = strings.ToLower(orderByParam[0])

		if orderBy != "asc" && orderBy != "desc" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	results, err := getAllGames(resultsPerPage, pageNum, orderBy)

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
