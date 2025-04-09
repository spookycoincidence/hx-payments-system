package service

import (
	"errors"
	"fmt"

	"github.com/huilen/hx-payments-system/wallet-service/internal/models"
	"github.com/huilen/hx-payments-system/wallet-service/internal/repository"
)

type WalletService interface {
	GetWallet(userID uint) (*models.Wallet, error)
	Deposit(userID uint, amount float64) error
	Withdraw(userID uint, amount float64) error
	CreateWallet(userID uint) (*models.Wallet, error)
}

type walletService struct {
	repo repository.WalletRepository
}

func NewWalletService(repo repository.WalletRepository) WalletService {
	return &walletService{
		repo: repo,
	}
}

func (s *walletService) GetWallet(userID uint) (*models.Wallet, error) {
	wallet, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("wallet not found for user %d", userID)
	}
	return wallet, nil
}

func (s *walletService) Deposit(userID uint, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	wallet, err := s.repo.GetByUserID(userID)
	if err != nil {
		return fmt.Errorf("wallet not found: %w", err)
	}

	wallet.Balance += amount
	return s.repo.Update(wallet)
}

func (s *walletService) Withdraw(userID uint, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	wallet, err := s.repo.GetByUserID(userID)
	if err != nil {
		return fmt.Errorf("wallet not found: %w", err)
	}

	if wallet.Balance < amount {
		return errors.New("insufficient funds")
	}

	wallet.Balance -= amount
	return s.repo.Update(wallet)
}

func (s *walletService) CreateWallet(userID uint) (*models.Wallet, error) {
	existing, err := s.repo.GetByUserID(userID)
	if err == nil && existing != nil {
		return nil, fmt.Errorf("wallet already exists for user %d", userID)
	}

	wallet := &models.Wallet{
		UserID:  userID,
		Balance: 0,
	}

	return s.repo.Create(wallet)
}
