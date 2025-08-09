package queries

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/krazy-code/devlink/models"
)

type DeveloperQueries struct {
	Pool *pgxpool.Pool
}

func (q *DeveloperQueries) GetDevelopers() ([]models.Developer, error) {
	query := `
        SELECT t1.id, t2.id AS user_id, t2.email, t2.name, t1.bio, t1.location, t1.website, t1.github, t1.created_at::text
        FROM developers AS t1
		INNER JOIN users AS t2 ON t1.user_id = t2.id
		ORDER BY t1.created_at DESC
    `
	rows, err := q.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("error querying developers: %w", err)
	}
	defer rows.Close()

	var developers []models.Developer
	for rows.Next() {
		var developer models.Developer
		err := rows.Scan(
			&developer.Id,
			&developer.UserId,
			&developer.Email,
			&developer.Name,
			&developer.Bio,
			&developer.Location,
			&developer.Website,
			&developer.Github,
			&developer.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning user row: %w", err)
		}
		developers = append(developers, developer)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating user rows: %w", err)
	}

	return developers, nil
}

func (q *DeveloperQueries) GetDeveloper(id int) (*models.Developer, error) {
	query := `
        SELECT t1.id, t2.id AS user_id, t2.email, t2.name, t1.bio, t1.location, t1.website, t1.github, t1.created_at::text
        FROM developers AS t1
		INNER JOIN users AS t2 ON t1.user_id = t2.id
		WHERE t1.id = $1
		ORDER BY t1.created_at DESC
    `

	developer := &models.Developer{}
	err := q.Pool.QueryRow(context.Background(), query, id).Scan(
		&developer.Id,
		&developer.UserId,
		&developer.Email,
		&developer.Name,
		&developer.Bio,
		&developer.Location,
		&developer.Website,
		&developer.Github,
		&developer.CreatedAt,
	)
	if err != nil {
		return &models.Developer{}, fmt.Errorf("error querying developer: %w", err)
	}
	return developer, nil
}

func (q *DeveloperQueries) CreateDeveloper(b *models.Developer) (int, error) {
	query := `
        INSERT INTO developers (user_id, bio, location, website, github)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
    `
	var id int
	err := q.Pool.QueryRow(context.Background(), query,
		&b.UserId,
		&b.Bio,
		&b.Location,
		&b.Website,
		&b.Github,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error insert developers: %w", err)
	}

	return id, nil
}

func (q *DeveloperQueries) UpdateDeveloper(id int, b *models.Developer) error {
	query := `
        UPDATE developers
        SET  github = $2, bio = $3, location = $4, website = $5
        WHERE id = $1
    `

	commandTag, err := q.Pool.Exec(context.Background(), query, id, b.Github, b.Bio, b.Location, b.Website)
	if err != nil {
		return fmt.Errorf("error updating developer: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no developer found with id %d", id)
	}

	return nil
}

func (q *DeveloperQueries) DeleteDeveloper(id int) error {
	query := `
        DELETE FROM developers
        WHERE id = $1
    `

	commandTag, err := q.Pool.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("error deleting developer: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no developer found with id %d", id)
	}

	return nil
}
