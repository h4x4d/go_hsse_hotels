package services

import (
	"context"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"strings"
)

func CreateHotel(hotel *models.Hotel) error {
	// connecting to database hotel
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), "db", os.Getenv("POSTGRES_PORT"), "hotel")
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return err
	}
	defer pool.Close()

	query := `INSERT INTO hotels`
	fieldNames := []string{}
	fields := []string{}
	values := []interface{}{}

	if hotel.Address != nil {
		fieldNames = append(fieldNames, "address")
		values = append(values, hotel.Address)
	}
	if hotel.City != nil {
		fieldNames = append(fieldNames, "city")
		values = append(values, hotel.City)
	}
	if hotel.Name != nil {
		fieldNames = append(fieldNames, "name")
		values = append(values, hotel.Name)
	}

	fieldNames = append(fieldNames, "id")
	values = append(values, hotel.ID)
	fieldNames = append(fieldNames, "hotel_class")
	values = append(values, hotel.HotelClass)

	for ind := 0; ind < len(fieldNames); ind++ {
		fields = append(fields, fmt.Sprintf("$%d", ind+1))
	}
	query += fmt.Sprintf(" (%s) VALUES (%s)", strings.Join(fieldNames, ", "),
		strings.Join(fields, ", "))

	for _, room := range hotel.Rooms {
		errCreateRoom := CreateRoom(room)
		if errCreateRoom != nil {
			return errCreateRoom
		}
	}

	_, errInsertHotel := pool.Exec(context.Background(), query, values...)
	return errInsertHotel
}
