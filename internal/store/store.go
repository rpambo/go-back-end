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
	}
	Users interface {
		Create(context.Context) error
	}
}

func NewStorage(db *sql.DB) Storage{
	return Storage{
		Posts: &PostsStore{db},
		Users: &UsersStore{db},
	}
}