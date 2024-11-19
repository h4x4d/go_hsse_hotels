package services

import (
	"context"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func GetRoomByID(RoomID int64) (*models.Room, error) {
	// connecting to database hotel
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), "db", os.Getenv("POSTGRES_PORT"), "hotel")
	pool, errPool := pgxpool.New(context.Background(), connStr)
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

	// scaning room
	errRoom := roomRow.Scan(&room.ID, room.HotelID, room.Cost, room.PersonCount)
	return room, errRoom
}
