package queries

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/krazy-code/devlink/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthQueries struct {
	Pool *pgxpool.Pool
}

func (q *AuthQueries) PostLogin(b *models.LoginRequest) (uuid.UUID, error) {
	query := `
		SELECT id, password_hash 
		FROM users 
		WHERE email=$1
	`
	var passwordHash string
	var userID uuid.UUID
	err := q.Pool.QueryRow(context.Background(),
		query, b.Email).Scan(&userID, &passwordHash)
	if err != nil {
		return userID, fmt.Errorf("invalid emails or password: %w", err)
	}

	if bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(b.Password)) != nil {
		return userID, fmt.Errorf("invalid emails or password: %w", err)
	}

	return userID, nil
}

func (q *AuthQueries) PostRegister(b *models.RegisterRequest) (uuid.UUID, error) {
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
		return uuid.New(), fmt.Errorf("email already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(b.Password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.New(), fmt.Errorf("failed to hash password")
	}

	var userID uuid.UUID

	if err := q.Pool.QueryRow(context.Background(), query, b.Name, b.Email, hashedPassword).Scan(&userID); err != nil {
		return uuid.New(), fmt.Errorf("invalid email or password: %w", err)
	}

	return userID, nil
}

func (q *AuthQueries) PostLogout(b models.LogoutRequest) (uuid.UUID, error) {
	userID, err := uuid.Parse(b.AccessTokenClaims["user_id"].(string))
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("logout error: %v", err)
	}

	if err := q.UpdateAccessToken(userID, ""); err != nil {
		return uuid.UUID{}, fmt.Errorf("logout error: %v", err)
	}

	return userID, nil
}

func (q *AuthQueries) UpdateAccessToken(id uuid.UUID, token string) (err error) {
	query := `
		UPDATE users 
		SET token = $2 
		WHERE id = $1
	`
	_, err = q.Pool.Exec(context.Background(), query, id, token)
	return
}
