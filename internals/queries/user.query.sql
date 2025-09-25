-- name: CreateUser :one
INSERT INTO users (name, email, profile_picture, clerkId)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByClerkId :one
SELECT * FROM users WHERE clerkId = $1;