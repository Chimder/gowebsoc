-- name: CreateChannel :one
INSERT INTO channels (name, created_at, updated_at)
VALUES ($1, NOW(), NOW())
RETURNING (id, name, created_at, updated_at);

-- name: GetChannels :many
SELECT * FROM channels;

-- name: GetChannel :one
SELECT * FROM channels WHERE id = $1;