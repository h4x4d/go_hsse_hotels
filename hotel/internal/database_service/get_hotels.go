package database_service

import (
	"context"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"strings"
)

func (ds *DatabaseService) GetAll(city *string, hotel_class *int64, name *string) ([]*models.Hotel, error) {
	query := `SELECT * FROM hotels`
	clauses := []string{}
	args := []interface{}{}

	if city != nil {
		clauses = append(clauses, fmt.Sprintf("city = $%d", len(clauses)+1))
		args = append(args, *city)
	}
	if hotel_class != nil {
		clauses = append(clauses, fmt.Sprintf("hotel_class = $%d", len(clauses)+1))
		args = append(args, *hotel_class)
	}
	if name != nil {
		clauses = append(clauses, fmt.Sprintf("NAME LIKE $%d", len(clauses)+1))
		args = append(args, "%"+*name+"%")
	}
	if len(clauses) > 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}

	rowsHotels, errQueryHotels := ds.pool.Query(context.Background(), query, args...)
	if errQueryHotels != nil {
		return nil, errQueryHotels
	}

	// getting hotels information
	result := make([]*models.Hotel, 0)
	for rowsHotels.Next() {
		// init currHotel
		currHotel := new(models.Hotel)
		currHotel.Name = new(string)
		currHotel.City = new(string)
		currHotel.Address = new(string)

		// scaning current hotel
		errHotel := rowsHotels.Scan(&currHotel.ID, currHotel.Name, currHotel.City,
			currHotel.Address, &currHotel.HotelClass, &currHotel.Cost, &currHotel.UserID)
		if errHotel != nil {
			return nil, errHotel
		}
		result = append(result, currHotel)
	}
	rowsHotels.Close()
	return result, nil
}
