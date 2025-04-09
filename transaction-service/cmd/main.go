package main

import (
	"log"
	"net/http"

	"user-service/internal/handler"
	"user-service/internal/service"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	transactionService := service.NewTransactionService()
	transactionHandler := handler.NewTransactionHandler(transactionService)

	r.HandleFunc("/transactions", transactionHandler.CreateTransaction).Methods("POST")

	log.Println("🚀 Transaction service corriendo en http://localhost:8083")
	log.Fatal(http.ListenAndServe(":8083", r))
}
