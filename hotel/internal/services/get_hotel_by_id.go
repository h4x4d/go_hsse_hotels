package services

import (
	"context"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
)

func GetHotelByID(HotelID int64) (*models.Hotel, error) {
	pool, errPool := utils.NewConnection()
	if errPool != nil {
		return nil, errPool
	}
	defer pool.Close()

	hotelRow, errGet := pool.Query(context.Background(),
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
	hotel.Rooms = make([]*models.Room, 0)

	// scaning hotel object
	errHotel := hotelRow.Scan(&hotel.ID, hotel.Name, hotel.City,
		hotel.Address, &hotel.HotelClass)
	return hotel, errHotel
}
