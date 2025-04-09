package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"notification-service/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendNotificationHandler(t *testing.T) {
	service := &service.NotificationService{}
	handler := NewNotificationHandler(service)

	body := map[string]string{
		"user_id": "user-123",
		"message": "Hola desde test!",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/notifications", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.SendNotification(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Notification sent")
}
