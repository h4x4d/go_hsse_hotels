package services

import (
	"context"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
	"strings"
)

func GetTags(RoomID *int64) ([]*models.Tag, error) {
	pool, err := utils.NewConnection()
	if err != nil {
		return nil, err
	}
	defer pool.Close()

	query := `SELECT * FROM tags`
	clauses := []string{}
	args := []interface{}{}

	if RoomID != nil {
		clauses = append(clauses, fmt.Sprintf("room_id = $%d", len(clauses)+1))
		args = append(args, RoomID)
	}

	if len(clauses) > 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}

	rowsTags, errQueryTags := pool.Query(context.Background(), query, args...)
	if errQueryTags != nil {
		return nil, err
	}

	// getting tags information
	result := make([]*models.Tag, 0)
	for rowsTags.Next() {
		// init currHotel
		currTag := new(models.Tag)
		currTag.Name = new(string)
		var RoomID int64

		// scanning Tag
		errScan := rowsTags.Scan(&RoomID, currTag.Name)
		if errScan != nil {
			return nil, err
		}
		result = append(result, currTag)
	}
	rowsTags.Close()
	return result, nil
}
