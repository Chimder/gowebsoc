package db

import (
	"log"
	"sync"

	"github.com/jmoiron/sqlx"
)

type Queries struct {
	GetPodchannelMessagesStmt *sqlx.Stmt
}

var (
	queries      *Queries
	queriesMutex sync.Mutex
)

func PrepareQueries(pgdb *sqlx.DB) *Queries {
	queriesMutex.Lock()
	defer queriesMutex.Unlock()

	if queries != nil {
		queries.Close()
	}

	queries = &Queries{}

	var err error
	queries.GetPodchannelMessagesStmt, err = pgdb.Preparex(`SELECT * FROM messages WHERE podchannel_id=$1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`)
	if err != nil {
		log.Fatalf("Error preparing GetPodchannelMessagesStmt: %v", err)
	}

	return queries
}

func (q *Queries) Close() {
	if q.GetPodchannelMessagesStmt != nil {
		q.GetPodchannelMessagesStmt.Close()
		q.GetPodchannelMessagesStmt = nil
	}
}
