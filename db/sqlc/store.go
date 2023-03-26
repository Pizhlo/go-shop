package db

import (
	"database/sql"
)

// SQLStore provides all functions to execute db queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

// Store provides all functions to execute db queries and transactions
type Store interface {
	Querier
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
