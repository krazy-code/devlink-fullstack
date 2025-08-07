package queries

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/krazy-code/devlink/models"
)

type UserQueries struct {
	Pool *pgxpool.Pool
}

func (q *UserQueries) GetUsers() ([]models.User, error) {
	query := `
        SELECT id, name, email, password_hash, created_at::text
        FROM users
		ORDER BY created_at DESC
    `
	rows, err := q.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("error querying users: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.PasswordHash,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning user row: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating user rows: %w", err)
	}

	return users, nil
}

func (q *UserQueries) CreateUser(t *models.User) (int, error) {
	query := `
        INSERT INTO users (name, email)
        VALUES ($1, $2)
        RETURNING id
    `

	var id int
	err := q.Pool.QueryRow(context.Background(), query, t.Name, t.Email).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error creating user: %w", err)
	}

	fmt.Printf("Created user with ID: %d\n", id)
	return id, nil
}
