package models

import "time"

type Order struct {
	ID         int       `json:"id" db:"id"`
	UserEmail  string    `json:"user_email" db:"user_email"`
	TotalPrice int       `json:"total_price" db:"total_price"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}
