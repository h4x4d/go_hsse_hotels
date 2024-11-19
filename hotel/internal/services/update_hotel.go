package services

import (
	"context"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"strings"
)

func UpdateHotel(hotel *models.Hotel) error {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), "db", os.Getenv("POSTGRES_PORT"), "hotel")
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return err
	}
	defer pool.Close()

	query := "SELECT * FROM hotels"
	settings := []string{}
	args := []interface{}{}

	if hotel.City != nil {
		settings = append(settings, fmt.Sprintf("city = $%d", len(settings)+1))
		args = append(args, *hotel.City)
	}
	if hotel.Address != nil {
		settings = append(settings, fmt.Sprintf("address = $%d", len(settings)+1))
		args = append(args, *hotel.Address)
	}
	if hotel.Name != nil {
		settings = append(settings, fmt.Sprintf("name = $%d", len(settings)+1))
		args = append(args, *hotel.Name)
	}
	settings = append(settings, fmt.Sprintf("id = $%d", len(settings)+1))
	args = append(args, hotel.ID)
	settings = append(settings, fmt.Sprintf("hotel_class = $%d", len(settings)+1))
	args = append(args, hotel.HotelClass)

	if len(clauses) > 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}

	_, errUpdateHotel := pool.Exec(context.Background(), query, args...)
	return errUpdateHotel
}
