package main

import (
	"log"
	"net"

	"github.com/rivaldi-fsociety/notification-service/internal/handler"
	"github.com/rivaldi-fsociety/notification-service/proto"
	"google.golang.org/grpc"
)

func main() {
	// Start gRPC Server
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterNotificationServiceServer(grpcServer, &handler.NotificationService{})

	// Start Echo HTTP Server
	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(200, "Hello from Echo!")
	// })

	// Run gRPC in a goroutine
	log.Println("Starting gRPC server on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}

	// // Start Echo
	// log.Println("Starting Echo server on :8080")
	// if err := e.Start(":8080"); err != nil {
	// 	log.Fatalf("Failed to start Echo server: %v", err)
	// }
}
