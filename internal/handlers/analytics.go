package handlers

import (
	"net/http"

	"go-pet-shop/internal/storage"

	"github.com/gin-gonic/gin"
)

type AnalyticsHandler struct {
	repo storage.Storage
}

func NewAnalyticsHandler(repo storage.Storage) *AnalyticsHandler {
	return &AnalyticsHandler{repo: repo}
}

func (h *AnalyticsHandler) GetPopularProducts(c *gin.Context) {
	products, err := h.repo.GetPopularProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
