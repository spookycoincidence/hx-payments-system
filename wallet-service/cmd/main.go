package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spookycoincidence/hx-payments-system/internal/db"
	"github.com/spookycoincidence/hx-payments-system/internal/models"
	"github.com/spookycoincidence/hx-payments-system/internal/repository"
	"github.com/spookycoincidence/hx-payments-system/internal/service"

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
