package services

import (
	"context"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/jackc/pgx/v5"
	"os"
)

func CreateHotel(hotel *models.Hotel) error {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), "db", os.Getenv("POSTGRES_PORT"), "hotel")
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	_, errInsertHotel := conn.Exec(context.Background(),
		"INSERT INTO hotels (id, name, city, address, hotel_class) VALUES ($1, $2, $3, $4, $5)",
		hotel.ID, hotel.Name, hotel.City, hotel.Address, hotel.HotelClass)
	return errInsertHotel
}
