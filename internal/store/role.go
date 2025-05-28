package store

import "database/sql"

type RoleStore struct {
	db *sql.DB
}