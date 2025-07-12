package models

import "go-pet-shop/models"

type Storage interface {
	// Products
	CreateProduct(p models.Product) error
	GetProductByID(id int) (models.Product, error)
	PlaceOrder(userEmail string, items []models.Order_items) (int, error)
	// Orders
	CreateOrder(order models.Orders) (int, error)
	GetOrderByID(id int) (models.Orders, error)
	GetOrdersByUserEmail(email string) ([]models.Orders, error)

	// Transactions
	CreateTransaction(tx models.Transaction) (int, error)
	GetTransactionByID(id int) (models.Transaction, error)
	GetTransactionsByOrderID(orderID int) ([]models.Transaction, error)

	// Users
	CreateUser(user models.Users) (int, error)
	GetUserByEmail(email string) (models.Users, error)
}
