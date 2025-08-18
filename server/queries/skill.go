package queries

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/krazy-code/devlink/models"
)

type SkillQueries struct {
	Pool *pgxpool.Pool
}

func (q *SkillQueries) GetSkills() ([]models.Skill, error) {
	query := `
        SELECT id, user_id, name
        FROM skills 
		ORDER BY created_at DESC
    `
	rows, err := q.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("error querying skills: %w", err)
	}
	defer rows.Close()

	var skills []models.Skill
	for rows.Next() {
		var skill models.Skill
		err := rows.Scan(
			&skill.Id,
			&skill.UserId,
			&skill.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning skill row: %w", err)
		}
		skills = append(skills, skill)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating skill rows: %w", err)
	}

	return skills, nil
}

func (q *SkillQueries) GetSkill(id uuid.UUID) (*models.Skill, error) {
	query := `
        SELECT id, user_id, name
        FROM skills
		WHERE id = $1
    `

	skill := &models.Skill{}
	err := q.Pool.QueryRow(context.Background(), query, id).Scan(
		&skill.Id,
		&skill.UserId,
		&skill.Name,
	)
	if err != nil {
		return &models.Skill{}, fmt.Errorf("error querying skill: %w", err)
	}
	return skill, nil
}

func (q *SkillQueries) CreateSkill(b *models.Skill) (uuid.UUID, error) {
	query := `
        INSERT INTO skills (user_id, name)
		VALUES ($1, $2)
		RETURNING id
    `
	var id uuid.UUID
	err := q.Pool.QueryRow(context.Background(), query,
		&b.UserId,
		&b.Name,
	).Scan(&id)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("error insert skills: %w", err)
	}

	return id, nil
}

func (q *SkillQueries) UpdateSkill(id uuid.UUID, b *models.Skill) error {
	query := `
        UPDATE skills
        SET  user_id = $2, name = $3
        WHERE id = $1
    `

	commandTag, err := q.Pool.Exec(context.Background(), query, id, b.UserId, b.Name)
	if err != nil {
		return fmt.Errorf("error updating skill: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no skill found with id %d", id)
	}

	return nil
}

func (q *SkillQueries) DeleteSkill(id uuid.UUID) error {
	query := `
        DELETE FROM skills
        WHERE id = $1
    `
	commandTag, err := q.Pool.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("error deleting skill: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no skill found with id %d", id)
	}

	return nil
}
