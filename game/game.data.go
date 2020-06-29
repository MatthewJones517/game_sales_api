package game

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/MatthewJones517/game_sales_api/database"
)

func getAllGames(resultsPerPage int, page int) ([]Game, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if page == 0 {
		page = 1
	}

	limitStart := resultsPerPage * (page - 1)
	limitEnd := resultsPerPage

	results, err := database.DbConn.QueryContext(ctx, `SELECT 
			games.rank, 
			name, 
			platform, 
			year, 
			genre, 
			publisher, 
			na_sales, 
			eu_sales, 
			jp_sales, 
			other_sales, 
			global_sales
		FROM games
		ORDER BY global_sales desc
		LIMIT ?, ?;`, limitStart, limitEnd)

	return processDatabaseResults(results, err)
}

func getSingleGameByRank(rank int) ([]Game, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	results, err := database.DbConn.QueryContext(ctx, `SELECT 
			games.rank, 
			name, 
			platform, 
			year, 
			genre, 
			publisher, 
			na_sales, 
			eu_sales, 
			jp_sales, 
			other_sales, 
			global_sales
		FROM games
		WHERE games.rank = ?`, rank)

	return processDatabaseResults(results, err)
}

func searchForGames(name string, platform string, genre string, publisher string, minSales string, maxSales string) ([]Game, error) {
	// Build the query string and collect arguments.
	firstConditionAdded := false
	arguments := []interface{}{}

	sql := "SELECT games.rank, name, platform, year, genre, publisher, na_sales, eu_sales, jp_sales, other_sales, global_sales FROM games WHERE "

	addSearchConditional("name", name, "=", &firstConditionAdded, &sql, &arguments)
	addSearchConditional("platform", platform, "=", &firstConditionAdded, &sql, &arguments)
	addSearchConditional("genre", genre, "=", &firstConditionAdded, &sql, &arguments)
	addSearchConditional("publisher", publisher, "=", &firstConditionAdded, &sql, &arguments)
	addSearchConditional("global_sales", minSales, ">", &firstConditionAdded, &sql, &arguments)
	addSearchConditional("global_sales", maxSales, "<", &firstConditionAdded, &sql, &arguments)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	results, err := database.DbConn.QueryContext(ctx, sql, arguments...)

	return processDatabaseResults(results, err)
}

func addSearchConditional(argumentName string, argument string, operator string, firstConditionAdded *bool, sql *string, arguments *[]interface{}) {
	if argument != "" {
		if *firstConditionAdded {
			*sql += " AND "
		}
		*sql += argumentName + " " + operator + " ?"
		*arguments = append(*arguments, argument)

		*firstConditionAdded = true
	}
}

func processDatabaseResults(results *sql.Rows, err error) ([]Game, error) {
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer results.Close()

	games := make([]Game, 0)

	for results.Next() {
		var game Game
		results.Scan(&game.Rank, &game.Name, &game.Platform, &game.Year, &game.Genre, &game.Publisher, &game.NASales, &game.EUSales, &game.JPSales, &game.OtherSales, &game.GlobalSales)

		games = append(games, game)
	}

	return games, nil
}
