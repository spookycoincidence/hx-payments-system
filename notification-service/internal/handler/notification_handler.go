package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notification-service/internal/service"
)

type NotificationHandler struct {
	service *service.NotificationService
}

func NewNotificationHandler(s *service.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		service: s,
	}
}

type notificationRequest struct {
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}

func (h *NotificationHandler) SendNotification(w http.ResponseWriter, r *http.Request) {
	var req notificationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := h.service.SendNotification(req.UserID, req.Message)
	if err != nil {
		http.Error(w, "failed to send notification", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Notification sent to user %s", req.UserID)
}
