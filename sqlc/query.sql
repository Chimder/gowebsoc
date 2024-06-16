-- name: CreateChannel :one
INSERT INTO channels (name, created_at, updated_at) VALUES ($1, NOW(), NOW())
RETURNING id, name, created_at, updated_at;

-- name: GetChannel :one
SELECT id, name, created_at, updated_at FROM channels WHERE id = $1;

-- name: CreatePodchannel :one
INSERT INTO podchannels (name, type, channel_id, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW())
RETURNING id, name, type, created_at, updated_at, channel_id;

-- name: GetPodchannels :many
SELECT id, name, type, created_at, updated_at, channel_id FROM podchannels WHERE channel_id = $1;
