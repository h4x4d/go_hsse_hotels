package database_service

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

type DatabaseService struct {
	pool *pgxpool.Pool
}

func GetDatabaseServiceFromContext(ctx context.Context) (*DatabaseService, error) {
	contextValue := ctx.Value("DatabaseService")
	if contextValue == nil {
		return nil, errors.New("Database service was not found in the context")
	}

	databaseService, castResult := contextValue.(*DatabaseService)
	if castResult == false {
		return nil, errors.New("Casting to database service failed")
	}
	return databaseService, nil
}

func ContextWithDatabaseService(databaseService *DatabaseService) context.Context {
	return context.WithValue(context.Background(), "DatabaseService", databaseService)
}

func Make() (*DatabaseService, error) {
	result := new(DatabaseService)
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), "db", os.Getenv("POSTGRES_PORT"), os.Getenv("HOTEL_DB_NAME"))
	newPool, errPool := pgxpool.New(context.Background(), connStr)
	if errPool != nil {
		return nil, errPool
	}
	result.pool = newPool
	return result, nil
}
