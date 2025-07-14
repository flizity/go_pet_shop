package models

type PopularProduct struct {
	ProductID int    `json:"product_id" db:"product_id"`
	Name      string `json:"name" db:"name"`
	Sold      int    `json:"sold" db:"sold"`
}
