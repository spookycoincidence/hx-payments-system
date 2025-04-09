package repository

import (
	"github.com/huilen/hx-payments-system/wallet-service/internal/db"
	"github.com/huilen/hx-payments-system/wallet-service/internal/models"
)

type WalletRepository interface {
	GetByUserID(userID uint) (*models.Wallet, error)
	Create(wallet *models.Wallet) error
	Update(wallet *models.Wallet) error
}

type walletRepository struct{}

func NewWalletRepository() WalletRepository {
	return &walletRepository{}
}

func (r *walletRepository) GetByUserID(userID uint) (*models.Wallet, error) {
	var wallet models.Wallet
	err := db.DB.Where("user_id = ?", userID).First(&wallet).Error
	return &wallet, err
}

func (r *walletRepository) Create(wallet *models.Wallet) error {
	return db.DB.Create(wallet).Error
}

func (r *walletRepository) Update(wallet *models.Wallet) error {
	return db.DB.Save(wallet).Error
}
