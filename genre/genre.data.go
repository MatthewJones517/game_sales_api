package genre

import (
	"context"
	"log"
	"time"

	"github.com/MatthewJones517/game_sales_api/database"
)

func getPlatforms() ([]Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	results, err := database.DbConn.QueryContext(ctx, `SELECT 
		genre as Genre, 
		sum(global_sales) as Sales 
		FROM gamesales.games 
		GROUP BY Genre 
		ORDER BY Sales desc;`)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer results.Close()

	genres := make([]Genre, 0)

	for results.Next() {
		var genre Genre
		results.Scan(&genre.Genre, &genre.Sales)

		genres = append(genres, genre)
	}

	return genres, nil
}
