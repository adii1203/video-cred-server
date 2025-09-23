-- name: CreateUser :one
INSERT INTO users (name, email, profile_picture, clerkId)
VALUES ($1, $2, $3, $4) RETURNING *;