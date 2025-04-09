package service

import (
	"fmt"
	"time"
	"user-service/internal/model"
)

type TransactionService struct{}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (ts *TransactionService) ProcessTransaction(userID string, amount float64) model.Transaction {
	transaction := model.Transaction{
		ID:        fmt.Sprintf("txn_%d", time.Now().UnixNano()),
		UserID:    userID,
		Amount:    amount,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	// Simula actualizar el wallet (sin llamada real)
	fmt.Printf("💰 Simulando actualización del wallet de usuario %s con monto %.2f\n", userID, amount)

	// Simula notificación (sin conexión real)
	fmt.Printf("📨 Enviando notificación a usuario %s sobre transacción %s\n", userID, transaction.ID)

	return transaction
}
