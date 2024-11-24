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
	gen.UnimplementedHotelServer
}

func Register(gRPCServer *grpc.Server) {
	gen.RegisterHotelServer(gRPCServer, &GRPCServer{})
}

func (serverApi *GRPCServer) GetHotel(
	ctx context.Context, in *gen.HotelRequest) (*gen.HotelResponse, error) {
	log.Println("in func")
	hotel, err := services.GetHotelByID(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}
	if hotel == nil {
		return nil, status.Errorf(codes.NotFound, "room %d not found", in.Id)
	}
	return &gen.HotelResponse{
		Id:         hotel.ID,
		Name:       *hotel.Name,
		City:       *hotel.City,
		Address:    *hotel.Address,
		HotelClass: hotel.HotelClass,
		Cost:       hotel.Cost,
	}, nil
}
