package services

import (
	"context"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"strings"
)

func GetRooms(HotelID *int64, Tag *string) ([]*models.Room, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), "db", os.Getenv("POSTGRES_PORT"), "hotel")
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	defer pool.Close()

	query := `SELECT * FROM rooms`
	clauses := []string{}
	args := []interface{}{}

	if HotelID != nil {
		clauses = append(clauses, fmt.Sprintf("hotel_id = $%d", len(clauses)+1))
		args = append(args, HotelID)
	}
	if Tag != nil {
		clause := fmt.Sprintf("$%d IN (SELECT tag FROM tags WHERE room_id = rooms.id)", len(clauses)+1)
		clauses = append(clauses, clause)
		args = append(args, Tag)
	}
	if len(clauses) > 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}

	rowsRooms, errQueryRooms := pool.Query(context.Background(), query, args...)
	if errQueryRooms != nil {
		return nil, err
	}

	// getting rooms information
	result := make([]*models.Room, 0)
	for rowsRooms.Next() {
		// init currRoom
		currRoom := new(models.Room)
		currRoom.Cost = new(int64)
		currRoom.HotelID = new(int64)
		currRoom.PersonCount = new(int64)
		currRoom.Tags = make([]*models.Tag, 0)

		// scanning currRoom
		errScan := rowsRooms.Scan(&currRoom.ID, currRoom.HotelID, currRoom.Cost, currRoom.PersonCount)
		if errScan != nil {
			return nil, err
		}

		// getting tags for current room
		currTags, errTags := GetTags(&currRoom.ID)
		if errTags != nil {
			return nil, err
		}
		currRoom.Tags = currTags
		result = append(result, currRoom)
	}
	rowsRooms.Close()
	return result, nil
}
