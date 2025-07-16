package postgres

import (
	"fmt"
	"go-pet-shop/models"
)

func (s *PostgresStorage) PlaceOrder(userEmail string, items []models.OrderItem) (int, error) {
	tx, err := s.DB.Beginx()
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	for _, item := range items {
		res, err := tx.Exec(
			`UPDATE products 
			SET stock = stock - $1 
			WHERE id = $2 AND stock >= $1`,
			item.Quantity, item.ProductID,
		)
		if err != nil {
			return 0, fmt.Errorf("failed to update stock: %w", err)
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return 0, fmt.Errorf("failed to check stock update: %w", err)
		}

		if rowsAffected == 0 {
			return 0, fmt.Errorf("product %d is out of stock or quantity unavailable", item.ProductID)
		}
	}

	var totalPrice int
	err = tx.Select(&totalPrice, `
		SELECT SUM(p.price * i.quantity)
		FROM (SELECT $1::int as product_id, $2::int as quantity) AS i
		JOIN products p ON p.id = i.product_id`,
		items[0].ProductID, items[0].Quantity)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate total: %w", err)
	}

	var orderID int
	err = tx.QueryRow(
		`INSERT INTO orders (user_email, total_price) 
		VALUES ($1, $2)
		RETURNING id`,
		userEmail, totalPrice,
	).Scan(&orderID)
	if err != nil {
		return 0, fmt.Errorf("failed to create order: %w", err)
	}

	for _, item := range items {
		_, err = tx.Exec(
			`INSERT INTO order_items (order_id, product_id, quantity)
			VALUES ($1, $2, $3)`,
			orderID, item.ProductID, item.Quantity,
		)
		if err != nil {
			return 0, fmt.Errorf("failed to add order item: %w", err)
		}
	}

	_, err = tx.Exec(
		`INSERT INTO transactions (order_id, amount, status)
		VALUES ($1, $2, 'completed')`,
		orderID, totalPrice,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to create transaction: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return orderID, nil
}
