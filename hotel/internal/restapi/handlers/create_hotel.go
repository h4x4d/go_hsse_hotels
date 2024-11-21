package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/services"
)

func CreateHotelHandler(params hotel.CreateHotelParams, _ interface{}) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	err := services.CreateHotel(params.Object)

	if err != nil {
		return utils.HandleInternalError(err)
	}
	result := new(hotel.CreateHotelOK)
	return result
}
