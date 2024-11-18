package room

import (
	"context"
	grpc_booking "github.com/h4x4d/go_hsse_hotels/booking/grpc_gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCServer struct {
	grpc_booking.UnimplementedBookingServer
	booking Booking
}

type Room struct {
	Id          int64 `json:"id"`
	HotelId     int64 `json:"hotel_id"`
	Cost        int32 `json:"cost"`
	PersonCount int32 `json:"person_count"`
}

type Booking interface {
	GetRoom(
		ctx context.Context,
		id int64,
	) (Room, error)
}

func Register(gRPCServer *grpc.Server, booking Booking) {
	grpc_booking.RegisterBookingServer(gRPCServer, &GRPCServer{booking: booking})
}

func (serverApi *GRPCServer) GetRoom(
	ctx context.Context, in *grpc_booking.RoomRequest) (*grpc_booking.RoomResponse, error) {
	room, err := serverApi.booking.GetRoom(ctx, in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}
	return &grpc_booking.RoomResponse{
		Id:          room.Id,
		HotelId:     room.HotelId,
		Cost:        room.Cost,
		PersonCount: room.PersonCount,
	}, nil
}
