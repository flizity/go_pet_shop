package postgres

import (
	"go-pet-shop/models"

	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(p models.Product) error {
	_, err := r.db.NamedExec(`
        INSERT INTO products (name, price, stock) 
        VALUES (:name, :price, :stock)`, p)
	return err
}
