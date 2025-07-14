package postgres

import (
	"go-pet-shop/models"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateTransaction(tx models.Transaction) (int, error) {
	var id int
	err := r.db.QueryRow(
		`INSERT INTO transactions (order_id, amount, status, created_at) VALUES ($1, $2, $3, $4) RETURNING id`,
		tx.OrderID, tx.Amount, tx.Status, tx.CreatedAt,
	).Scan(&id)
	return id, err
}

func (r *TransactionRepository) GetTransactionByID(id int) (models.Transaction, error) {
	var tx models.Transaction
	err := r.db.Get(&tx, "SELECT id, order_id, amount, status, created_at FROM transactions WHERE id = $1", id)
	return tx, err
}

func (r *TransactionRepository) GetTransactionsByOrderID(orderID int) ([]models.Transaction, error) {
	var txs []models.Transaction
	err := r.db.Select(&txs, "SELECT id, order_id, amount, status, created_at FROM transactions WHERE order_id = $1", orderID)
	return txs, err
}
