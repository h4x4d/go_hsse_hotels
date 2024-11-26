package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/database_service"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/utils"
	"net/http"
)

func UpdateHotelHandler(params hotel.UpdateHotelParams, _ interface{}) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	databaseService, contextErr := database_service.GetDatabaseServiceFromContext(params.HTTPRequest.Context())
	if contextErr != nil {
		return middleware.Error(http.StatusInternalServerError, contextErr.Error())
	}

	newHotel := params.Object
	newHotel.ID = params.HotelID
	updatedHotel, errGet := databaseService.GetById(params.HotelID)
	if errGet != nil {
		return utils.HandleInternalError(errGet)
	}
	if updatedHotel != nil {
		// deleting old hotel
		_, errDelete := databaseService.DeleteByID(params.HotelID)
		if errDelete != nil {
			return utils.HandleInternalError(errDelete)
		}
	}

	// creating new Hotel
	_, createErr := databaseService.Create(newHotel)
	if createErr != nil {
		return utils.HandleInternalError(createErr)
	}
	result := new(hotel.UpdateHotelOK)
	result.SetPayload(newHotel)
	return result
}

type UpdateHotelHandlerType func(params hotel.UpdateHotelParams, _ interface{}) middleware.Responder

func (h UpdateHotelHandlerType) AddDatabaseService(databaseService *database_service.DatabaseService) UpdateHotelHandlerType {
	return func(params hotel.UpdateHotelParams, principal interface{}) middleware.Responder {
		params.HTTPRequest = params.HTTPRequest.WithContext(database_service.ContextWithDatabaseService(databaseService))
		return h(params, principal)
	}
}
