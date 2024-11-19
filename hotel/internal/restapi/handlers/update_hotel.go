package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/services"
	"net/http"
)

func UpdateHotelHandler(params hotel.UpdateHotelParams, _ interface{}) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	newHotel := params.Object
	updatedHotel, errGet := services.GetHotelByID(params.HotelID)
	if errGet != nil {
		return middleware.Error(http.StatusInternalServerError, errGet.Error())
	}
	if updatedHotel != nil {
		if len(newHotel.Rooms) == 0 {
			newHotel.Rooms = updatedHotel.Rooms
		}
		// deleting old hotel
		errDelete := services.DeleteHotelByID(params.HotelID)
		if errDelete != nil {
			return middleware.Error(http.StatusInternalServerError, errDelete.Error())
		}
	}
	// creating new Hotel
	createErr := services.CreateHotel(newHotel)
	if createErr != nil {
		return middleware.Error(http.StatusInternalServerError, createErr.Error())
	}
	result := new(hotel.UpdateHotelOK)
	return result
}
