package handler

import (
	"goSql/models"
	"goSql/utils"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UserH struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *UserH {
	return &UserH{db: db}
}

type Userr struct {
	db string
}

// @Summary      List accounts
// @Description  get accounts
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200  {array}   model.User
// @Router       /accounts [get]
func (u *UserH) CreateChannel(w http.ResponseWriter, r *http.Request) {
	channel := models.Channel{Name: "football"}

	query := `INSERT INTO channels (name, created_at, updated_at)
	 VALUES ($1, NOW(), NOW()) RETURNING id, name, created_at, updated_at`
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

func (u *UserH) GetChannels(w http.ResponseWriter, r *http.Request) {
	var channels []models.Channel

	query := `SELECT * FROM channels`
	err := u.db.Select(&channels, query)
	if err != nil {
		utils.WriteError(w, 500, "Get channels err:", err)
		return
	}

	if err := utils.WriteJSON(w, 200, channels); err != nil {
		utils.WriteError(w, 500, "Get channels write", err)
		return
	}
}

type ChannelWithPodchannels struct {
	Channel     models.Channel      `json:"channel"`
	Podchannels []models.Podchannel `json:"podchannels"`
}

func (u *UserH) GetChannel(w http.ResponseWriter, r *http.Request) {
	channelID := chi.URLParam(r, "channelID")

	var channel models.Channel
	var podchannels []models.Podchannel

	query := `SELECT * FROM channels WHERE id=$1`
	err := u.db.Get(&channel, query, channelID)
	if err != nil {
		utils.WriteError(w, 500, "Get channel err:", err)
		return
	}

	podquery := `SELECT * FROM podchannels WHERE channel_id=$1`
	err = u.db.Select(&podchannels, podquery, channel.ID)
	if err != nil {
		utils.WriteError(w, 500, "Get podchannels err:", err)
		return
	}

	channelWithPod := ChannelWithPodchannels{Channel: channel, Podchannels: podchannels}

	if err := utils.WriteJSON(w, 200, channelWithPod); err != nil {
		utils.WriteError(w, 500, "Get channel write", err)
		return
	}
}

func (u *UserH) CreatePodchannel(w http.ResponseWriter, r *http.Request) {
	podchannel := models.Podchannel{
		Name:      "",
		Type:      "",
		ChannelID: 0,
	}

	query := `INSERT INTO podchannels (name, type, channel_id, created_at, updated_at)
	 VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id, name, created_at, updated_at`
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
	podchannelID := chi.URLParam(r, "podchannelID")

	var podchannel models.Podchannel
	query := `SELECT * FROM podchannels WHERE id=$1`
	err := u.db.Get(&podchannel, query, podchannelID)

	if err != nil {
		utils.WriteError(w, 500, "Get podchannel err:", err)
		return
	}

	if err := utils.WriteJSON(w, 200, podchannel); err != nil {
		utils.WriteError(w, 500, "Get podchannel write", err)
		return
	}
}

func (u *UserH) GetPodchannels(w http.ResponseWriter, r *http.Request) {
	var podchannels []models.Podchannel
	query := `SELECT * FROM podchannels`
	err := u.db.Select(&podchannels, query)

	if err != nil {
		utils.WriteError(w, 500, "Get podchannels err:", err)
		return
	}

	if err := utils.WriteJSON(w, 200, podchannels); err != nil {
		utils.WriteError(w, 500, "Get podchannels write", err)
		return
	}
}
