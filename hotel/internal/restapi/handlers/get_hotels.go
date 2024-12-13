package handlers

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	models2 "github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/utils"
	"net/http"
)

func (handler *Handler) GetHotelsHandler(params hotel.GetHotelsParams) (responder middleware.Responder) {
	_, span := handler.tracer.Start(context.Background(), "get hotels")
	defer span.End()
	defer utils.CatchPanic(&responder)

	payload, err := handler.Database.GetAll(params.City, params.HotelClass, params.Name)
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
