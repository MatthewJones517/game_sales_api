package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MatthewJones517/game_sales_api/database"
	"github.com/MatthewJones517/game_sales_api/game"
	"github.com/MatthewJones517/game_sales_api/genre"
	"github.com/MatthewJones517/game_sales_api/platform"
	"github.com/MatthewJones517/game_sales_api/publisher"
)

func main() {
	database.SetupDatabase()
	platform.SetupRoutes()
	genre.SetupRoutes()
	publisher.SetupRoutes()
	game.SetupRoutes()

	rootHandler := http.HandlerFunc(handleRoot)
	http.Handle("/", rootHandler)

	fmt.Println(("Now listening on port 5050."))
	log.Fatal(http.ListenAndServe(":5050", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><head><title>Game Sales Data API</title></head><body>This is an API meant to be consumed with something like <a href='https://www.postman.com/'>Postman</a>. Check out the <a href='https://github.com/MatthewJones517/game_sales_api'>Github Repo</a> for more information.</body>")
}
