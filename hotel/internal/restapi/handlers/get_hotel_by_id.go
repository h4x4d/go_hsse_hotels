package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/utils"
)

func (handler *Handler) GetHotelByIDHandler(params hotel.GetHotelByIDParams) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)
	hotelByID, err := handler.Database.GetById(params.HotelID)

	if err != nil {
		return utils.HandleInternalError(err)
	}
	if hotelByID == nil {
		return new(hotel.GetHotelByIDNotFound)
	}

	result := new(hotel.GetHotelByIDOK)
	result = result.WithPayload(hotelByID)
	return result
}
