package main

import (
	"log"
	"os"

	"go-pet-shop/internal/config"
	"go-pet-shop/internal/handlers"
	"go-pet-shop/internal/storage/postgres"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig("E:\\VS CODE PROJECTS\\go_pet_shop\\config\\local.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := postgres.NewPostgresConnection(cfg.DB.URL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	storage := postgres.NewStorage(db)
	productHandler := handlers.NewProductHandler(storage.ProductRepository)
	statusHandler := handlers.NewStatusHandler()
	paymentHandler := handlers.NewPaymentHandler(storage)
	userHandler := handlers.NewUserHandler(storage)
	orderHandler := handlers.NewOrderHandler(storage)
	checkoutHandler := handlers.NewCheckoutHandler(storage)
	historyHandler := handlers.NewHistoryHandler(storage)
	analyticsHandler := handlers.NewAnalyticsHandler(storage)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Server is running")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/status", statusHandler.CheckStatus)
	r.GET("/products", productHandler.GetAllProducts)
	r.POST("/products", productHandler.CreateProduct)
	r.POST("/payments", paymentHandler.ProcessPayment)
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users", userHandler.GetAllUsers)
	r.POST("/orders", orderHandler.CreateOrder)
	r.GET("/orders", orderHandler.GetOrdersByUserEmail)
	r.POST("/checkout", checkoutHandler.PlaceOrder)
	r.GET("/history", historyHandler.GetUserOrderHistory)
	r.GET("/analytics/popular-products", analyticsHandler.GetPopularProducts)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
