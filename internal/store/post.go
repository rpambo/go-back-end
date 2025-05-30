package store

import (
	"context"
	"database/sql"
	"errors"

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

func (s *PostsStore) GetById(ctx context.Context, id int64) (*types.Post, error){
	query := `
			SELECT id, user_id, content, title, create_at, update_at, tags, version
			FROM posts
			WHERE id = $1
			`
	
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var Posts types.Post
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&Posts.ID,
		&Posts.UserID,
		&Posts.Title,
		&Posts.CreatedAt,
		&Posts.UpdatedAt,
		pq.Array(&Posts.Tags),
		&Posts.Version,
	)

	if err != nil{
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &Posts, nil
}