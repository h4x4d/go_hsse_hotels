package services

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func DeleteHotelByID(HotelID int64) (*int64, error) {
	// connecting to database
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), "db", os.Getenv("POSTGRES_PORT"), "hotel")
	pool, errPool := pgxpool.New(context.Background(), connStr)
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
