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

func (r *ProductRepository) CreateProduct(p models.Product) (int, error) {
	var id int
	err := r.db.QueryRow(
		`INSERT INTO products (name, price, stock) VALUES ($1, $2, $3) RETURNING id`,
		p.Name, p.Price, p.Stock,
	).Scan(&id)
	return id, err
}

func (r *ProductRepository) GetProductByID(id int) (models.Product, error) {
	var product models.Product
	err := r.db.Get(&product, "SELECT id, name, price, stock FROM products WHERE id = $1", id)
	return product, err
}

func (r *ProductRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Select(&products, "SELECT id, name, price, stock FROM products")
	return products, err
}

func (r *ProductRepository) UpdateProduct(p models.Product) error {
	_, err := r.db.Exec(
		"UPDATE products SET name = $1, price = $2, stock = $3 WHERE id = $4",
		p.Name, p.Price, p.Stock, p.ID,
	)
	return err
}

func (r *ProductRepository) DeleteProduct(id int) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id = $1", id)
	return err
}

func (r *ProductRepository) GetPopularProducts() ([]models.PopularProduct, error) {
	var products []models.PopularProduct
	err := r.db.Select(&products, "SELECT product_id, COUNT(*) as popularity FROM order_items GROUP BY product_id ORDER BY popularity DESC")
	return products, err
}
