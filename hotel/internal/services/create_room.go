package services

import (
	"context"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
	"strings"
)

func CreateRoom(room *models.Room) (*int64, error) {
	pool, err := utils.NewConnection()
	if err != nil {
		return nil, err
	}
	defer pool.Close()

	query := `INSERT INTO rooms`
	fieldNames := []string{}
	fields := []string{}
	values := []interface{}{}

	if room.Cost != nil {
		fieldNames = append(fieldNames, "cost")
		values = append(values, room.Cost)
	}

	if room.HotelID != nil {
		fieldNames = append(fieldNames, "hotel_id")
		values = append(values, room.HotelID)
	}

	if room.PersonCount != nil {
		fieldNames = append(fieldNames, "person_count")
		values = append(values, room.PersonCount)
	}
	if room.ID != 0 {
		fieldNames = append(fieldNames, "id")
		values = append(values, room.ID)
	}

	for ind := 0; ind < len(fieldNames); ind++ {
		fields = append(fields, fmt.Sprintf("$%d", ind+1))
	}
	query += fmt.Sprintf(" (%s) VALUES (%s) RETURNING id", strings.Join(fieldNames, ", "),
		strings.Join(fields, ", "))
	errInsertRoom := pool.QueryRow(context.Background(), query, values...).Scan(&room.ID)
	if errInsertRoom != nil {
		return nil, errInsertRoom
	}
	for _, tag := range room.Tags {
		errCreateTag := CreateTag(room.ID, tag)
		if errCreateTag != nil {
			return nil, errCreateTag
		}
	}
	return &room.ID, nil
}
