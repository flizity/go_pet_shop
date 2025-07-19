package postgres

import (
	"go-pet-shop/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user models.User) (int, error) {
	var id int
	err := r.db.QueryRow(
		`INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`,
		user.Name, user.Email,
	).Scan(&id)
	return id, err
}

func (r *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Get(&user, "SELECT id, name, email FROM users WHERE email = $1", email)
	return user, err
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Select(&users, "SELECT id, name, email FROM users")
	return users, err
}

func (r *UserRepository) DeleteUser(id string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
