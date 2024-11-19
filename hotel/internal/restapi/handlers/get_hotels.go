package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/services"
	"net/http"
)

func GetHotelsHandler(params hotel.GetHotelsParams) (responder middleware.Responder) {
	// catching panic
	defer utils.CatchPanic(&responder)

	payload, err := services.GetHotels(params.City, params.HotelClass, params.Name, params.Tag)
	if err != nil {
		return middleware.Error(http.StatusInternalServerError, err.Error())
	}

	result := new(hotel.GetHotelsOK)
	result = result.WithPayload(payload)

	return result
}
