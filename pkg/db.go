package pkg

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func InitDB() *pgx.Conn {
	cnf, err := pgx.ParseConfig(os.Getenv("DATABASE_URI"))
	if err != nil {
		log.Fatal("Error: could not parse db uri")
	}

	db, err := pgx.ConnectConfig(context.Background(), cnf)
	if err != nil {
		log.Fatal("Error: database connection fail")
	}

	return db
}
