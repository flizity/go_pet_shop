package postgres

import (
	"go-pet-shop/models"

	"github.com/jmoiron/sqlx"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(order models.Order) (int, error) {
	var id int
	err := r.db.QueryRow(
		`INSERT INTO orders (user_email, total_price, created_at) VALUES ($1, $2, $3) RETURNING id`,
		order.UserEmail, order.TotalPrice, order.CreatedAt,
	).Scan(&id)
	return id, err
}

func (r *OrderRepository) GetOrdersByUserEmail(email string) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Select(&orders, "SELECT id, user_email, total_price, created_at FROM orders WHERE user_email = $1", email)
	return orders, err
}
