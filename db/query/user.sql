-- name: GetUser :one
SELECT *
FROM users
WHERE username = $1
LIMIT 1;

-- name: CreateUser :one
INSERT INTO
    users (username, email, hashed_password)
VALUES
    ($1, $2, $3) RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM users
WHERE id = $1;