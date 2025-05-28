package store

import "database/sql"



type CommentStore struct {
	db *sql.DB
}
