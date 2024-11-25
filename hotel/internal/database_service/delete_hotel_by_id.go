package database_service

import (
	"context"
)

func (ds *DatabaseService) DeleteByID(HotelID int64) (*int64, error) {
	// checking for existence current hotel
	hotelExists, errExistence := ds.Exists(HotelID)
	if errExistence != nil {
		return nil, errExistence
	}
	if !hotelExists {
		return nil, nil
	}

	// delete information from booking TODO

	// deleting hotel
	deletedId := new(int64)
	errDeleteHotel := ds.pool.QueryRow(context.Background(),
		"DELETE FROM hotels WHERE id = $1 RETURNING id", HotelID).Scan(&deletedId)
	if errDeleteHotel != nil {
		return nil, errDeleteHotel
	}
	return deletedId, nil
}
