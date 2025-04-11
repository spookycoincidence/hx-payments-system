package repository_test

import (
	"testing"

	"github.com/spookycoincidence/hx-payments-system/internal/db"
	"github.com/spookycoincidence/hx-payments-system/internal/models"
	"github.com/spookycoincidence/hx-payments-system/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndGetWallet(t *testing.T) {
	db.Connect() // fake: no ejecutarlo en realidad
	r := repository.NewWalletRepository()

	userID := uint(99)
	w := &models.Wallet{UserID: userID, Balance: 0}
	r.Create(w)

	result, err := r.GetByUserID(userID)
	assert.NoError(t, err)
	assert.Equal(t, userID, result.UserID)
}
