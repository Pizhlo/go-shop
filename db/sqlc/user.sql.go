// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user.sql

package db

import (
	"context"

	"github.com/lib/pq"
)

const createUser = `-- name: CreateUser :one
INSERT INTO
    users (username, email, hashed_password, first_name, last_name)
VALUES
    ($1, $2, $3, $4, $5) RETURNING id, username, email, hashed_password, favourites, first_name, last_name
`

type CreateUserParams struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.HashedPassword,
		arg.FirstName,
		arg.LastName,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		pq.Array(&i.Favourites),
		&i.FirstName,
		&i.LastName,
	)
	return i, err
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteAuthor, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, username, email, hashed_password, favourites, first_name, last_name
FROM users
WHERE username = $1
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		pq.Array(&i.Favourites),
		&i.FirstName,
		&i.LastName,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, email, hashed_password, favourites, first_name, last_name
FROM users
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		pq.Array(&i.Favourites),
		&i.FirstName,
		&i.LastName,
	)
	return i, err
}
