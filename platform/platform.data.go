package platform

import (
	"context"
	"log"
	"time"

	"github.com/MatthewJones517/game_sales_api/database"
)

func getPlatforms() ([]Platform, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	results, err := database.DbConn.QueryContext(ctx, `SELECT 
		platform as Platform, 
		sum(global_sales) as Sales 
		FROM gamesales.games 
		GROUP BY Platform 
		ORDER BY Sales desc;`)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer results.Close()

	platforms := make([]Platform, 0)

	for results.Next() {
		var platform Platform
		results.Scan(&platform.Platform, &platform.Sales)

		platforms = append(platforms, platform)
	}

	return platforms, nil
}
