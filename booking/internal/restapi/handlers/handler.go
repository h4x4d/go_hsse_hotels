package handlers

import (
	"github.com/h4x4d/go_hsse_hotels/booking/internal/database_service"
	"github.com/h4x4d/go_hsse_hotels/pkg/jaeger"
	"go.opentelemetry.io/otel/trace"
	"log"
)

type Handler struct {
	Database *database_service.DatabaseService
	tracer   trace.Tracer
}

func NewHandler(connStr string) (*Handler, error) {
	db, err := database_service.NewDatabaseService(connStr)
	if err != nil {
		return nil, err
	}

	tracer, err := jaeger.InitTracer("Booking")
	if err != nil {
		log.Fatal("init tracer", err)
	}
	return &Handler{db, tracer}, nil
}

func (handler *Handler) GetTracer() trace.Tracer {
	return handler.tracer
}
