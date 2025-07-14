package postgres

import (
	"go-pet-shop/models"

	"github.com/jmoiron/sqlx"
)

type OrderItemRepository struct {
	db *sqlx.DB
}

func (r *OrderItemRepository) AddOrderItem(orderItem models.OrderItem) error {
	_, err := r.db.Exec(
		`INSERT INTO order_items (order_id, product_id, quantity) VALUES ($1, $2, $3)`,
		orderItem.OrderID, orderItem.ProductID, orderItem.Quantity,
	)
	return err
}

func (r *OrderItemRepository) GetOrderItemsByOrderID(orderID int) ([]models.OrderItem, error) {
	var items []models.OrderItem
	err := r.db.Select(&items, "SELECT id, order_id, product_id, quantity FROM order_items WHERE order_id = $1", orderID)
	return items, err
}
