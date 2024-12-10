package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/restapi/operations/customer"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/utils"
	"net/http"
)

func (handler *Handler) CreateBooking(params customer.CreateBookingParams, user *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	if user != nil && user.Role == "customer" {
		bookingId, errCreate := handler.Database.Create(
			params.Object.DateFrom,
			params.Object.DateTo,
			params.Object.HotelID,
			user.UserID,
		)
		if errCreate != nil {
			return utils.HandleInternalError(errCreate)
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
