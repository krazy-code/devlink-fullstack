package queries

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/krazy-code/devlink/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthQueries struct {
	Pool *pgxpool.Pool
}

func (q *AuthQueries) PostLogin(b *models.LoginRequest) (int, error) {
	query := `
		SELECT id, password_hash 
		FROM users 
		WHERE email=$1
	`
	var passwordHash string
	var userID int
	err := q.Pool.QueryRow(context.Background(),
		query, b.Email).Scan(&userID, &passwordHash)
	if err != nil {
		return 0, fmt.Errorf("invalid email or password: %w", err)
	}

	if bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(b.Password)) != nil {
		return 0, fmt.Errorf("invalid email or password: %w", err)
	}

	return userID, nil
}

func (q *AuthQueries) PostRegister(b *models.RegisterRequest) (int, error) {
	query := `
		INSERT INTO users (name,email, password_hash) 
		VALUES ($1, $2, $3)
		RETURNING id
	`
	queryGetUser := `
        SELECT email
        FROM users
		WHERE email=$1
    `
	err := q.Pool.QueryRow(context.Background(), queryGetUser, b.Email).Scan()
	if err == nil {
		return 0, fmt.Errorf("email already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(b.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("failed to hash password")
	}

	var userID int

	if err := q.Pool.QueryRow(context.Background(), query, b.Name, b.Email, hashedPassword).Scan(&userID); err != nil {
		return 0, fmt.Errorf("invalid email or password: %w", err)
	}

	return userID, nil
}

func (q *AuthQueries) PostLogout(b *models.RegisterRequest) (int, error) {
	query := `
		INSERT INTO users (name,email, password_hash) 
		VALUES ($1, $2, $3)
		RETURNING id
	`
	queryGetUser := `
        SELECT email
        FROM users
		WHERE email=$1
    `
	err := q.Pool.QueryRow(context.Background(), queryGetUser, b.Email).Scan()
	if err == nil {
		return 0, fmt.Errorf("email already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(b.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("failed to hash password")
	}

	var userID int

	if err := q.Pool.QueryRow(context.Background(), query, b.Name, b.Email, hashedPassword).Scan(&userID); err != nil {
		return 0, fmt.Errorf("invalid email or password: %w", err)
	}

	// if err := q.Pool.QueryRow(context.Background(), query, b.Name, b.Email, hashedPassword).Scan(&userID); err != nil {
	// 	return 0, fmt.Errorf("invalid email or password: %w", err)
	// }

	return userID, nil
}
