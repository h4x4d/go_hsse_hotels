package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/restapi/operations/hotelier"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/utils"
	"net/http"
)

func (handler *Handler) GetBooking(params hotelier.GetBookingParams, user *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	if user.Role == "hotelier" {
		bookings, errGet := handler.Database.GetAll(params.HotelID)
		if errGet != nil {
			return utils.HandleInternalError(errGet)
		}

		result := new(hotelier.GetBookingOK)
		result.SetPayload(bookings)
		return result
	} else if user.Role == "customer" {
		errCode := int64(hotelier.GetBookingForbiddenCode)
		result := new(hotelier.GetBookingForbidden)
		result.SetPayload(&models.Error{
			ErrorMessage:    "You doesn't have permission to create a booking",
			ErrorStatusCode: &errCode,
		})
		return result
	} else {
		// here must be bad request
		errCode := int64(http.StatusBadRequest)
		result := new(hotelier.GetBookingForbidden)
		result.SetPayload(&models.Error{
			ErrorMessage:    "Your role does not exist",
			ErrorStatusCode: &errCode,
		})
		return result
	}
}
