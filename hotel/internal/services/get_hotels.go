package services

import (
	"context"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/jackc/pgx/v5"
	"os"
)

func GetHotels(city *string, hotel_class *int64, name *string, tag *string) ([]*models.Hotel, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), "db", os.Getenv("POSTGRES_PORT"), "hotel")
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	// getting query for hotels
	rowsHotels, errQueryHotels := conn.Query(context.Background(), "SELECT * FROM hotels")
	if errQueryHotels != nil {
		return nil, err
	}

	// getting hotels information
	result := make([]*models.Hotel, 0)
	for rowsHotels.Next() {
		// init currHotel
		currHotel := new(models.Hotel)
		currHotel.Name = new(string)
		currHotel.City = new(string)
		currHotel.Address = new(string)
		currHotel.Rooms = make([]*models.Room, 0)

		// scaning currHotel
		errHotel := rowsHotels.Scan(&currHotel.ID, currHotel.Name, currHotel.City,
			currHotel.Address, &currHotel.HotelClass)
		if errHotel != nil {
			return nil, err
		}
		result = append(result, currHotel)
	}
	rowsHotels.Close()
	// getting rooms information
	for _, currHotel := range result {
		// getting query for rooms
		rowsRooms, errQueryRooms := conn.Query(context.Background(),
			"SELECT * FROM rooms WHERE hotel_id = $1", currHotel.ID)
		if errQueryRooms != nil {
			return nil, err
		}
		// reading rooms
		for rowsRooms.Next() {
			currRoom := new(models.Room)
			currRoom.Cost = new(int64)
			currRoom.HotelID = new(int64)
			currRoom.PersonCount = new(int64)

			// scanning room
			errRoom := rowsRooms.Scan(&currRoom.ID, currRoom.HotelID, currRoom.Cost, currRoom.PersonCount)
			if errRoom != nil {
				return nil, err
			}
			currHotel.Rooms = append(currHotel.Rooms, currRoom)
		}
		rowsRooms.Close()
	}
	// getting tags information
	for _, currHotel := range result {
		for _, currRoom := range currHotel.Rooms {
			// getting query for tags
			rowsTags, errQueryTags := conn.Query(context.Background(),
				"SELECT * FROM tags WHERE room_id = $1", currRoom.ID)
			if errQueryTags != nil {
				return nil, err
			}
			// reading rooms
			for rowsTags.Next() {
				currTag := new(models.Tag)
				currTag.Name = new(string)
				var currTagId int

				// scanning tag
				errTag := rowsTags.Scan(&currTagId, currTag.Name)
				if errTag != nil {
					return nil, err
				}
				currRoom.Tags = append(currRoom.Tags, currTag)
			}
			rowsTags.Close()
		}
	}
	return result, nil
}
