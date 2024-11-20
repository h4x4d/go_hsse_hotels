package room

import (
	"context"
	gen "github.com/h4x4d/go_hsse_hotels/hotel/internal/grpc/gen"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCServer struct {
	gen.UnimplementedRoomServer
	hotelServer HotelServer
}

type HotelServer interface {
	GetRoom(
		ctx context.Context,
		id int64,
	) (*models.Room, error)
}

func Register(gRPCServer *grpc.Server, hotel HotelServer) {
	gen.RegisterRoomServer(gRPCServer, &GRPCServer{hotelServer: hotel})
}

func (serverApi *GRPCServer) GetRoom(
	ctx context.Context, in *gen.RoomRequest) (*gen.RoomResponse, error) {
	room, err := services.GetRoomByID(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}
	return &gen.RoomResponse{
		Id:          room.ID,
		HotelId:     *room.HotelID,
		Cost:        int32(*room.Cost),
		PersonCount: int32(*room.PersonCount),
	}, nil
}
