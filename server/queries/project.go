package queries

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/krazy-code/devlink/models"
)

type ProjectQueries struct {
	Pool *pgxpool.Pool
}

func (q *ProjectQueries) GetProjects() ([]models.Project, error) {
	query := `
        SELECT id, user_id, title, description, tech_stack, project_url, created_at::text
        FROM projects AS t1
		ORDER BY created_at DESC
    `
	rows, err := q.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("error querying projects: %w", err)
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var project models.Project
		err := rows.Scan(
			&project.Id,
			&project.UserId,
			&project.Title,
			&project.Description,
			&project.Techstacks,
			&project.ProjectURL,
			&project.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning project row: %w", err)
		}
		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating user rows: %w", err)
	}

	return projects, nil
}

func (q *ProjectQueries) GetProject(id uuid.UUID) (*models.Project, error) {
	query := `
        SELECT t1.id, t2.id AS user_id, t2.email, t2.name, t1.bio, t1.location, t1.website, t1.github, t1.created_at::text
        FROM projects AS t1
		INNER JOIN users AS t2 ON t1.user_id = t2.id
		WHERE t1.id = $1
		ORDER BY t1.created_at DESC
    `

	project := &models.Project{}
	err := q.Pool.QueryRow(context.Background(), query, id).Scan(
		&project.Id,
		&project.UserId,
		&project.Description,
		&project.ProjectURL,
		&project.Techstacks,
		&project.Title,
		&project.CreatedAt,
	)
	if err != nil {
		return &models.Project{}, fmt.Errorf("error querying project: %w", err)
	}
	return project, nil
}

func (q *ProjectQueries) CreateProject(b *models.Project) (uuid.UUID, error) {
	query := `
        INSERT INTO projects (user_id, title, description, tech_stack, project_url)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
    `
	var id uuid.UUID
	err := q.Pool.QueryRow(context.Background(), query,
		&b.UserId,
		&b.Title,
		&b.Description,
		&b.Techstacks,
		&b.ProjectURL,
	).Scan(&id)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("error insert projects: %w", err)
	}

	return id, nil
}

func (q *ProjectQueries) UpdateProject(id uuid.UUID, b *models.Project) error {
	query := `
        UPDATE projects
        SET  title = $2, link = $3, tech_stacks = $4, descrition = $5
        WHERE id = $1
    `

	commandTag, err := q.Pool.Exec(context.Background(), query, id, b.Title, b.ProjectURL, b.Techstacks, b.Description)
	if err != nil {
		return fmt.Errorf("error updating developer: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no developer found with id %d", id)
	}

	return nil
}

func (q *ProjectQueries) DeleteProject(id uuid.UUID) error {
	query := `
        DELETE FROM projects
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
