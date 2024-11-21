package services

import (
	"context"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
	"strings"
)

func DeleteTag(roomId int64, tag *models.Tag) (bool, error) {
	pool, err := utils.NewConnection()
	if err != nil {
		return false, err
	}
	defer pool.Close()

	query := `DELETE FROM tags`
	clauses := []string{}
	args := []interface{}{}

	if tag.Name != nil {
		clauses = append(clauses, fmt.Sprintf("name = $%d", len(clauses)+1))
		args = append(args, tag.Name)
	}
	clauses = append(clauses, fmt.Sprintf("room_id = $%d", len(clauses)+1))
	args = append(args, roomId)

	query += " WHERE " + strings.Join(clauses, " AND ")
	query += " RETURNING id"

	deletedQuery, errDeleteTag := pool.Query(context.Background(), query, args...)
	if errDeleteTag != nil {
		return false, errDeleteTag
	}
	// more than one tag can be deleted TODO
	if !deletedQuery.Next() {
		return false, nil
	}
	return true, nil
}
