package models

import "time"

type OrderDetail struct {
	OrderID    int       `db:"order_id" json:"order_id"`
	UserEmail  string    `db:"user_email" json:"user_email"`
	TotalPrice int       `db:"total_price" json:"total_price"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	ProductID  int       `db:"product_id" json:"product_id"`
	Quantity   int       `db:"quantity" json:"quantity"`
}
