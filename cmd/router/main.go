package router

import (
	"log"
	"restaurant/Storage/postgres"
	"restaurant/api"
	"restaurant/config"
)

func Router() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	r := api.NewRouter(db)
	r.Run(config.Load().USER_ROUTER)
}
