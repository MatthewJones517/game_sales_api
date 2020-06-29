package publisher

import (
	"context"
	"log"
	"time"

	"github.com/MatthewJones517/game_sales_api/database"
)

func getPublishers() ([]Publisher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	results, err := database.DbConn.QueryContext(ctx, `SELECT 
		publisher as Publisher, 
		sum(global_sales) as Sales 
		FROM gamesales.games 
		GROUP BY Publisher 
		ORDER BY Sales desc;`)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer results.Close()

	publishers := make([]Publisher, 0)

	for results.Next() {
		var publisher Publisher
		results.Scan(&publisher.Publisher, &publisher.Sales)

		publishers = append(publishers, publisher)
	}

	return publishers, nil
}
