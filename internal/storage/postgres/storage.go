package postgres

import (
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	ProductRepository *ProductRepository
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		ProductRepository: NewProductRepository(db),
	}
}
