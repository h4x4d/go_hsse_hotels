package services

import (
	"context"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func GetHotelByID(HotelID int64) (*models.Hotel, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), "db", os.Getenv("POSTGRES_PORT"), "hotel")
	pool, errPool := pgxpool.New(context.Background(), connStr)
	if errPool != nil {
		return nil, errPool
	}
	defer pool.Close()

	hotelRow, errGet := pool.Query(context.Background(),
		"SELECT FROM hotels WHERE id = $1", HotelID)
	if errGet != nil {
		return nil, errGet
	}
	if !hotelRow.Next() {
		return nil, nil
	}

	hotel := new(models.Hotel)
	hotel.Name = new(string)
	hotel.City = new(string)
	hotel.Address = new(string)
	hotel.Rooms = make([]*models.Room, 0)

	// scaning hotel
	errHotel := hotelRow.Scan(&hotel.ID, hotel.Name, hotel.City,
		hotel.Address, &hotel.HotelClass)
	return hotel, errHotel
}
