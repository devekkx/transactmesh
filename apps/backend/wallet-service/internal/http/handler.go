package http

import "github.com/devekkx/transactmesh/wallet-service/internal/wallet"

type Handler struct {
	service *wallet.Service
}

func NewHandler(s *wallet.Service) *Handler {
	return &Handler{service: s}
}

type debitRequest struct {
	WalletID string `json:"wallet_id" binding:"required"`
	Amount   int64  `json:"amount" binding:"required,gt=0"`
}
