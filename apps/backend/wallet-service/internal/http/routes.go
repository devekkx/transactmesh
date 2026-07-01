package http

import "github.com/gin-gonic/gin"

func (h *Handler) RegisterRoutes(r *gin.Engine) {

	wallet := r.Group("/wallet")
	{
		wallet.POST("/debit", h.Debit)
	}
}
