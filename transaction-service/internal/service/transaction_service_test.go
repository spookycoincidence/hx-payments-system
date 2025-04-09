package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	// Simular la creación de una transacción
	service := &TransactionService{}
	tx := service.CreateTransaction("user-123", 100.0)

	assert.NotNil(t, tx)
	assert.Equal(t, "user-123", tx.UserID)
	assert.Equal(t, 100.0, tx.Amount)
}
