package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	httpHandler "github.com/huilen/hx-payments-system/wallet-service/internal/transport/http"
	"github.com/stretchr/testify/assert"
)

type dummyService struct{}

func (d *dummyService) GetWallet(userID uint) (*models.Wallet, error) {
	return &models.Wallet{UserID: userID, Balance: 100}, nil
}

// Implementar dummy deposit, withdraw, etc.

func TestGetWallet(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	svc := &dummyService{}
	httpHandler.NewWalletHandler(r, svc)

	req, _ := http.NewRequest("GET", "/wallet/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
