-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: AddUser :one
INSERT INTO users (name, password, email)
VALUES ($1, $2, $3)
RETURNING *;