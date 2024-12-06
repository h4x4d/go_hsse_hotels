package database_service

import (
	"context"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
)

func (ds *DatabaseService) GetById(HotelID int64) (*models.Hotel, error) {
	hotelRow, errGet := ds.pool.Query(context.Background(),
		"SELECT * FROM hotels WHERE id = $1", HotelID)
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

	// scaning hotel object
	errHotel := hotelRow.Scan(&hotel.ID, hotel.Name, hotel.City,
		hotel.Address, &hotel.HotelClass, &hotel.Cost, &hotel.UserID)
	hotelRow.Close()

	return hotel, errHotel
}
