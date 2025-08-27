package main

import (
	"database/sql"
	"log"

	"github.com/Viczdera/bank/api"
	db "github.com/Viczdera/bank/db/sqlc"
	"github.com/Viczdera/bank/util"

	_ "github.com/lib/pq"
)

func main() {
	//load config from env
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load config 💿", err)
	}

	//establish connection to db
	connDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Could not connect to DB 💿", err)
	}

	store := db.NewStore(connDB)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("server no gree start! 😭", err)
	}

}
