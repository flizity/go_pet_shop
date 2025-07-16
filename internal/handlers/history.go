package handlers

import (
	"net/http"

	"go-pet-shop/internal/storage"

	"github.com/gin-gonic/gin"
)

type HistoryHandler struct {
	repo storage.Storage
}

func NewHistoryHandler(repo storage.Storage) *HistoryHandler {
	return &HistoryHandler{repo: repo}
}

func (h *HistoryHandler) GetUserOrderHistory(c *gin.Context) {
	email := c.Param("email")

	history, err := h.repo.GetUserOrderHistory(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, history)
}
