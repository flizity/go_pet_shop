package postgres

import (
	"go-pet-shop/models"

	"github.com/jmoiron/sqlx"
)

type OrderDetailRepository struct {
	db *sqlx.DB
}

func NewOrderDetailRepository(db *sqlx.DB) *OrderDetailRepository {
	return &OrderDetailRepository{db: db}
}

func (r *OrderDetailRepository) GetOrderDetailsByOrderID(orderID int) ([]models.OrderDetail, error) {
	var details []models.OrderDetail
	err := r.db.Select(&details, `
		SELECT o.id as order_id, o.user_email, o.total_price, o.created_at, oi.product_id, oi.quantity
		FROM orders o
		LEFT JOIN order_items oi ON o.id = oi.order_id
		WHERE o.id = $1
		ORDER BY o.created_at DESC, o.id, oi.id
	`, orderID)
	return details, err
}
