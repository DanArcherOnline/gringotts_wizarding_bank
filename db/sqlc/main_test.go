package db

import (
	// "context"
	"database/sql"
	"log"
	"os"
	"testing"

	// "github.com/jackc/pgx/v5/pgxpool"
	"github.com/danarcheronline/gringotts_wizarding_bank/db/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	// var err error
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config.", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database.", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
