package services

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func DeleteRoomByID(RoomID int64) (*int64, error) {
	// connecting to database hotel
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), "db", os.Getenv("POSTGRES_PORT"), "hotel")
	pool, errPool := pgxpool.New(context.Background(), connStr)
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

	// actually TAGS are deleting TODO
	isTagDeleted, errDeleteTag := DeleteTag(RoomID, nil)
	if errDeleteTag != nil {
		return nil, errDeleteTag
	}
	if !isTagDeleted {
		return nil, nil
	}

	// deleting room itself
	queryDeleted, errDeleteRoom := pool.Query(context.Background(),
		"DELETE FROM rooms WHERE id = $1 RETURNING id", RoomID)
	if errDeleteRoom != nil {
		return nil, errDeleteRoom
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
