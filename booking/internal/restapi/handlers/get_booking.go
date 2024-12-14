package handlers

import (
	"context"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/grpc/client"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/restapi/operations/hotelier"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log/slog"
	"net/http"
)

func (handler *Handler) GetBooking(params hotelier.GetBookingParams, user *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	// Tracing
	ctx, span := handler.tracer.Start(context.Background(), "get booking")
	defer span.End()
	traceId := fmt.Sprintf("%s", span.SpanContext().TraceID())
	ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceId)

	if user != nil && user.Role == "hotelier" {
		hotel, hotelErr := client.GetHotelById(ctx, params.HotelID)
		if hotelErr != nil {
			if statusCode, ok := status.FromError(hotelErr); ok && statusCode.Code() == codes.NotFound {
				// Logging
				slog.Error(
					"failed get bookings",
					slog.String("method", "GET"),
					slog.Group("user-properties",
						slog.String("user-id", user.UserID),
						slog.String("role", user.Role),
						slog.Int("telegram-id", user.TelegramID),
					),
					slog.Group("booking-properties",
						slog.Int64("hotel-id", *params.HotelID),
					),
					slog.Int("status_code", http.StatusNotFound),
					slog.String("error", "Not found"),
				)

				code := int64(http.StatusNotFound)
				return &hotelier.GetBookingNotFound{
					Payload: &models.Error{
						ErrorStatusCode: &code,
						ErrorMessage:    fmt.Sprintf("Hotel with id %d not found", *params.HotelID),
					},
				}
			}
			return utils.HandleInternalError(hotelErr)
		}
		if hotel.UserID == user.UserID {
			bookings, errGet := handler.Database.GetAll(params.HotelID)
			if errGet != nil {
				return utils.HandleInternalError(errGet)
			}

			// Logging
			slog.Info(
				"get bookings",
				slog.String("method", "GET"),
				slog.Group("user-properties",
					slog.String("user-id", user.UserID),
					slog.String("role", user.Role),
					slog.Int("telegram-id", user.TelegramID),
				),
				slog.Group("booking-properties",
					slog.Int64("hotel-id", *params.HotelID),
				),
				slog.Int("status_code", hotelier.GetBookingOKCode),
			)

			result := new(hotelier.GetBookingOK)
			result.SetPayload(bookings)
			return result
		}
	}
	if user == nil {
		user = &models.User{
			UserID:     "empty",
			Role:       "empty",
			TelegramID: 0,
		}
	}
	// Logging
	slog.Error(
		"failed get bookings",
		slog.String("method", "GET"),
		slog.Group("user-properties",
			slog.String("user-id", user.UserID),
			slog.String("role", user.Role),
			slog.Int("telegram-id", user.TelegramID),
		),
		slog.Group("booking-properties",
			slog.Int64("hotel-id", *params.HotelID),
		),
		slog.Int("status_code", hotelier.GetBookingForbiddenCode),
		slog.String("error", "Not enough rights"),
	)

	errCode := int64(hotelier.GetBookingForbiddenCode)
	result := new(hotelier.GetBookingForbidden)
	result.SetPayload(&models.Error{
		ErrorMessage:    "You don't have permission to get this bookings",
		ErrorStatusCode: &errCode,
	})
	return result
}
