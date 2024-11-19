package errors

import "fmt"

type RoomNotFound struct {
	RoomID int
}

func (e *RoomNotFound) Error() string {
	return fmt.Sprintf("room with ID %d not found", e.RoomID)
}
