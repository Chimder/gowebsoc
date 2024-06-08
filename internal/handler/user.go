package handler

import (
	"goSql/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UserH struct {
	db *sqlx.DB
}

type Channel struct {
	ID        string `db:"id" json:"id"`
	Name      string `db:"name"`
	OwnerID   string `db:"owner_id"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func NewUser(db *sqlx.DB) *UserH {
	return &UserH{db: db}
}

func (u *UserH) Create(w http.ResponseWriter, r *http.Request) {

	channel := Channel{
		ID:      uuid.NewString(),
		Name:    "Football",
		OwnerID: "lol",
	}

	query := `INSERT INTO channels (id, name, owner_id, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING *`
	err := u.db.Get(&channel, query, channel.ID, channel.Name, channel.OwnerID)
	if err != nil {
		utils.WriteError(w, 500, "Create channel err:", err)
		return
	}

}
