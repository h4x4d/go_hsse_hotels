package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/utils"
)

func (handler *Handler) CreateHotelHandler(params hotel.CreateHotelParams,
	_ *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	id, err := handler.Database.Create(params.Object)
	if err != nil {
		return utils.HandleInternalError(err)
	}

	result := new(hotel.CreateHotelOK)
	result.SetPayload(&hotel.CreateHotelOKBody{ID: *id})
	return result
}
