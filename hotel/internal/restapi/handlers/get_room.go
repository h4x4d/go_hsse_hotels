package handlers

import (
	"errors"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/room"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/services"
	errors2 "github.com/h4x4d/go_hsse_hotels/hotel/internal/services/errors"
	"net/http"
)

func GetRoomHandler(params room.GetRoomByIDParams) (responder middleware.Responder) {
	// catching panic
	defer utils.CatchPanic(&responder)

	payload, err := services.GetRoom(int(params.RoomID))
	notFound := new(errors2.RoomNotFound)
	if errors.As(err, &notFound) {
		notFound := int64(http.StatusNotFound)
		return room.NewGetRoomByIDNotFound().WithPayload(
			&models.Error{ErrorMessage: fmt.Sprintf("Element %d Not Found", params.RoomID),
				ErrorStatusCode: &notFound})
	}
	if err != nil {
		return middleware.Error(http.StatusInternalServerError, err.Error())
	}

	result := new(room.GetRoomByIDOK)
	result.SetPayload(payload)

	return result
}
