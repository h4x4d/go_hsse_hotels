package services

import (
	"context"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
)

func GetRoomByID(RoomID int64) (*models.Room, error) {
	pool, errPool := utils.NewConnection()
	if errPool != nil {
		return nil, errPool
	}
	defer pool.Close()

	roomRow, errGet := pool.Query(context.Background(),
		"SELECT * FROM rooms WHERE id = $1", RoomID)
	if errGet != nil {
		return nil, errGet
	}
	if !roomRow.Next() {
		return nil, nil
	}

	room := new(models.Room)
	room.PersonCount = new(int64)
	room.Cost = new(int64)
	room.HotelID = new(int64)
	room.Tags = make([]*models.Tag, 0)

	// scaning room
	errRoom := roomRow.Scan(&room.ID, &room.HotelID, &room.Cost, &room.PersonCount)
	roomRow.Close()

	rowsTags, errQueryTags := pool.Query(context.Background(), "SELECT * FROM tags WHERE room_id = $1", room.ID)
	if errQueryTags != nil {
		return nil, errQueryTags
	}

	for rowsTags.Next() {
		currTag := new(models.Tag)
		currTag.Name = new(string)
		var currTagId int

		// scanning tag
		errTag := rowsTags.Scan(&currTagId, currTag.Name)
		if errTag != nil {
			return nil, errTag
		}
		room.Tags = append(room.Tags, currTag)
	}
	rowsTags.Close()

	return room, errRoom
}
