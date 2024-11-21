package services

import (
	"context"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
)

func DeleteRoomByID(RoomID int64) (*int64, error) {
	// connecting to database hotel
	pool, errPool := utils.NewConnection()
	if errPool != nil {
		return nil, errPool
	}
	defer pool.Close()

	// checking for existence current hotel
	roomExists, errExistence := RoomExists(RoomID)
	if errExistence != nil {
		return nil, errExistence
	}
	if !roomExists {
		return nil, nil
	}

	// deleting room itself
	deletedId := new(int64)
	errDeleteRoom := pool.QueryRow(context.Background(),
		"DELETE FROM rooms WHERE id = $1 RETURNING id", RoomID).Scan(&deletedId)
	if errDeleteRoom != nil {
		return nil, errDeleteRoom
	}
	return deletedId, nil
}
