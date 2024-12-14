package database_service

import (
	"context"
)

func (ds *DatabaseService) Exists(HotelID int64) (bool, error) {
	hotelRow, errGet := ds.pool.Query(context.Background(),
		"SELECT id FROM hotels WHERE id = $1", HotelID)
	if errGet != nil {
		return false, errGet
	}
	status := hotelRow.Next()
	hotelRow.Close()
	return status, nil
}
