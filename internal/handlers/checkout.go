package handlers

import (
	"net/http"

	"go-pet-shop/internal/storage"
	"go-pet-shop/models"

	"github.com/gin-gonic/gin"
)

type CheckoutHandler struct {
	repo storage.Storage
}

func NewCheckoutHandler(repo storage.Storage) *CheckoutHandler {
	return &CheckoutHandler{repo: repo}
}

func (h *CheckoutHandler) PlaceOrder(c *gin.Context) {
	var request struct {
		UserEmail string             `json:"user_email"`
		Items     []models.OrderItem `json:"items"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert []models.OrderItem to []models.PlaceOrder
	var placeOrders []models.PlaceOrder
	for _, item := range request.Items {
		placeOrders = append(placeOrders, models.PlaceOrder{
			UserEmail: request.UserEmail,
			Items:     []models.OrderItem{item},
		})
	}

	orderID, err := h.repo.PlaceOrder(request.UserEmail, placeOrders)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"order_id": orderID})
}
