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

func (handler *Handler) UpdateHotelHandler(params hotel.UpdateHotelParams, user *models.User) (responder middleware.Responder) {
	// Tracing
	_, span := handler.tracer.Start(context.Background(), "update hotel")
	defer span.End()

	defer utils.CatchPanic(&responder)

	existing, errGet := handler.Database.GetById(params.HotelID)
	if errGet != nil {
		// Logging
		slog.Error(
			"failed update hotel",
			slog.String("method", "PUT"),
			slog.Group("user-properties",
				slog.String("user-id", user.UserID),
				slog.String("role", user.Role),
				slog.Int("telegram-id", user.TelegramID),
			),
			slog.Int("status_code", http.StatusInternalServerError),
			slog.String("error", errGet.Error()),
		)

		return utils.HandleInternalError(errGet)
	}
	if existing == nil {
		// Logging
		slog.Error(
			"failed update hotel",
			slog.String("method", "PUT"),
			slog.Group("user-properties",
				slog.String("user-id", user.UserID),
				slog.String("role", user.Role),
				slog.Int("telegram-id", user.TelegramID),
			),
			slog.Int("status_code", hotel.UpdateHotelNotFoundCode),
			slog.String("error", fmt.Sprintf("Hotel with id %d not found", params.HotelID)),
		)

		code := int64(hotel.UpdateHotelNotFoundCode)
		return &hotel.UpdateHotelNotFound{Payload: &models.Error{
			ErrorMessage:    fmt.Sprintf("Hotel with id %d not found", params.HotelID),
			ErrorStatusCode: &code,
		}}
	}
	if existing.UserID != user.UserID {
		// Logging
		slog.Error(
			"failed update hotel",
			slog.String("method", "PUT"),
			slog.Group("user-properties",
				slog.String("user-id", user.UserID),
				slog.String("role", user.Role),
				slog.Int("telegram-id", user.TelegramID),
			),
			slog.Group("user-owner-properties",
				slog.String("user-id", existing.UserID),
			),
			slog.Int("status_code", hotel.UpdateHotelForbiddenCode),
			slog.String("error", "Not enough rights"),
		)

		code := int64(hotel.UpdateHotelForbiddenCode)
		return &hotel.UpdateHotelForbidden{Payload: &models.Error{
			ErrorMessage:    "You can't edit hotel that does not belong to you",
			ErrorStatusCode: &code,
		}}
	}

	newHotel := params.Object
	updated, errUpdate := handler.Database.Update(params.HotelID, newHotel)
	if errUpdate != nil {
		// Logging
		slog.Error(
			"failed update hotel",
			slog.String("method", "Put"),
			slog.Group("user-properties",
				slog.String("user-id", user.UserID),
				slog.String("role", user.Role),
				slog.Int("telegram-id", user.TelegramID),
			),
			slog.Group("hotel-properties",
				slog.String("name", *params.Object.Name),
				slog.Int64("cost", params.Object.Cost),
				slog.Int64("hotel-id", params.Object.ID),
				slog.Int64("hotel-class", params.Object.HotelClass),
				slog.String("address", *params.Object.Address),
				slog.String("city", *params.Object.City),
			),
			slog.Int("status_code", http.StatusInternalServerError),
			slog.String("error", errUpdate.Error()),
		)

		return utils.HandleInternalError(errUpdate)
	}

	// Logging
	slog.Info(
		"update hotel",
		slog.String("method", "PUT"),
		slog.Group("user-properties",
			slog.String("user-id", user.UserID),
			slog.String("role", user.Role),
			slog.Int("telegram-id", user.TelegramID),
		),
		slog.Group("old-hotel-properties",
			slog.String("name", *existing.Name),
			slog.Int64("hotel-id", existing.ID),
			slog.Int64("cost", existing.Cost),
			slog.Int64("hotel-class", existing.HotelClass),
			slog.String("address", *existing.Address),
			slog.String("city", *existing.City),
		),
		slog.Group("new-hotel-properties",
			slog.String("name", *params.Object.Name),
			slog.Int64("hotel-id", params.Object.ID),
			slog.Int64("cost", params.Object.Cost),
			slog.Int64("hotel-class", params.Object.HotelClass),
			slog.String("address", *params.Object.Address),
			slog.String("city", *params.Object.City),
		),
		slog.Int("status_code", hotel.UpdateHotelOKCode),
	)

	result := new(hotel.UpdateHotelOK)
	result.SetPayload(updated)
	return result
}
