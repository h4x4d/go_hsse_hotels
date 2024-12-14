package handlers

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/utils"
	"log/slog"
)

func (handler *Handler) CreateHotelHandler(params hotel.CreateHotelParams,
	user *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)
	// Tracing
	_, span := handler.tracer.Start(context.Background(), "create hotel")
	defer span.End()

	if user == nil || user.Role != "hotelier" {
		if user == nil {
			user = &models.User{
				UserID:     "empty",
				Role:       "empty",
				TelegramID: 0,
			}
		}
		// Logging
		slog.Error(
			"failed create new hotel",
			slog.String("method", "POST"),
			slog.Group("user-properties",
				slog.String("user-id", user.UserID),
				slog.String("role", user.Role),
				slog.Int("telegram-id", user.TelegramID),
			),
			slog.Group("hotel-properties",
				slog.String("name", *params.Object.Name),
				slog.Int64("cost", params.Object.Cost),
				slog.Int64("hotel-class", params.Object.HotelClass),
				slog.String("address", *params.Object.Address),
				slog.String("city", *params.Object.City),
			),
			slog.Int("status_code", hotel.CreateHotelForbiddenCode),
			slog.String("error", "Creation of hotels allowed only to hoteliers"),
		)

		code := int64(hotel.CreateHotelForbiddenCode)
		result := hotel.CreateHotelForbidden{Payload: &models.Error{
			"Creation of hotels allowed only to hoteliers",
			&code,
		}}
		return &result
	}

	id, err := handler.Database.Create(params.Object, user)
	if err != nil {
		// Logging
		slog.Error(
			"failed create new hotel",
			slog.String("method", "POST"),
			slog.Group("user-properties",
				slog.String("user-id", user.UserID),
				slog.String("role", user.Role),
				slog.Int("telegram-id", user.TelegramID),
			),
			slog.Group("hotel-properties",
				slog.String("name", *params.Object.Name),
				slog.Int64("cost", params.Object.Cost),
				slog.Int64("hotel-class", params.Object.HotelClass),
				slog.String("address", *params.Object.Address),
				slog.String("city", *params.Object.City),
				slog.Int64("id", params.Object.ID),
			),
			slog.Int("status_code", hotel.CreateHotelForbiddenCode),
			slog.String("error", err.Error()),
		)

		return utils.HandleInternalError(err)
	}

	// Logging
	slog.Info(
		"created new hotel",
		slog.String("method", "POST"),
		slog.Group("user-properties",
			slog.String("userId", user.UserID),
			slog.String("role", user.Role),
			slog.Int("telegram-id", user.TelegramID),
		),
		slog.Group("hotel-properties",
			slog.String("name", *params.Object.Name),
			slog.Int64("cost", params.Object.Cost),
			slog.Int64("hotel-class", params.Object.HotelClass),
			slog.String("address", *params.Object.Address),
			slog.String("city", *params.Object.City),
			slog.Int64("id", *id),
		),
		slog.Int("status_code", hotel.CreateHotelOKCode),
	)

	result := new(hotel.CreateHotelOK)
	result.SetPayload(&hotel.CreateHotelOKBody{ID: *id})
	return result
}
