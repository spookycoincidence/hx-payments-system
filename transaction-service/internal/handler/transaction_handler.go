package handler

import (
	"encoding/json"
	"net/http"
	"user-service/internal/service"
)

type TransactionHandler struct {
	Service *service.TransactionService
}

func NewTransactionHandler(s *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{Service: s}
}

type TransactionRequest struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var req TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "❌ Error al parsear la solicitud", http.StatusBadRequest)
		return
	}

	transaction := h.Service.ProcessTransaction(req.UserID, req.Amount)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}
