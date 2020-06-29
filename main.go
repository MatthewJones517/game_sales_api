package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MatthewJones517/game_sales_api/database"
	"github.com/MatthewJones517/game_sales_api/platform"
)

func main() {
	database.SetupDatabase()
	platform.SetupRoutes()

	fmt.Println(("Now listening on port 5050."))
	log.Fatal(http.ListenAndServe(":5050", nil))
}
