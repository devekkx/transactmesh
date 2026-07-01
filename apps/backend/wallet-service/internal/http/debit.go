package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Debit(c *gin.Context) {

	var req debitRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.service.Debit(c.Request.Context(), req.WalletID, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
