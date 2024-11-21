package handlers

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	models2 "github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/room"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/services"
	"net/http"
)

func DeleteRoomByIDHandler(params room.DeleteRoomByIDParams, _ interface{}) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	deletedRoomId, errDelete := services.DeleteRoomByID(params.RoomID)
	if errDelete != nil {
		return utils.HandleInternalError(errDelete)
	}
	if deletedRoomId == nil {
		notFound := int64(http.StatusNotFound)
		return new(room.DeleteRoomByIDNotFound).WithPayload(&models2.Error{
			ErrorMessage:    fmt.Sprintf("Room with id %d not found", params.RoomID),
			ErrorStatusCode: &notFound,
		})
	}
	result := new(room.DeleteRoomByIDOK)
	return result
}
