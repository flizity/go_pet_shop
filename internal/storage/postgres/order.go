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

func (r *OrderRepository) GetOrderItemsByOrderID(orderID int) ([]models.OrderItem, error) {
	var items []models.OrderItem
	err := r.db.Select(&items, "SELECT order_id, product_id, quantity FROM order_items WHERE order_id = $1", orderID)
	return items, err
}

func (r *OrderRepository) GetUserOrderHistory(email string) ([]models.OrderDetail, error) {
	var orderDetails []models.OrderDetail
	err := r.db.Select(&orderDetails, `
		SELECT o.id as order_id, o.user_email, o.total_price, o.created_at, 
		       oi.product_id, oi.quantity
		FROM orders o
		JOIN order_items oi ON o.id = oi.order_id
		WHERE o.user_email = $1
	`, email)
	return orderDetails, err
}

func (r *OrderRepository) PlaceOrder(userEmail string, items []models.PlaceOrder) (int, error) {
	// Начало транзакции
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	// Вставка заказа
	var orderID int
	err = tx.QueryRow(
		"INSERT INTO orders (user_email, total_price, created_at) VALUES ($1, $2, NOW()) RETURNING id",
		userEmail, calculateTotalPriceFromOrderItems(items),
	).Scan(&orderID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// Вставка позиций заказа
	for _, item := range items {
		for _, orderItem := range item.Items {
			_, err = tx.Exec(
				"INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)",
				orderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price,
			)
			if err != nil {
				tx.Rollback()
				return 0, err
			}
		}
	}

	// Завершение транзакции
	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return orderID, nil
}

func calculateTotalPriceFromOrderItems(items []models.PlaceOrder) float64 {
	total := 0.0
	for _, item := range items {
		for _, orderItem := range item.Items {
			total += orderItem.Price * float64(orderItem.Quantity)
		}
	}
	return total
}
