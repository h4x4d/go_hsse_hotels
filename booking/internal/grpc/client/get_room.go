package client

import (
	"context"
	gen "github.com/h4x4d/go_hsse_hotels/booking/internal/grpc/gen"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/grpc/utils"
	"log"
)

func GetRoomById(roomId int) {
	conn, err := utils.ConnectToHotel()
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := gen.NewRoomClient(conn)
	bookList, err := client.GetRoom(context.Background(), &gen.RoomRequest{Id: int64(roomId)})
	if err != nil {
		log.Fatalf("failed to get book list: %v", err)
	}
	log.Printf("room: %v", bookList)
}
