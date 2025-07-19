package handlers

import (
	"net/http"

	"go-pet-shop/internal/storage"
	"go-pet-shop/models"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	Storage storage.Storage
}

func NewOrderHandler(storage storage.Storage) *OrderHandler {
	return &OrderHandler{
		Storage: storage,
	}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	orderID, err := h.Storage.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order_id": orderID})
}

func (h *OrderHandler) GetOrdersByUserEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	orders, err := h.Storage.GetOrdersByUserEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
