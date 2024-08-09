-- name: GetPodchannels :many
SELECT * FROM podchannels WHERE channel_id = $1;

-- name: CreatePodChannel :one
INSERT INTO podchannels (name, types, channel_id, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW())
RETURNING (id, name, created_at, updated_at);