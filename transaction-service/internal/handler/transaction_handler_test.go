package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-service/internal/service"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransactionHandler(t *testing.T) {
	transactionService := &service.TransactionService{}
	handler := NewTransactionHandler(transactionService)

	body := map[string]interface{}{
		"user_id": "user-123",
		"amount":  200.0,
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.CreateTransaction(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Contains(t, rr.Body.String(), "transaction_id")
}
