package queries

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/krazy-code/devlink/models"
)

type FollowQueries struct {
	Pool *pgxpool.Pool
}

// Create a follow
func (q *FollowQueries) CreateFollow(followerID, followedID uuid.UUID) error {
	_, err := q.Pool.Exec(context.Background(),
		`INSERT INTO follows (follower_id, followed_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`,
		followerID, followedID,
	)
	return err
}

// Delete a follow
func (q *FollowQueries) DeleteFollow(followerID, followedID uuid.UUID) error {
	_, err := q.Pool.Exec(context.Background(),
		`DELETE FROM follows WHERE follower_id=$1 AND followed_id=$2`,
		followerID, followedID,
	)
	return err
}

// List developers followed by a user
func (q *FollowQueries) ListFollowed(followerID uuid.UUID) ([]models.Follow, error) {
	rows, err := q.Pool.Query(context.Background(),
		`SELECT id, follower_id, followed_id, created_at::text FROM follows WHERE follower_id=$1`,
		followerID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var follows []models.Follow
	for rows.Next() {
		var f models.Follow
		if err := rows.Scan(&f.ID, &f.FollowerID, &f.FollowedID, &f.CreatedAt); err != nil {
			return nil, err
		}
		follows = append(follows, f)
	}
	return follows, nil
}
