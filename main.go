package main

import (
	"database/sql"
	"log"

	"github.com/danarcheronline/gringotts_wizarding_bank/api"
	db "github.com/danarcheronline/gringotts_wizarding_bank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/gringotts_wizarding_bank?sslmode=disable"
	address  = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to database.", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(address)
	if err != nil {
		log.Fatal("Cannot start server.", err)
	}
}
