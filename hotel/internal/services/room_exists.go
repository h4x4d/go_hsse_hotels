package services

import (
	"context"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
)

func RoomExists(RoomID int64) (bool, error) {
	pool, errPool := utils.NewConnection()
	if errPool != nil {
		return false, errPool
	}
	defer pool.Close()

	roomRow, errGet := pool.Query(context.Background(),
		"SELECT id FROM rooms WHERE id = $1", RoomID)
	if errGet != nil {
		return false, errGet
	}
	status := roomRow.Next()
	roomRow.Close()
	return status, nil
}
