package services

import (
	"context"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"strings"
)

func CreateTag(roomId int64, tag *models.Tag) error {
	// connecting to database
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), "db", os.Getenv("POSTGRES_PORT"), "hotel")
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return err
	}
	defer pool.Close()

	query := `INSERT INTO tags`
	fieldNames := []string{}
	fields := []string{}
	values := []interface{}{}

	if tag.Name != nil {
		fieldNames = append(fieldNames, "tag")
		values = append(values, tag.Name)
	}

	fieldNames = append(fieldNames, "room_id")
	values = append(values, roomId)

	for ind := 0; ind < len(fieldNames); ind++ {
		fields = append(fields, fmt.Sprintf("$%d", ind+1))
	}
	query += fmt.Sprintf(" (%s) VALUES (%s)", strings.Join(fieldNames, ", "),
		strings.Join(fields, ", "))

	_, errInsertTag := pool.Exec(context.Background(), query, values...)
	return errInsertTag
}
