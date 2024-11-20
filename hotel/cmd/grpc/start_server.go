package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/h4x4d/go_hsse_hotels/hotel/internal/grpc/handlers"
	"google.golang.org/grpc"
)

func StartServer(group *sync.WaitGroup) {
	defer group.Done()

	port := os.Getenv("HOTEL_GRPC_PORT")
	host := os.Getenv("HOTEL_HOST")
	if port == "" {
		port = "8889"
	}
	if host == "" {
		host = "0.0.0.0"
	}

	address := host + ":" + port

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", address, err)
	}

	grpcServer := grpc.NewServer()
	handlers.Register(grpcServer)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Starting gRPC server on %s...", address)
		if err := grpcServer.Serve(listener); err != nil {
			log.Printf("gRPC server stopped: %v", err)
		}
	}()

	<-stop
	log.Println("Shutting down gRPC server...")
	gracefulStop(grpcServer)
}

func gracefulStop(server *grpc.Server) {
	const timeout = 2 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan struct{})
	go func() {
		server.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		log.Println("gRPC server gracefully stopped.")
	case <-ctx.Done():
		log.Println("Timeout reached; forcing gRPC server shutdown.")
		server.Stop()
	}
}
