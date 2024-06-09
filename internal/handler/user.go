package handler

import (
	"goSql/model"
	"goSql/utils"
	"net/http"

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
	channel := model.Channel{Name: "football"}
	query := `INSERT INTO channels (name, created_at, updated_at) VALUES ($1, NOW(), NOW()) RETURNING id, name, created_at, updated_at`
	err := u.db.Get(&channel, query, channel.Name)
	if err != nil {
		utils.WriteError(w, 500, "Create channel err:", err)
		return
	}

	if err := utils.WriteJSON(w, 200, channel); err != nil {
		utils.WriteError(w, 500, "Create channel write", err)
		return
	}
}
