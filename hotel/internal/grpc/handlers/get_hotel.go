package handlers

import (
	"context"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/database_service"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/grpc/gen"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"os"
)

type GRPCServer struct {
	Database *database_service.DatabaseService
	gen.UnimplementedHotelServer
}

func NewGRPCServer() (*GRPCServer, error) {
	db, err := database_service.NewDatabaseService(fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), "db", os.Getenv("POSTGRES_PORT"), os.Getenv("HOTEL_DB_NAME")))
	if err != nil {
		return nil, err
	}
	return &GRPCServer{Database: db}, nil
}

func Register(gRPCServer *grpc.Server) {
	server, err := NewGRPCServer()
	if err != nil {
		os.Exit(1)
	}
	gen.RegisterHotelServer(gRPCServer, server)
}

func (serverApi *GRPCServer) GetHotel(
	ctx context.Context, in *gen.HotelRequest) (*gen.HotelResponse, error) {

	tracer := otel.Tracer("Hotel")
	md, _ := metadata.FromIncomingContext(ctx)
	if len(md.Get("x-trace-id")) > 0 {
		traceIdString := md.Get("x-trace-id")[0]
		traceId, err := trace.TraceIDFromHex(traceIdString)
		if err != nil {
			return nil, err
		}
		spanContext := trace.NewSpanContext(trace.SpanContextConfig{
			TraceID: traceId,
		})
		ctx = trace.ContextWithSpanContext(ctx, spanContext)
	} else {
		ctx = context.Background()
	}
	_, span := tracer.Start(ctx, "get hotel")
	defer span.End()

	hotel, err := serverApi.Database.GetById(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}
	if hotel == nil {
		return nil, status.Errorf(codes.NotFound, "hotel %d not found", in.Id)
	}
	return &gen.HotelResponse{
		Id:         hotel.ID,
		Name:       *hotel.Name,
		City:       *hotel.City,
		Address:    *hotel.Address,
		HotelClass: hotel.HotelClass,
		Cost:       hotel.Cost,
		UserId:     hotel.UserID,
	}, nil
}
