package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/restapi/operations/customer"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/utils"
)

func (handler *Handler) UpdateBooking(params customer.UpdateBookingParams, user *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	if user.Role == "customer" {
		booking, errUpdate := handler.Database.Update(params.BookingID, params.Object)
		if errUpdate != nil {
			return utils.HandleInternalError(errUpdate)
		}

		result := new(customer.UpdateBookingOK)
		result.SetPayload(&customer.UpdateBookingOKBody{BookingID: booking.BookingID})
		return result
	} else if user.Role == "hotelier" {
		errCode := int64(customer.UpdateBookingForbiddenCode)
		result := new(customer.UpdateBookingForbidden)
		result.SetPayload(&models.Error{
			ErrorMessage:    "You doesn't have permission to create a booking",
			ErrorStatusCode: &errCode,
		})
		return result
	} else {
		errCode := int64(customer.UpdateBookingBadRequestCode)
		result := new(customer.UpdateBookingBadRequest)
		result.SetPayload(&models.Error{
			ErrorMessage:    "Your role does not exist",
			ErrorStatusCode: &errCode,
		})
		return result
	}
}
