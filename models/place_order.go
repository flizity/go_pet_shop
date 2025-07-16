package models

type PlaceOrder struct {
	UserEmail string      `json:"user_email" binding:"required,email"`
	Items     []OrderItem `json:"items" binding:"required,gt=0"`
}
