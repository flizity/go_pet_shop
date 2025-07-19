package tests

import (
	"go-pet-shop/internal/storage/postgres"
	"go-pet-shop/models"
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sqlx.DB {
	dsn := os.Getenv("TEST_DB_DSN")
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}

	_, err = db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE")
	if err != nil {
		log.Fatalf("Failed to clean test database: %v", err)
	}

	return db
}

func TestUserRepository_CreateUser(t *testing.T) {
	db := setupTestDB()
	repo := postgres.NewUserRepository(db)

	user := models.User{
		Name:  "Test User",
		Email: "test@example.com",
	}

	id, err := repo.CreateUser(user)

	assert.NoError(t, err)
	assert.Greater(t, id, 0)
}
