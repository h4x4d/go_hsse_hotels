package handlers

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/grpc/client"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/restapi/operations/hotelier"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func (handler *Handler) GetBooking(params hotelier.GetBookingParams, user *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	if user != nil && user.Role == "hotelier" {
		hotel, hotelErr := client.GetHotelById(params.HotelID)
		if hotelErr != nil {
			if statusCode, ok := status.FromError(hotelErr); ok && statusCode.Code() == codes.NotFound {
				code := int64(http.StatusNotFound)
				return &hotelier.GetBookingNotFound{
					Payload: &models.Error{
						ErrorStatusCode: &code,
						ErrorMessage:    fmt.Sprintf("Hotel with id %d not found", params.HotelID),
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

			result := new(hotelier.GetBookingOK)
			result.SetPayload(bookings)
			return result
		}
	}
	errCode := int64(hotelier.GetBookingForbiddenCode)
	result := new(hotelier.GetBookingForbidden)
	result.SetPayload(&models.Error{
		ErrorMessage:    "You don't have permission to get this bookings",
		ErrorStatusCode: &errCode,
	})
	return result
}
