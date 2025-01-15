package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB() *pgxpool.Pool {
	connString := os.Getenv("DATABASE_URL")
	pool, err := pgxpool.New(context.Background(), connString)

	if err != nil {
		log.Fatalf("Failed to connect to db %v\n", err)
		os.Exit(1)
	}

	return pool
}
