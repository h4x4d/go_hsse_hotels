package services

import (
	"context"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
	"strings"
)

func CreateRoom(room *models.Room) error {
	pool, err := utils.NewConnection()
	if err != nil {
		return err
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

	fieldNames = append(fieldNames, "id")
	values = append(values, room.ID)

	for ind := 0; ind < len(fieldNames); ind++ {
		fields = append(fields, fmt.Sprintf("$%d", ind+1))
	}
	query += fmt.Sprintf(" (%s) VALUES (%s)", strings.Join(fieldNames, ", "),
		strings.Join(fields, ", "))

	for _, tag := range room.Tags {
		errCreateTag := CreateTag(room.ID, tag)
		if errCreateTag != nil {
			return errCreateTag
		}
	}

	_, errInsertRoom := pool.Exec(context.Background(), query, values...)
	return errInsertRoom
}
