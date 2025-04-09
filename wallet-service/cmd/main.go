package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/huilen/hx-payments-system/wallet-service/internal/db"
	"github.com/huilen/hx-payments-system/wallet-service/internal/models"
	"github.com/huilen/hx-payments-system/wallet-service/internal/repository"
	"github.com/huilen/hx-payments-system/wallet-service/internal/service"
	transport "github.com/huilen/hx-payments-system/wallet-service/internal/transport/http"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Connect()
	db.DB.AutoMigrate(&models.Wallet{})

	r := gin.Default()

	repo := repository.NewWalletRepository()
	svc := service.NewWalletService(repo)
	transport.NewWalletHandler(r, svc)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
