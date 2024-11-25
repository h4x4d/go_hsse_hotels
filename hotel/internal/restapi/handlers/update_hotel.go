package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/utils"
)

func (handler *Handler) UpdateHotelHandler(params hotel.UpdateHotelParams, _ interface{}) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	newHotel := params.Object
	updated, errUpdate := handler.Database.Update(params.HotelID, newHotel)
	if errUpdate != nil {
		return utils.HandleInternalError(errUpdate)
	}

	result := new(hotel.UpdateHotelOK)
	result.SetPayload(updated)
	return result
}
