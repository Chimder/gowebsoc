package handler

import (
	"goSql/models"
	"goSql/utils"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UserH struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *UserH {
	return &UserH{db: db}
}

func (u *UserH) Create(w http.ResponseWriter, r *http.Request) {
	channel := models.Channel{Name: "football"}

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

type ChannelWithPodchannels struct {
	Channel     models.Channel      `json:"channel"`
	Podchannels []models.Podchannel `json:"podchannels"`
}

func (u *UserH) GetChannel(w http.ResponseWriter, r *http.Request) {
	var channel models.Channel
	var podchannel []models.Podchannel

	query := `SELECT * FROM "channels" WHERE id=$1`
	podquery := `SELECT * FROM "podchannels" WHERE channel_id=$1`
	err := u.db.Get(&channel, query, 2)
	if err != nil {
		utils.WriteError(w, 500, "Create podchannel err:", err)
		return
	}
	err = u.db.Select(&podchannel, podquery, channel.ID)
	if err != nil {
		utils.WriteError(w, 500, "Create podchannel err:", err)
		return
	}

	channelWithPod := ChannelWithPodchannels{Channel: channel, Podchannels: podchannel}

	if err := utils.WriteJSON(w, 200, channelWithPod); err != nil {
		utils.WriteError(w, 500, "Create podchannel write", err)
		return
	}
}

func (u *UserH) CreatePodchannel(w http.ResponseWriter, r *http.Request) {
	podchannel := models.Podchannel{Name: "Barsa", Type: "team", ChannelID: 2}

	// Выполнение запроса с возвращением id, created_at и updated_at
	query := `INSERT INTO podchannels (name, type, channel_id, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id, name, created_at, updated_at`
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
	var podchannel models.Podchannel
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
