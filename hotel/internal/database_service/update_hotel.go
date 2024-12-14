package database_service

import (
	"context"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"strings"
)

func (ds *DatabaseService) Update(id int64, hotel *models.Hotel) (*models.Hotel, error) {
	query := `UPDATE hotels SET`
	var fieldNames []string
	var values []interface{}

	if hotel.Address != nil {
		fieldNames = append(fieldNames, fmt.Sprintf("address = $%d", len(values)+1))
		values = append(values, *hotel.Address)
	}
	if hotel.City != nil {
		fieldNames = append(fieldNames, fmt.Sprintf("city = $%d", len(values)+1))
		values = append(values, *hotel.City)
	}
	if hotel.Name != nil {
		fieldNames = append(fieldNames, fmt.Sprintf("name = $%d", len(values)+1))
		values = append(values, *hotel.Name)
	}
	if hotel.Cost != 0 {
		fieldNames = append(fieldNames, fmt.Sprintf("cost = $%d", len(values)+1))
		values = append(values, hotel.Cost)
	}
	if hotel.ID != 0 {
		fieldNames = append(fieldNames, fmt.Sprintf("id = $%d", len(values)+1))
		values = append(values, hotel.ID)
	}
	if hotel.HotelClass != 0 {
		fieldNames = append(fieldNames, fmt.Sprintf("hotel_class = $%d", len(values)+1))
		values = append(values, hotel.HotelClass)
	}
	query += fmt.Sprintf(" %s WHERE %s RETURNING *", strings.Join(fieldNames, ", "),
		fmt.Sprintf("id = $%d", len(values)+1))
	values = append(values, id)
	fmt.Println(query, values)
	errGetHotel := ds.pool.QueryRow(context.Background(), query, values...).Scan(&hotel.ID, hotel.Name,
		hotel.City, hotel.Address, &hotel.HotelClass, &hotel.Cost, &hotel.UserID)
	return hotel, errGetHotel
}
