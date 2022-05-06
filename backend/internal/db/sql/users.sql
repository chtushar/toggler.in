-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: AddUser :one
INSERT INTO users (name, password, email)
VALUES ($1, $2, $3)
RETURNING *;

-- name: VerifyEmail :one
UPDATE users
SET email_verified = true
WHERE id = $1
RETURNING *;