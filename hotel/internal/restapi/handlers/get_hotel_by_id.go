package handlers

import (
	"context"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/utils"
	"log/slog"
	"net/http"
)

func (handler *Handler) GetHotelByIDHandler(params hotel.GetHotelByIDParams) (responder middleware.Responder) {
	// Tracing
	_, span := handler.tracer.Start(context.Background(), "get hotel by id")
	defer span.End()

	defer utils.CatchPanic(&responder)
	hotelByID, err := handler.Database.GetById(params.HotelID)

	if err != nil {
		// Logging
		slog.Error(
			"failed get hotel by id",
			slog.Group("hotel-properties",
				slog.Int64("hotelId", params.HotelID),
			),
			slog.Int("status_code", http.StatusInternalServerError),
			slog.String("error", err.Error()),
		)

		return utils.HandleInternalError(err)
	}
	if hotelByID == nil {
		// Logging
		slog.Info(
			"failed get hotel by id",
			slog.Group("hotel-properties",
				slog.Int64("hotelId", params.HotelID),
			),
			slog.Int("status_code", hotel.GetHotelByIDNotFoundCode),
			slog.String("error", "Hotel with hotelId not found"),
		)

		errCode := int64(hotel.GetHotelByIDNotFoundCode)
		return &hotel.GetHotelByIDNotFound{Payload: &models.Error{
			ErrorStatusCode: &errCode,
			ErrorMessage:    fmt.Sprintf("Hotel with id %d not found", params.HotelID),
		}}
	}
	// Logging
	slog.Info(
		"get hotel by id",
		slog.Group("hotel-properties",
			slog.Int64("hotelId", params.HotelID),
		),
		slog.Int("status_code", hotel.GetHotelByIDOKCode),
	)

	result := new(hotel.GetHotelByIDOK)
	result = result.WithPayload(hotelByID)
	return result
}
