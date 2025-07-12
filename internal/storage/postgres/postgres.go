package postgres

import (
	"fmt"

	"go-pet-shop/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sqlx.DB
}

func New(databaseURL string) (*Storage, error) {
	db, err := sqlx.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

func (s *Storage) CreateProduct(p models.Product) error {
	_, err := s.db.Exec(
		`INSERT INTO products (name, price, stock) 
		VALUES ($1, $2, $3)`,
		p.Name, p.Price, p.Stock,
	)
	return err
}

// order.go
func (s *Storage) PlaceOrder(userEmail string, items []models.Order_items) (int, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var orderID int
	err = tx.QueryRow(
		`INSERT INTO orders (user_email, total_price) 
		VALUES ($1, 0) 
		RETURNING id`, userEmail,
	).Scan(&orderID)

	if err != nil {
		return 0, err
	}

	total := 0
	for _, item := range items {
		var stock, price int
		err = tx.QueryRow("SELECT stock, price FROM products WHERE id = $1", item.ProductID).Scan(&stock, &price)
		if err != nil {
			return 0, err
		}
		if stock < item.Quantity {
			return 0, fmt.Errorf("product %d not available", item.ProductID)
		}

		_, err = tx.Exec(
			`INSERT INTO order_items (order_id, product_id, quantity) 
			VALUES ($1, $2, $3)`,
			orderID, item.ProductID, item.Quantity,
		)
		if err != nil {
			return 0, err
		}

		_, err = tx.Exec(
			"UPDATE products SET stock = stock - $1 WHERE id = $2",
			item.Quantity, item.ProductID,
		)
		if err != nil {
			return 0, err
		}

		total += price * item.Quantity
	}

	_, err = tx.Exec(
		"UPDATE orders SET total_price = $1 WHERE id = $2",
		total, orderID,
	)
	if err != nil {
		return 0, err
	}

	_, err = tx.Exec(
		`INSERT INTO transactions (order_id, amount, status) 
		VALUES ($1, $2, 'completed')`,
		orderID, total,
	)
	if err != nil {
		return 0, err
	}

	return orderID, tx.Commit()
}
