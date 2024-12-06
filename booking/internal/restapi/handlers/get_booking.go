package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/restapi/operations/hotelier"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/utils"
)

func (handler *Handler) GetBooking(params hotelier.GetBookingParams, user *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	if user != nil && user.Role == "hotelier" {
		bookings, errGet := handler.Database.GetAll(params.HotelID)
		if errGet != nil {
			return utils.HandleInternalError(errGet)
		}

		result := new(hotelier.GetBookingOK)
		result.SetPayload(bookings)
		return result
	} else {
		errCode := int64(hotelier.GetBookingForbiddenCode)
		result := new(hotelier.GetBookingForbidden)
		result.SetPayload(&models.Error{
			ErrorMessage:    "You doesn't have permission to get a bookings",
			ErrorStatusCode: &errCode,
		})
		return result
	}
}
