package services

import (
	"context"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
)

func DeleteHotelByID(HotelID int64) (*int64, error) {
	pool, errPool := utils.NewConnection()
	if errPool != nil {
		return nil, errPool
	}
	defer pool.Close()

	// checking for existence current hotel
	hotelExists, errExistence := HotelExists(HotelID)
	if errExistence != nil {
		return nil, errExistence
	}
	if !hotelExists {
		return nil, nil
	}

	// deleting rooms
	roomsToDelete, errRooms := GetRooms(&HotelID, nil)
	if errRooms != nil {
		return nil, errRooms
	}
	for _, room := range roomsToDelete {
		deletedRoomID, errDeleteRoom := DeleteRoomByID(room.ID)
		if errDeleteRoom != nil {
			return nil, errDeleteRoom
		}
		if deletedRoomID == nil {
			return nil, nil
		}
	}

	// delete information from booking TODO

	// deleting hotel
	queryDeleted, errDeleteHotel := pool.Query(context.Background(),
		"DELETE FROM hotels WHERE id = $1 RETURNING id", HotelID)
	if errDeleteHotel != nil {
		return nil, errDeleteHotel
	}
	if !queryDeleted.Next() {
		return nil, nil
	}
	deletedId := new(int64)
	errScan := queryDeleted.Scan(&deletedId)
	if errScan != nil {
		return nil, errScan
	}
	return deletedId, nil
}
