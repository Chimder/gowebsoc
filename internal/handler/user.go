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
	channel := model.Channel{Name: "Games"}
	query := `INSERT INTO channels (name, created_at, updated_at) VALUES ($1, NOW(), NOW())`
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
func (u *UserH) GetChannel(w http.ResponseWriter, r *http.Request) {
	var channel model.Channel
	var podchannels []model.Podchannel

	query := `SELECT * FROM "channels" WHERE id=$1`
	// podquery := `SELECT * FROM "podchannels" WHERE "channel_id"=$1`
	podquery := `SELECT id, name, type, created_at, updated_at, deleted_at, channel_id FROM "podchannels" WHERE channel_id=$1`

	err := u.db.Get(&channel, query, 4)
	if err != nil {
		utils.WriteError(w, 500, "Get channel err:", err)
		return
	}

	err = u.db.Select(&podchannels, podquery, 4)
	if err != nil {
		utils.WriteError(w, 500, "Get podchannels err:", err)
		return
	}
	channel.PodChannels = podchannels

	if err := utils.WriteJSON(w, 200, channel); err != nil {
		utils.WriteError(w, 500, "Create podchannel write", err)
		return
	}
}

func (u *UserH) CreatePodchannel(w http.ResponseWriter, r *http.Request) {
	podchannel := model.Podchannel{Name: "Persna", Type: "ver 5", ChannelID: 4}

	query := `INSERT INTO podchannels (name, type, channel_id, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW())`
	err := u.db.Get(&podchannel, query, podchannel.Name, podchannel.Type, podchannel.ChannelID)
	if err != nil {
		utils.WriteError(w, 500, "Create podchannel err:", err)
		return
	}

	if err := utils.WriteJSON(w, 200, podchannel); err != nil {
		utils.WriteError(w, 500, "Create podchannel write", err)
		return
	}
}

func (u *UserH) GetPodchannel(w http.ResponseWriter, r *http.Request) {
	var podchannel model.Podchannel
	query := `SELECT * FROM "podchannels"`
	err := u.db.Get(&podchannel, query)

	if err != nil {
		utils.WriteError(w, 500, "Create podchannel err:", err)
		return
	}

	if err := utils.WriteJSON(w, 200, podchannel); err != nil {
		utils.WriteError(w, 500, "Create podchannel write", err)
		return
	}
}
