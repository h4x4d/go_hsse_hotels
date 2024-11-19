package services

import (
	"context"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"strings"
)

func GetHotels(city *string, hotel_class *int64, name *string, tag *string) ([]*models.Hotel, error) {
	// connecting to database hotel
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), "db", os.Getenv("POSTGRES_PORT"), "hotel")
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	defer pool.Close()

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

	rowsHotels, errQueryHotels := pool.Query(context.Background(), query, args...)
	if errQueryHotels != nil {
		return nil, err
	}

	// getting hotels information
	result := make([]*models.Hotel, 0)
	for rowsHotels.Next() {
		// init currHotel
		currHotel := new(models.Hotel)
		currHotel.Name = new(string)
		currHotel.City = new(string)
		currHotel.Address = new(string)
		currHotel.Rooms = make([]*models.Room, 0)

		// scaning current hotel
		errHotel := rowsHotels.Scan(&currHotel.ID, currHotel.Name, currHotel.City,
			currHotel.Address, &currHotel.HotelClass)
		if errHotel != nil {
			return nil, err
		}

		// getting rooms for current hotel
		currRooms, errRooms := GetRooms(&currHotel.ID, nil)
		if errRooms != nil {
			return nil, errRooms
		}
		currHotel.Rooms = currRooms

		result = append(result, currHotel)
	}
	rowsHotels.Close()
	return result, nil
}
