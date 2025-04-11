package service_test

import (
	"testing"

	"github.com/spookycoincidence/hx-payments-system/internal/models"
	"github.com/spookycoincidence/hx-payments-system/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepo struct {
	mock.Mock
}

func (m *mockRepo) GetByUserID(userID uint) (*models.Wallet, error) {
	args := m.Called(userID)
	return args.Get(0).(*models.Wallet), args.Error(1)
}
func (m *mockRepo) Create(wallet *models.Wallet) (*models.Wallet, error) {
	args := m.Called(wallet)
	return args.Get(0).(*models.Wallet), args.Error(1)
}
func (m *mockRepo) Update(wallet *models.Wallet) error {
	return m.Called(wallet).Error(0)
}

func TestDeposit(t *testing.T) {
	mockRepo := new(mockRepo)
	s := service.NewWalletService(mockRepo)

	wallet := &models.Wallet{UserID: 1, Balance: 100}
	mockRepo.On("GetByUserID", uint(1)).Return(wallet, nil)
	mockRepo.On("Update", mock.Anything).Return(nil)

	err := s.Deposit(1, 50)
	assert.NoError(t, err)
	assert.Equal(t, float64(150), wallet.Balance)
}
