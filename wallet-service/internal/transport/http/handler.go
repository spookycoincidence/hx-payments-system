package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/spookycoincidence/hx-payments-system/internal/service"
)

type WalletHandler struct {
	service service.WalletService
}

func NewWalletHandler(service service.WalletService) *WalletHandler {
	return &WalletHandler{service: service}
}

func (h *WalletHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/wallets", h.CreateWallet).Methods("POST")
	r.HandleFunc("/wallets/{userID}", h.GetWallet).Methods("GET")
	r.HandleFunc("/wallets/{userID}/deposit", h.Deposit).Methods("POST")
	r.HandleFunc("/wallets/{userID}/withdraw", h.Withdraw).Methods("POST")
}

func (h *WalletHandler) CreateWallet(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		UserID uint `json:"user_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	wallet, err := h.service.CreateWallet(payload.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	json.NewEncoder(w).Encode(wallet)
}

func (h *WalletHandler) GetWallet(w http.ResponseWriter, r *http.Request) {
	userID, err := parseUserID(r)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	wallet, err := h.service.GetWallet(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(wallet)
}

func (h *WalletHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	userID, err := parseUserID(r)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	var payload struct {
		Amount float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil || payload.Amount <= 0 {
		http.Error(w, "invalid amount", http.StatusBadRequest)
		return
	}

	if err := h.service.Deposit(userID, payload.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *WalletHandler) Withdraw(w http.ResponseWriter, r *http.Request) {
	userID, err := parseUserID(r)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	var payload struct {
		Amount float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil || payload.Amount <= 0 {
		http.Error(w, "invalid amount", http.StatusBadRequest)
		return
	}

	if err := h.service.Withdraw(userID, payload.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func parseUserID(r *http.Request) (uint, error) {
	vars := mux.Vars(r)
	idStr := vars["userID"]
	id, err := strconv.Atoi(idStr)
	return uint(id), err
}
