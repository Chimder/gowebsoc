-- name: GetMessages :many
SELECT * FROM messages
WHERE podchannel_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: CreateMessage :exec
INSERT INTO messages (content, author_id, podchannel_id, created_at)
VALUES ($1, $2, $3, NOW());