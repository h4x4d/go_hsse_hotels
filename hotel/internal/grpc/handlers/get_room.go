package handlers

import (
	"context"
	gen "github.com/h4x4d/go_hsse_hotels/hotel/internal/grpc/gen"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type GRPCServer struct {
	HotelServer
	gen.UnimplementedRoomServer
}

type HotelServer interface {
	GetRoom(
		ctx context.Context,
		id int64,
	) (*gen.RoomResponse, error)
}

func Register(gRPCServer *grpc.Server) {
	gen.RegisterRoomServer(gRPCServer, &GRPCServer{})
}

func (serverApi *GRPCServer) GetRoom(
	ctx context.Context, in *gen.RoomRequest) (*gen.RoomResponse, error) {
	log.Println("in func")
	room, err := services.GetRoomByID(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}
	if room == nil {
		return nil, status.Errorf(codes.NotFound, "room %d not found", in.Id)
	}
	return &gen.RoomResponse{
		Id:          room.ID,
		HotelId:     *room.HotelID,
		Cost:        int32(*room.Cost),
		PersonCount: int32(*room.PersonCount),
	}, nil
}
