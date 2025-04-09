package service

import "fmt"

type NotificationService struct{}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (n *NotificationService) SendNotification(userID, message string) {
	fmt.Printf("🔔 Notificación enviada al usuario %s: %s\n", userID, message)
}
