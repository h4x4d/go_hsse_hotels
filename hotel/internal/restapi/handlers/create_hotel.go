package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/utils"
)

func (handler *Handler) CreateHotelHandler(params hotel.CreateHotelParams,
	user *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	if user == nil || user.Role != "hotelier" {
		code := int64(hotel.CreateHotelForbiddenCode)
		result := hotel.CreateHotelForbidden{Payload: &models.Error{
			"Creation of hotels allowed only to hoteliers",
			&code,
		}}
		return &result
	}

	id, err := handler.Database.Create(params.Object, user)
	if err != nil {
		return utils.HandleInternalError(err)
	}

	result := new(hotel.CreateHotelOK)
	result.SetPayload(&hotel.CreateHotelOKBody{ID: *id})
	return result
}
