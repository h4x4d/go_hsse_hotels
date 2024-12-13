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
)

func (handler *Handler) UpdateBooking(params customer.UpdateBookingParams, user *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)
	ctx, span := handler.tracer.Start(context.Background(), "update booking")
	traceId := fmt.Sprintf("%s", span.SpanContext().TraceID())
	ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceId)

	defer span.End()
	owner, err := handler.Database.CheckOwnership(ctx, params.BookingID, user)
	if err != nil {
		return utils.HandleInternalError(err)
	}
	if !owner {
		errCode := int64(customer.GetBookingByIDForbiddenCode)
		result := new(customer.GetBookingByIDForbidden)
		result.SetPayload(&models.Error{
			ErrorMessage:    "You don't have permission to update this booking",
			ErrorStatusCode: &errCode,
		})
		return result
	}
	booking, errUpdate := handler.Database.Update(ctx, params.BookingID, params.Object)
	if errUpdate != nil {
		return utils.HandleInternalError(errUpdate)
	}

	notifyErr := handler.KafkaConn.SendNotification(
		pkg_models.Notification{
			Name: "Booking update",
			Text: fmt.Sprintf("Your booking with booking_id %d was updated successfully",
				params.BookingID),
			TelegramID: user.TelegramID,
		})
	if notifyErr != nil {
		return utils.HandleInternalError(notifyErr)
	}
	hotel, hotelErr := client.GetHotelById(ctx, booking.HotelID)
	if hotelErr != nil {
		return utils.HandleInternalError(hotelErr)
	}
	tgId, tgErr := handler.KeyCloak.GetTelegramId(hotel.UserID)
	if tgErr != nil {
		return utils.HandleInternalError(tgErr)
	}

	notifyErr2 := handler.KafkaConn.SendNotification(
		pkg_models.Notification{
			Name: "Booking update",
			Text: fmt.Sprintf("Your hotel %d booking with booking_id %d was updated",
				*params.Object.HotelID, params.BookingID),
			TelegramID: tgId,
		})
	if notifyErr2 != nil {
		return utils.HandleInternalError(notifyErr2)
	}

	result := new(customer.UpdateBookingOK)
	result.SetPayload(booking)
	return result
}
