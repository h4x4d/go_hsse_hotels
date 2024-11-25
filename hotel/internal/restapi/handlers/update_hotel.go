package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/utils"
)

func (handler *Handler) UpdateHotelHandler(params hotel.UpdateHotelParams, _ interface{}) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	newHotel := params.Object
	newHotel.ID = params.HotelID
	updatedHotel, errGet := handler.Database.GetById(params.HotelID)
	if errGet != nil {
		return utils.HandleInternalError(errGet)
	}
	if updatedHotel != nil {
		// deleting old hotel
		_, errDelete := handler.Database.DeleteByID(params.HotelID)
		if errDelete != nil {
			return utils.HandleInternalError(errDelete)
		}
	}

	// creating new Hotel
	_, createErr := handler.Database.Create(newHotel)
	if createErr != nil {
		return utils.HandleInternalError(createErr)
	}
	result := new(hotel.UpdateHotelOK)
	result.SetPayload(newHotel)
	return result
}
