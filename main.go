package main

import (
	"context"
	"log"

	"github.com/akoqassym/simplebank/api"
	"github.com/akoqassym/simplebank/config"
	db "github.com/akoqassym/simplebank/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/lib/pq"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load the config:", err)
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
