package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/database_service"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/utils"
	"net/http"
)

func CreateHotelHandler(params hotel.CreateHotelParams, _ interface{}) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	databaseService, contextErr := database_service.GetDatabaseServiceFromContext(params.HTTPRequest.Context())
	if contextErr != nil {
		return middleware.Error(http.StatusInternalServerError, contextErr.Error())
	}

	id, err := databaseService.Create(params.Object)

	if err != nil {
		return utils.HandleInternalError(err)
	}
	result := new(hotel.CreateHotelOK)
	result.SetPayload(&hotel.CreateHotelOKBody{ID: *id})
	return result
}

type CreateHotelHandlerType func(params hotel.CreateHotelParams, _ interface{}) middleware.Responder

func (h CreateHotelHandlerType) AddDatabaseService(databaseService *database_service.DatabaseService) CreateHotelHandlerType {
	return func(params hotel.CreateHotelParams, principal interface{}) middleware.Responder {
		params.HTTPRequest = params.HTTPRequest.WithContext(database_service.ContextWithDatabaseService(databaseService))
		return h(params, principal)
	}
}
