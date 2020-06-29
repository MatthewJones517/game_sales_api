package game

import (
	"context"
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
