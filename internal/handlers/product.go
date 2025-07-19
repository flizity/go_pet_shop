package handlers

import (
	"net/http"

	"go-pet-shop/internal/storage"
	"go-pet-shop/models"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	repo storage.ProductRepository
}

func NewProductHandler(repo storage.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.repo.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.repo.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":    "Product created successfully",
		"product_id": id,
	})
}
