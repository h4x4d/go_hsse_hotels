package services

import (
	"context"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
)

func HotelExists(HotelID int64) (bool, error) {
	pool, errPool := utils.NewConnection()
	if errPool != nil {
		return false, errPool
	}
	defer pool.Close()

	hotelRow, errGet := pool.Query(context.Background(),
		"SELECT id FROM hotels WHERE id = $1", HotelID)
	if errGet != nil {
		return false, errGet
	}
	status := hotelRow.Next()
	hotelRow.Close()
	return status, nil
}
