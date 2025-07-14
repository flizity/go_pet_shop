package models

import "go-pet-shop/models"

type Storage interface {
	// Пользователи
	CreateUser(user models.User) error
	GetUserByEmail(email string) (models.User, error)
	GetAllUsers() ([]models.User, error)

	// Товары
	CreateProduct(product models.Product) error
	GetProductByID(id int) (models.Product, error)
	GetAllProducts() ([]models.Product, error)
	UpdateProduct(product models.Product) error
	DeleteProduct(id int) error

	// Заказы и позиции
	CreateOrder(order models.Order) (int, error)
	GetOrdersByUserEmail(email string) ([]models.Order, error)
	GetOrderItemsByOrderID(orderID int) ([]models.OrderItem, error)
	AddOrderItem(orderItem models.OrderItem) error

	// Оформление заказа в транзакции
	PlaceOrder(userEmail string, items []models.PlaceOrder) (orderID int, err error)

	// История заказов
	GetUserOrderHistory(email string) ([]models.OrderDetail, error)

	// Аналитика
	GetPopularProducts() ([]models.PopularProduct, error)
}
