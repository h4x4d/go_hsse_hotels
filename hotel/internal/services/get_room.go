package services

import (
	"context"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/services/errors"
	"github.com/jackc/pgx/v5"
	"os"
)

func GetRoom(roomId int) (*models.Room, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), "db", os.Getenv("POSTGRES_PORT"), "hotel")
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	roomRows, queryErr := conn.Query(context.Background(), "SELECT * FROM rooms WHERE id = $1", roomId)
	if queryErr != nil {
		return nil, queryErr
	}
	room := new(models.Room)
	room.Tags = make([]*models.Tag, 0)
	room.Cost = new(int64)
	room.PersonCount = new(int64)
	room.HotelID = new(int64)

	nextErr := roomRows.Next()
	if nextErr == false {
		return nil, &errors.RoomNotFound{RoomID: roomId}
	}
	scanErr := roomRows.Scan(&room.ID, &room.HotelID, &room.Cost, &room.PersonCount)
	if scanErr != nil {
		return nil, scanErr
	}
	roomRows.Close()

	rowsTags, errQueryTags := conn.Query(context.Background(), "SELECT * FROM tags WHERE room_id = $1", room.ID)
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

	return room, nil
}
