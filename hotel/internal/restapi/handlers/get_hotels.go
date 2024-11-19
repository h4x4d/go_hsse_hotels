package handlers

import (
	"context"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
	"github.com/jackc/pgx/v5"
	"net/http"
	"os"
)

func GetHotelsHandler(params hotel.GetHotelsParams) (responder middleware.Responder) {
	// catching panic
	defer utils.CatchPanic(&responder)

	// connecting to database hotel
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), "db", os.Getenv("POSTGRES_PORT"), "hotel")
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return middleware.Error(http.StatusInternalServerError, err.Error())
	}
	defer conn.Close(context.Background())

	// getting query for hotels
	rowsHotels, errQueryHotels := conn.Query(context.Background(), "SELECT * FROM hotels")
	if errQueryHotels != nil {
		return middleware.Error(http.StatusInternalServerError, errQueryHotels.Error())
	}

	// getting hotels information
	result := new(hotel.GetHotelsOK)
	result = result.WithPayload(make([]*models.Hotel, 0))
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
			return middleware.Error(http.StatusInternalServerError, errHotel.Error())
		}
		result.Payload = append(result.Payload, currHotel)
	}
	rowsHotels.Close()
	// getting rooms information
	for _, currHotel := range result.Payload {
		// getting query for rooms
		rowsRooms, errQueryRooms := conn.Query(context.Background(),
			"SELECT * FROM rooms WHERE hotel_id = $1", currHotel.ID)
		if errQueryRooms != nil {
			return middleware.Error(http.StatusInternalServerError, errQueryRooms.Error())
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
				return middleware.Error(http.StatusInternalServerError, errRoom.Error())
			}
			currHotel.Rooms = append(currHotel.Rooms, currRoom)
		}
		rowsRooms.Close()
	}
	// getting tags information
	for _, currHotel := range result.Payload {
		for _, currRoom := range currHotel.Rooms {
			// getting query for tags
			rowsTags, errQueryTags := conn.Query(context.Background(),
				"SELECT * FROM tags WHERE room_id = $1", currRoom.ID)
			if errQueryTags != nil {
				return middleware.Error(http.StatusInternalServerError, errQueryTags.Error())
			}
			// reading rooms
			for rowsTags.Next() {
				currTag := new(models.Tag)
				currTag.Name = new(string)
				var currTagId int

				// scanning tag
				errTag := rowsTags.Scan(&currTagId, currTag.Name)
				if errTag != nil {
					return middleware.Error(http.StatusInternalServerError, errTag.Error())
				}
				currRoom.Tags = append(currRoom.Tags, currTag)
			}
			rowsTags.Close()
		}
	}
	return result
}
