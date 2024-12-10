package handlers

import (
	"github.com/h4x4d/go_hsse_hotels/booking/internal/database_service"
	"github.com/h4x4d/go_hsse_hotels/pkg/client"
	"github.com/h4x4d/go_hsse_hotels/pkg/notification"
)

type Handler struct {
	Database  *database_service.DatabaseService
	KafkaConn *notification.KafkaConnection
	KeyCloak  *client.Client
}

func NewHandler(connStr string) (*Handler, error) {
	db, err := database_service.NewDatabaseService(connStr)
	if err != nil {
		return nil, err
	}
	conn, kafkaErr := notification.NewEnvKafkaConnection()
	if kafkaErr != nil {
		return nil, kafkaErr
	}
	return &Handler{db, conn, nil}, nil
}
