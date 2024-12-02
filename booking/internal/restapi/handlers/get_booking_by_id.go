package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/restapi/operations/customer"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/utils"
	"net/http"
)

func (handler *Handler) GetBookingByID(params customer.GetBookingByIDParams, user *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	if user.Role == "customer" {
		booking, errGet := handler.Database.GetByID(params.BookingID)
		if errGet != nil {
			return utils.HandleInternalError(errGet)
		}

		result := new(customer.GetBookingByIDOK)
		result.SetPayload(booking)
		return result
	} else if user.Role == "hotelier" {
		errCode := int64(customer.GetBookingByIDForbiddenCode)
		result := new(customer.GetBookingByIDForbidden)
		result.SetPayload(&models.Error{
			ErrorMessage:    "You doesn't have permission to create a booking",
			ErrorStatusCode: &errCode,
		})
		return result
	} else {
		// here must be bad request
		errCode := int64(http.StatusBadRequest)
		result := new(customer.GetBookingByIDForbidden)
		result.SetPayload(&models.Error{
			ErrorMessage:    "Your role does not exist",
			ErrorStatusCode: &errCode,
		})
		return result
	}
}
