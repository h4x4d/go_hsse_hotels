package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/database_service"
	models2 "github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/utils"
	"net/http"
)

func GetHotelsHandler(params hotel.GetHotelsParams) (responder middleware.Responder) {
	// catching panic
	defer utils.CatchPanic(&responder)

	databaseService, contextErr := database_service.GetDatabaseServiceFromContext(params.HTTPRequest.Context())
	if contextErr != nil {
		return middleware.Error(http.StatusInternalServerError, contextErr.Error())
	}

	payload, err := databaseService.GetAll(params.City, params.HotelClass, params.Name)
	if err != nil {
		return utils.HandleInternalError(err)
	}

	if len(payload) == 0 {
		notFound := int64(http.StatusNotFound)
		return new(hotel.GetHotelsNotFound).WithPayload(&models2.Error{
			ErrorMessage:    "Suitable hotels not found",
			ErrorStatusCode: &notFound,
		})
	}

	result := new(hotel.GetHotelsOK)
	result = result.WithPayload(payload)
	return result
}

type GetHotelsHandlerType func(params hotel.GetHotelsParams) middleware.Responder

func (h GetHotelsHandlerType) AddDatabaseService(databaseService *database_service.DatabaseService) GetHotelsHandlerType {
	return func(params hotel.GetHotelsParams) middleware.Responder {
		params.HTTPRequest = params.HTTPRequest.WithContext(database_service.ContextWithDatabaseService(databaseService))
		return h(params)
	}
}
