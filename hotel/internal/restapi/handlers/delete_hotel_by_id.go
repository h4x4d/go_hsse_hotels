package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/services"
	"net/http"
)

func DeleteHotelByIDHandler(params hotel.DeleteHotelByIDParams, _ interface{}) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	deletedHotel, errGet := services.GetHotelByID(params.HotelID)
	if errGet != nil {
		return middleware.Error(http.StatusInternalServerError, errGet.Error())
	}
	if deletedHotel == nil {
		return new(hotel.DeleteHotelByIDNotFound)
	}

	errDelete := services.DeleteHotelByID(params.HotelID)

	if errDelete != nil {
		return middleware.Error(http.StatusInternalServerError, errDelete.Error())
	}
	result := new(hotel.DeleteHotelByIDOK)
	return result
}
