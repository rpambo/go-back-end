package store

import (
	"context"
	"database/sql"

	"github.com/rpambo/go-back-end/types"
)

type CommentStore struct {
	db *sql.DB
}

func (s *CommentStore) GetPostByID(ctx context.Context, PostID int64) ([]types.Comment, error){
	query := `
				SELECT c.id, c.post_id, c.user_id, c.content, user.username, user.id FROM comments c
				JOIN users on users.id = c.user_id
				WHERE c.post = $s1
				ORBER BY c.create_at DESC;
				`
	
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, PostID)

	if err != nil{
		return nil, err
	}

	defer rows.Close()

	comments := []types.Comment{}
	for rows.Next() {
		var c types.Comment
		c.User = types.User{}
		err := rows.Scan(&c.ID, &c.PostID, &c.UserID, &c.Content, &c.CreatedAt, &c.User.Username, &c.User.ID)
		if err != nil {
			return nil, err
		}

		comments = append(comments, c)
	}

	return comments, nil
}

func (s *CommentStore) Create(ctx context.Context, comment *types.Comment) error {
	query := `
			INSERT INTO comments(pots_id, user_id, content)
			VALUES($1, $2, $3) RETURNING (id, create_at)
			`
	
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		comment.PostID,
		comment.UserID,
		comment.Content,
	).Scan(&comment.ID, comment.CreatedAt)

	if err != nil {
		return err
	}

	return nil
} 
