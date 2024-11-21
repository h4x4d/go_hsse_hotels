package utils

import (
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func ConnectTo(host *string, port *string) (*grpc.ClientConn, error) {
	return grpc.NewClient(fmt.Sprintf("%s:%s", *host, *port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func ConnectToHotel() (*grpc.ClientConn, error) {
	port := os.Getenv("HOTEL_GRPC_PORT")
	if port == "" {
		return nil, errors.New("HOTEL port is not not found")
	}
	host := "hotel"
	return ConnectTo(&host, &port)
}
