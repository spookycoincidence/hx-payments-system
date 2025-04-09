package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendNotification(t *testing.T) {
	service := &NotificationService{}

	err := service.SendNotification("user-123", "Test message")

	assert.NoError(t, err)
}
