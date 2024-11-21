package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/services"
)

func UpdateHotelHandler(params hotel.UpdateHotelParams, _ interface{}) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	newHotel := params.Object
	newHotel.ID = params.HotelID
	updatedHotel, errGet := services.GetHotelByID(params.HotelID)
	if errGet != nil {
		return utils.HandleInternalError(errGet)
	}
	if updatedHotel != nil {
		// adding new rooms or deleting them otherwise
		if len(newHotel.Rooms) == 0 {
			newHotel.Rooms = append(newHotel.Rooms, updatedHotel.Rooms...)
		}

		// deleting old hotel
		_, errDelete := services.DeleteHotelByID(params.HotelID)
		if errDelete != nil {
			return utils.HandleInternalError(errDelete)
		}
	}
	// creating new Hotel
	id, createErr := services.CreateHotel(newHotel)
	if createErr != nil {
		return utils.HandleInternalError(createErr)
	}
	result := new(hotel.UpdateHotelOK)
	result.SetPayload(&hotel.UpdateHotelOKBody{ID: *id})
	return result
}
