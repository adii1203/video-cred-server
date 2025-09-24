-- name: CreateCred :one
INSERT INTO creds (user_id, cred_name, resume_url, video_url)
VALUES 
($1, $2, $3, $4) 
RETURNING *;