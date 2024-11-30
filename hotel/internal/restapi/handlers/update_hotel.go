package handlers

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/utils"
)

func (handler *Handler) UpdateHotelHandler(params hotel.UpdateHotelParams, _ *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	exists, errGet := handler.Database.Exists(params.HotelID)
	if errGet != nil {
		return utils.HandleInternalError(errGet)
	}
	if exists == false {
		code := int64(hotel.UpdateHotelNotFoundCode)
		return &hotel.UpdateHotelNotFound{Payload: &models.Error{
			ErrorMessage:    fmt.Sprintf("no hotel with id %d", params.HotelID),
			ErrorStatusCode: &code,
		}}
	}

	newHotel := params.Object
	updated, errUpdate := handler.Database.Update(params.HotelID, newHotel)
	if errUpdate != nil {
		return utils.HandleInternalError(errUpdate)
	}

	result := new(hotel.UpdateHotelOK)
	result.SetPayload(updated)
	return result
}
