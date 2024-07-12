package main

import (
	"log"
	"restaurant/Storage/postgres"
	"restaurant/api"
	"restaurant/config"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	r := api.NewRouter(db)
	r.Run(config.Load().USER_ROUTER)
}
