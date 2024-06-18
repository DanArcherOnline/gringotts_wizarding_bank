package main

import (
	"database/sql"
	"log"

	"github.com/danarcheronline/gringotts_wizarding_bank/api"
	db "github.com/danarcheronline/gringotts_wizarding_bank/db/sqlc"
	"github.com/danarcheronline/gringotts_wizarding_bank/db/util"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/gringotts_wizarding_bank?sslmode=disable"
	address  = "0.0.0.0:8080"
)

func main() {
	config, err := util.LoadConfig("./")
	if err != nil {
		log.Fatal("Cannot load config.", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
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
