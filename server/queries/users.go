package queries

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/krazy-code/devlink/models"
	"golang.org/x/crypto/bcrypt"
)

type UserQueries struct {
	Pool *pgxpool.Pool
}

func (q *UserQueries) GetUsers() ([]models.User, error) {
	query := `
        SELECT id, name, email, created_at::text
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
			&user.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning user arow: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating user rows: %w", err)
	}

	return users, nil
}

func (q *UserQueries) GetUser(id uuid.UUID) (*models.UserDetail, error) {
	query := `
		SELECT t1.id, t1.name, t1.email, 
		COALESCE(t2.github_url, ''), COALESCE(t2.bio, ''), COALESCE(t2.website_url, ''),
		COALESCE(t2.linkedin_url, ''), COALESCE(t2.location, ''),
		COALESCE(t1.avatar, ''), t1.created_at::text
		FROM users as t1 
		LEFT JOIN developers as t2 ON t2.user_id = t1.id
		WHERE t1.id = $1
	`

	user := &models.UserDetail{}
	err := q.Pool.QueryRow(context.Background(), query, id).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.GithubURL,
		&user.Bio,
		&user.WebsiteURL,
		&user.LinkedinURL,
		&user.Location,
		&user.Avatar,
		&user.CreatedAt,
	)
	if err != nil {
		return &models.UserDetail{}, fmt.Errorf("error querying user: %w", err)
	}
	return user, nil
}

func (q *UserQueries) CreateUser(b *models.User) (int, error) {
	query := `
        INSERT INTO users (name, email, passwordhash)
        VALUES ($1, $2, $3)
        RETURNING id
    `

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(b.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("failed to hash password")
	}

	var id int

	if err := q.Pool.QueryRow(context.Background(), query, b.Name, b.Email, hashedPassword).Scan(&id); err != nil {
		return 0, fmt.Errorf("error creating user: %w", err)
	}

	fmt.Printf("Created user with ID: %d\n", id)
	return id, nil
}

func (q *UserQueries) UpdateUser(id uuid.UUID, b *models.User) error {
	query := `
        UPDATE users
        SET  name = $2, email = $3, bio = $4, avatar = $5
        WHERE id = $1
    `

	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(b.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return fmt.Errorf("failed to hash password")
	// }

	commandTag, err := q.Pool.Exec(context.Background(), query, id, b.Name, b.Email, b.Bio, b.Avatar)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no user found with id %d", id)
	}

	return nil
}

func (q *UserQueries) DeleteUser(id uuid.UUID) error {
	query := `
        DELETE FROM users
        WHERE id = $1
    `

	commandTag, err := q.Pool.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no user found with id %d", id)
	}

	return nil
}
