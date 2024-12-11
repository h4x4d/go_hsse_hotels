package handlers

import (
	"github.com/h4x4d/go_hsse_hotels/booking/internal/database_service"
	"github.com/h4x4d/go_hsse_hotels/pkg/client"
	"github.com/h4x4d/go_hsse_hotels/pkg/notification"
	"go.opentelemetry.io/otel/trace"
	"log"
)

type Handler struct {
	Database  *database_service.DatabaseService
	KafkaConn *notification.KafkaConnection
	KeyCloak  *client.Client
	tracer   trace.Tracer
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
	tracer, err := jaeger.InitTracer("Booking")
	if err != nil {
		log.Fatal("init tracer", err)
	}
	return &Handler{db, conn, nil, tracer}, nil
}

func (handler *Handler) GetTracer() trace.Tracer {
	return handler.tracer
}
