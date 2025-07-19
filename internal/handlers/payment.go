package handlers

import (
	"net/http"

	"go-pet-shop/internal/storage"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	repo storage.Storage
}

func NewPaymentHandler(repo storage.Storage) *PaymentHandler {
	return &PaymentHandler{repo: repo}
}

func (h *PaymentHandler) ProcessPayment(c *gin.Context) {
	var paymentRequest struct {
		OrderID int     `json:"order_id"`
		Amount  float64 `json:"amount"`
		Method  string  `json:"method"`
	}

	if err := c.ShouldBindJSON(&paymentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Логика обработки платежа
	err := h.repo.ProcessPayment(paymentRequest.OrderID, paymentRequest.Amount, paymentRequest.Method)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Payment failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment successful"})
}
