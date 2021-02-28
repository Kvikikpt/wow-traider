package db

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
)

var (
	Pool *pgx.Conn
)

//starting db
func InitDb(databaseUrl string) {
	var dbErr error
	Pool, dbErr = pgx.Connect(context.Background(), databaseUrl)
	if dbErr != nil {
		log.Fatalf("Unable to connect to database: %v", dbErr)
	}
}