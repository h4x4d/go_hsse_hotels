package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	models2 "github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/room"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/services"
	"net/http"
)

func GetRoomsHandler(params room.GetRoomsParams) (responder middleware.Responder) {
	// catching panic
	defer utils.CatchPanic(&responder)

	rooms, err := services.GetRooms(params.HotelID, params.Tag)
	if err != nil {
		return utils.HandleInternalError(err)
	}

	if len(rooms) == 0 {
		notFound := int64(http.StatusNotFound)
		return new(room.GetRoomsNotFound).WithPayload(&models2.Error{
			ErrorMessage:    "Suitable rooms not found",
			ErrorStatusCode: &notFound,
		})
	}

	result := new(room.GetRoomsOK)
	result = result.WithPayload(rooms)
	return result
}
