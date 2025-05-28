package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
	"github.com/rpambo/go-back-end/types"
)

type PostsStore struct {
	db *sql.DB
}

func (s *PostsStore) Create(ctx context.Context, posts *types.Post) error {
	query := `
			INSERT INTO post(content, title, user_id, tags)
			VALUES($1, $2, $3, $4) RETURNING id, creat_at, update_at
			`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		posts.Content,
		posts.Title,
		posts.UserID,
		pq.Array(posts.Tags)).Scan(&posts.ID, &posts.CreatedAt, &posts.UpdatedAt)
	
	if err != nil {
		return err
	}

	return nil
}