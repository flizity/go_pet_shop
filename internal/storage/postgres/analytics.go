package postgres

import (
	"go-pet-shop/models"

	"github.com/jmoiron/sqlx"
)

type AnalyticsRepository struct {
	db *sqlx.DB
}

func NewAnalyticsRepository(db *sqlx.DB) *AnalyticsRepository {
	return &AnalyticsRepository{db: db}
}

func (r *AnalyticsRepository) GetPopularProducts() ([]models.PopularProduct, error) {
	var products []models.PopularProduct
	err := r.db.Select(&products, `
		SELECT p.id as product_id, p.name, SUM(oi.quantity) as sold
		FROM products p
		JOIN order_items oi ON p.id = oi.product_id
		GROUP BY p.id, p.name
		ORDER BY sold DESC
		LIMIT 10
	`)
	return products, err
}
