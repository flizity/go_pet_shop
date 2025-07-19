package postgres

import (
	"go-pet-shop/models"

	"github.com/jmoiron/sqlx"
)

type Storage struct {
	ProductRepository     *ProductRepository
	OrderRepository       *OrderRepository
	UserRepository        *UserRepository
	TransactionRepository *TransactionRepository
	DB                    *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		ProductRepository:     NewProductRepository(db),
		OrderRepository:       NewOrderRepository(db),
		UserRepository:        NewUserRepository(db),
		TransactionRepository: NewTransactionRepository(db),
		DB:                    db,
	}
}

func (s *Storage) ProcessPayment(orderID int, amount float64, method string) error {
	_, err := s.DB.Exec(
		`INSERT INTO transactions (order_id, amount, method, status) VALUES ($1, $2, $3, 'completed')`,
		orderID, amount, method,
	)
	return err
}

func (s *Storage) AddOrderItem(orderItem models.OrderItem) error {
	_, err := s.DB.Exec(
		`INSERT INTO order_items (order_id, product_id, quantity) VALUES ($1, $2, $3)`,
		orderItem.OrderID, orderItem.ProductID, orderItem.Quantity,
	)
	return err
}

func (s *Storage) CreateOrder(order models.Order) (int, error) {
	return s.OrderRepository.CreateOrder(order)
}

func (s *Storage) CreateProduct(product models.Product) (int, error) {
	return s.ProductRepository.CreateProduct(product)
}

func (s *Storage) CreateUser(user models.User) (int, error) {
	return s.UserRepository.CreateUser(user)
}

func (s *Storage) CreateTransaction(tx models.Transaction) (int, error) {
	return s.TransactionRepository.CreateTransaction(tx)
}

func (s *Storage) DeleteProduct(id int) error {
	return s.ProductRepository.DeleteProduct(id)
}

func (s *Storage) DeleteUser(id string) error {
	return s.UserRepository.DeleteUser(id)
}

func (s *Storage) GetAllProducts() ([]models.Product, error) {
	return s.ProductRepository.GetAllProducts()
}

func (s *Storage) GetAllUsers() ([]models.User, error) {
	return s.UserRepository.GetAllUsers()
}

func (s *Storage) GetOrderItemsByOrderID(orderID int) ([]models.OrderItem, error) {
	return s.OrderRepository.GetOrderItemsByOrderID(orderID)
}

func (s *Storage) GetOrdersByUserEmail(email string) ([]models.Order, error) {
	return s.OrderRepository.GetOrdersByUserEmail(email)
}

func (s *Storage) GetPopularProducts() ([]models.PopularProduct, error) {
	return s.ProductRepository.GetPopularProducts()
}

func (s *Storage) GetProductByID(id int) (models.Product, error) {
	return s.ProductRepository.GetProductByID(id)
}

func (s *Storage) GetTransactionByID(id int) (models.Transaction, error) {
	return s.TransactionRepository.GetTransactionByID(id)
}

func (s *Storage) GetTransactionsByOrderID(orderID int) ([]models.Transaction, error) {
	return s.TransactionRepository.GetTransactionsByOrderID(orderID)
}

func (s *Storage) GetUserByEmail(email string) (models.User, error) {
	return s.UserRepository.GetUserByEmail(email)
}

func (s *Storage) GetUserOrderHistory(email string) ([]models.OrderDetail, error) {
	return s.OrderRepository.GetUserOrderHistory(email)
}

func (s *Storage) PlaceOrder(userEmail string, items []models.PlaceOrder) (orderID int, err error) {
	return s.OrderRepository.PlaceOrder(userEmail, items)
}

func (s *Storage) UpdateProduct(product models.Product) error {
	return s.ProductRepository.UpdateProduct(product)
}
