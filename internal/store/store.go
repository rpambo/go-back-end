package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/rpambo/go-back-end/types"
)

var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Posts interface {
		Create(context.Context, *types.Post) error
		GetById(context.Context, int64) (*types.Post, error)
	}
	Users interface {
		Create(context.Context) error
	}
	Comments interface {
		Create(context.Context, *types.Comment) error
		GetPostByID(context.Context, int64) ([]types.Comment, error)
	}
}

func NewStorage(db *sql.DB) Storage{
	return Storage{
		Posts: &PostsStore{db},
		Users: &UsersStore{db},
		Comments: &CommentStore{db},
	}
}