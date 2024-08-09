// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: messages.sql

package queries

import (
	"context"
)

const createMessage = `-- name: CreateMessage :exec
INSERT INTO messages (content, author_id, podchannel_id, created_at)
VALUES ($1, $2, $3, NOW())
`

type CreateMessageParams struct {
	Content      string `db:"content" json:"content"`
	AuthorID     string `db:"author_id" json:"author_id"`
	PodchannelID int32  `db:"podchannel_id" json:"podchannel_id"`
}

func (q *Queries) CreateMessage(ctx context.Context, arg CreateMessageParams) error {
	_, err := q.db.Exec(ctx, createMessage, arg.Content, arg.AuthorID, arg.PodchannelID)
	return err
}

const getMessages = `-- name: GetMessages :many
SELECT id, created_at, updated_at, content, author_id, podchannel_id FROM messages
WHERE podchannel_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetMessagesParams struct {
	PodchannelID int32 `db:"podchannel_id" json:"podchannel_id"`
	Limit        int32 `db:"limit" json:"limit"`
	Offset       int32 `db:"offset" json:"offset"`
}

func (q *Queries) GetMessages(ctx context.Context, arg GetMessagesParams) ([]Message, error) {
	rows, err := q.db.Query(ctx, getMessages, arg.PodchannelID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Content,
			&i.AuthorID,
			&i.PodchannelID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}