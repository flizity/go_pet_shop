package app

import (
	"log"
	"os"

	"go-pet-shop/internal/config"
	"go-pet-shop/internal/handlers"
	"go-pet-shop/internal/storage/postgres"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("./config/local.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database connection
	db, err := postgres.NewPostgresConnection(cfg.DB.URL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repositories
	storage := &postgres.Storage{
		ProductRepository: postgres.NewProductRepository(db),
	}

	// Initialize handlers
	productHandler := handlers.NewProductHandler(storage.ProductRepository)
	statusHandler := handlers.NewStatusHandler()

	// Initialize Gin router
	r := gin.Default()

	// Define routes
	r.GET("/status", statusHandler.CheckStatus)
	r.GET("/products", productHandler.GetAllProducts)
	r.POST("/products", productHandler.CreateProduct)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s", port)
	r.Run(":" + port)
}
