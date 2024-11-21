package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/room"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/services"
)

func CreateRoomHandler(params room.CreateRoomParams, _ interface{}) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	err := services.CreateRoom(params.Object)

	if err != nil {
		return utils.HandleInternalError(err)
	}
	result := new(room.CreateRoomOK)
	return result
}
