package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/room"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/services"
	"net/http"
)

func DeleteRoomByIDHandler(params room.DeleteRoomByIDParams, _ interface{}) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	deletedRoomId, errDelete := services.DeleteRoomByID(params.RoomID)
	if errDelete != nil {
		return middleware.Error(http.StatusInternalServerError, errDelete.Error())
	}
	if deletedRoomId == nil {
		return new(room.DeleteRoomByIDNotFound)
	}
	result := new(room.DeleteRoomByIDOK)
	return result
}
