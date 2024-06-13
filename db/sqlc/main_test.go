package db

import (
	// "context"
	"database/sql"
	"log"
	"os"
	"testing"

	// "github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

const (
	dbSource = "postgresql://root:secret@localhost:5432/gringotts_wizarding_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open("postgres", dbSource)

	if err != nil {
		log.Fatal("Cannot connect to database.", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
