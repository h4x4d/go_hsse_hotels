package handlers

import "github.com/h4x4d/go_hsse_hotels/hotel/internal/database_service"

type Handler struct {
	Database *database_service.DatabaseService
}

func NewHandler(connStr string) (*Handler, error) {
	db, err := database_service.NewDatabaseService(connStr)
	if err != nil {
		return nil, err
	}
	return &Handler{db}, nil
}
