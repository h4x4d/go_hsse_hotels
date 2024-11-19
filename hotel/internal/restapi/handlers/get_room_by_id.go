package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/room"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/services"
	"net/http"
)

func GetRoomByIDHandler(params room.GetRoomByIDParams) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	roomByID, err := services.GetRoomByID(params.RoomID)

	if err != nil {
		return middleware.Error(http.StatusInternalServerError, err.Error())
	}
	if roomByID == nil {
		return new(room.GetRoomByIDNotFound)
	}

	result := new(room.GetRoomByIDOK)
	result = result.WithPayload(roomByID)
	return result
}
