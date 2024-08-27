package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/tehuticode/subodai-finance/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

func InitDB() error {
    connStr := "user=sub0x dbname=subodai_finance password=password sslmode=disable"
    var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        return fmt.Errorf("error opening database: %w", err)
    }

    if err = DB.Ping(); err != nil {
        return fmt.Errorf("error connecting to database: %w", err)
    }

    fmt.Println("Successfully connected to database")
    return nil
}

// CreateUser inserts a new user into the database
func CreateUser(username, email, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	var user models.User
	err = DB.QueryRow(
		"INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id, username, email, created_at",
		username, email, string(hashedPassword),
	).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Add other database operations here as needed
