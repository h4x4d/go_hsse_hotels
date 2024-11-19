package services

import (
	"context"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"strings"
)

func DeleteTag(roomId int64, tag *models.Tag) (bool, error) {
	// connecting to database hotel
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), "db", os.Getenv("POSTGRES_PORT"), "hotel")
	pool, err := pgxpool.New(context.Background(), connStr)
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
