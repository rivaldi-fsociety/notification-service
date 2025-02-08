package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/rivaldi-fsociety/notification-service/proto"
)

// NotificationHandler processes notifications
type NotificationService struct {
	proto.UnimplementedNotificationServiceServer
}

// SendNotification sends an email or SMS
func (s *NotificationService) SendNotification(ctx context.Context, req *proto.SendNotificationRequest) (*proto.SendNotificationResponse, error) {
	log.Printf("Sending notification to %s: %s", req.UserId, req.Message)
	// Simulate sending email
	fmt.Printf("📩 Notification sent to %s: %s\n", req.UserId, req.Message)
	return &proto.SendNotificationResponse{Success: true}, nil
}
