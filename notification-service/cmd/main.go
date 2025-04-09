package main

import (
	"log"
	"time"

	"notification-service/internal/service"
)

func main() {
	ns := service.NewNotificationService()

	log.Println("🚀 Notification service corriendo (simulado)")
	time.Sleep(1 * time.Second)
	ns.SendNotification("user123", "¡Tu transacción fue exitosa!")
}
