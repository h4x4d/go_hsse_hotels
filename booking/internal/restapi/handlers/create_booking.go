package handlers

import (
	"context"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/grpc/client"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/restapi/operations/customer"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/utils"
	pkg_models "github.com/h4x4d/go_hsse_hotels/pkg/models"
	"google.golang.org/grpc/metadata"
	"net/http"
)

func (handler *Handler) CreateBooking(params customer.CreateBookingParams, user *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	ctx, span := handler.tracer.Start(context.Background(), "create booking")
	defer span.End()
	traceId := fmt.Sprintf("%s", span.SpanContext().TraceID())
	ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceId)

	if user != nil && user.Role == "customer" {
		bookingId, errCreate := handler.Database.Create(ctx,
			params.Object.DateFrom,
			params.Object.DateTo,
			params.Object.HotelID,
			user.UserID,
		)
		if errCreate != nil {
			return utils.HandleInternalError(errCreate)
		}

		notifyErr := handler.KafkaConn.SendNotification(
			pkg_models.Notification{
				Name: "New booking",
				Text: fmt.Sprintf("Your booking with booking_id %d was created successfully",
					*bookingId),
				TelegramID: user.TelegramID,
			})
		if notifyErr != nil {
			return utils.HandleInternalError(notifyErr)
		}
		hotel, hotelErr := client.GetHotelById(ctx, params.Object.HotelID)
		if hotelErr != nil {
			return utils.HandleInternalError(hotelErr)
		}
		tgId, tgErr := handler.KeyCloak.GetTelegramId(hotel.UserID)
		if tgErr != nil {
			return utils.HandleInternalError(tgErr)
		}

		notifyErr2 := handler.KafkaConn.SendNotification(
			pkg_models.Notification{
				Name: "New Booking",
				Text: fmt.Sprintf("Your hotel %d was booked with booking_id %d",
					*params.Object.HotelID, *bookingId),
				TelegramID: tgId,
			})
		if notifyErr2 != nil {
			return utils.HandleInternalError(notifyErr2)
		}

		result := new(customer.CreateBookingOK)
		result.SetPayload(&customer.CreateBookingOKBody{BookingID: *bookingId})
		return result
	} else {
		errCode := int64(http.StatusForbidden)
		result := new(customer.CreateBookingForbidden)
		result.SetPayload(&models.Error{
			ErrorMessage:    "You don't have permission to create a booking",
			ErrorStatusCode: &errCode,
		})
		return result
	}
}
