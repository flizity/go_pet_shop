package handlers

import (
	"net/http"

	"go-pet-shop/internal/storage"
	"go-pet-shop/models"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Storage storage.Storage
}

func NewUserHandler(storage storage.Storage) *UserHandler {
	return &UserHandler{
		Storage: storage,
	}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.Storage.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	id, err := h.Storage.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	user.ID = id
	c.JSON(http.StatusCreated, user)
}
