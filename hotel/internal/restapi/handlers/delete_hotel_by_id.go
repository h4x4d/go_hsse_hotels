package handlers

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	models2 "github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/services"
	"net/http"
)

func DeleteHotelByIDHandler(params hotel.DeleteHotelByIDParams, _ interface{}) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	deletedHotelId, errDelete := services.DeleteHotelByID(params.HotelID)
	if errDelete != nil {
		return utils.HandleInternalError(errDelete)
	}
	if deletedHotelId == nil {
		notFound := int64(http.StatusNotFound)
		return new(hotel.DeleteHotelByIDNotFound).WithPayload(&models2.Error{
			ErrorMessage:    fmt.Sprintf("Hotel with id %d not found", params.HotelID),
			ErrorStatusCode: &notFound,
		})
	}
	result := new(hotel.DeleteHotelByIDOK)
	result.SetPayload(&models2.Result{Status: "success"})
	return result
}
