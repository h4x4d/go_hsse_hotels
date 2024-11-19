package services

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func HotelExists(HotelID int64) (bool, error) {
	// connecting to database hotel
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), "db", os.Getenv("POSTGRES_PORT"), "hotel")
	pool, errPool := pgxpool.New(context.Background(), connStr)
	if errPool != nil {
		return false, errPool
	}
	defer pool.Close()

	hotelRow, errGet := pool.Query(context.Background(),
		"SELECT id FROM hotels WHERE id = $1", HotelID)
	if errGet != nil {
		return false, errGet
	}
	if !hotelRow.Next() {
		return true, nil
	}
	return false, nil
}
