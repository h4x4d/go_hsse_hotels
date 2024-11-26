package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/database_service"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/utils"
	"net/http"
)

func GetHotelByIDHandler(params hotel.GetHotelByIDParams) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	databaseService, contextErr := database_service.GetDatabaseServiceFromContext(params.HTTPRequest.Context())
	if contextErr != nil {
		return middleware.Error(http.StatusInternalServerError, contextErr.Error())
	}

	hotelByID, err := databaseService.GetById(params.HotelID)

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

type GetHotelByIDHandlerType func(params hotel.GetHotelByIDParams) middleware.Responder

func (h GetHotelByIDHandlerType) AddDatabaseService(databaseService *database_service.DatabaseService) GetHotelByIDHandlerType {
	return func(params hotel.GetHotelByIDParams) middleware.Responder {
		params.HTTPRequest = params.HTTPRequest.WithContext(database_service.ContextWithDatabaseService(databaseService))
		return h(params)
	}
}
