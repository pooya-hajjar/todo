package models

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
)

var PostgresDB pgx.Conn

func ConnectToPostgres(wg *sync.WaitGroup) {
	defer wg.Done()

	pgDsn := os.Getenv("POSTGRES_DSN")

	db, err := pgx.Connect(context.Background(), pgDsn)
	if err != nil {
		log.Fatal("db connection error !")
	}

	PostgresDB = *db
}
