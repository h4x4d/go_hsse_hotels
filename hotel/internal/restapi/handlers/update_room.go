package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/room"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/services"
	"net/http"
)

func UpdateRoomHandler(params room.UpdateRoomParams, _ interface{}) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	newRoom := params.Object
	newRoom.ID = params.RoomID
	updatedRoom, errGet := services.GetRoomByID(params.RoomID)
	if errGet != nil {
		return middleware.Error(http.StatusInternalServerError, errGet.Error())
	}
	if updatedRoom != nil {
		// adding tags or deleting otherwise
		if len(newRoom.Tags) > 0 {
			newRoom.Tags = append(newRoom.Tags, updatedRoom.Tags...)
		}

		// deleting old room
		_, errDelete := services.DeleteRoomByID(params.RoomID)
		if errDelete != nil {
			return middleware.Error(http.StatusInternalServerError, errDelete.Error())
		}
	}
	// creating new Room
	createErr := services.CreateRoom(newRoom)
	if createErr != nil {
		return middleware.Error(http.StatusInternalServerError, createErr.Error())
	}
	result := new(room.UpdateRoomOK)
	return result
}
