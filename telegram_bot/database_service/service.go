package database_service

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

type DatabaseService struct {
	pool *pgxpool.Pool
}

func NewDatabaseService() (*DatabaseService, error) {
	// in the future 127.0.0.1 -> db
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), "db", os.Getenv("POSTGRES_PORT"), os.Getenv("TELEGRAM_DB_NAME"))

	result := new(DatabaseService)
	newPool, errPool := pgxpool.New(context.Background(), connStr)
	if errPool != nil {
		return nil, errPool
	}
	result.pool = newPool
	return result, nil
}
