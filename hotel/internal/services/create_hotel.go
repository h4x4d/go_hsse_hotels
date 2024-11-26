package services

import (
	"context"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
	"strings"
)

func CreateHotel(hotel *models.Hotel) (*int64, error) {
	pool, err := utils.NewConnection()
	if err != nil {
		return nil, err
	}
	defer pool.Close()

	query := `INSERT INTO hotels`
	var fieldNames []string
	var fields []string
	var values []interface{}

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
	if hotel.ID != 0 {
		fieldNames = append(fieldNames, "id")
		values = append(values, hotel.ID)
	}
	fieldNames = append(fieldNames, "hotel_class")
	values = append(values, hotel.HotelClass)

	for ind := 0; ind < len(fieldNames); ind++ {
		fields = append(fields, fmt.Sprintf("$%d", ind+1))
	}
	query += fmt.Sprintf(" (%s) VALUES (%s) RETURNING id", strings.Join(fieldNames, ", "),
		strings.Join(fields, ", "))
	errInsertHotel := pool.QueryRow(context.Background(), query, values...).Scan(&hotel.ID)
	if errInsertHotel != nil {
		return nil, errInsertHotel
	}

	for _, room := range hotel.Rooms {
		room.HotelID = &hotel.ID
		_, errCreateRoom := CreateRoom(room)
		if errCreateRoom != nil {
			return nil, errCreateRoom
		}
	}

	return &hotel.ID, errInsertHotel
}