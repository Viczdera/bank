package db

//same package as crud code
//main_test used to set up connection with query object

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:secret@localhost:8080/s_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

// testing
func TestMain(t *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Could not connect to DB", err)
	}

	//use connection to create new test queries object
	testQueries = New(testDB)

	//run and report back to test runner via the o.exit command
	os.Exit(t.Run())

}
