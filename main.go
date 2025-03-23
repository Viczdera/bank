package main

import (
	"database/sql"
	"log"

	"github.com/Viczdera/bank/api"
	db "github.com/Viczdera/bank/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://root:secret@localhost:8080/s_bank?sslmode=disable"
	serverAddress = "0.0.0.0:9090"
)

func main() {
	//establish connection to db

	connDB, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Could not connect to DB 💿", err)
	}

	store := db.NewStore(connDB)
	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("server no gree start! 😭", err)
	}

}
