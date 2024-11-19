package handlers

import (
	"github.com/go-openapi/runtime/middleware"
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
		return middleware.Error(http.StatusInternalServerError, err.Error())
	}

	if len(rooms) == 0 {
		return new(room.GetRoomsNotFound)
	}

	result := new(room.GetRoomsOK)
	result = result.WithPayload(rooms)
	return result
}
