package handler

import (
	"context"
	"encoding/json"
	"goSql/models"
	"goSql/utils"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

type UserH struct {
	pgdb *sqlx.DB
	rdb  *redis.Client
}

func NewUser(pgdb *sqlx.DB, rdb *redis.Client) *UserH {
	return &UserH{pgdb: pgdb, rdb: rdb}
}
func (u *UserH) ProcessMessages() {
	ctx := context.Background()
	for {
		msgData, err := u.rdb.BLPop(ctx, 0, "messageQueue").Result()
		if err != nil {
			log.Printf("Redis pop error: %s\n", err)
			continue
		}

		if len(msgData) < 2 {
			continue
		}

		var message models.Message
		if err := json.Unmarshal([]byte(msgData[1]), &message); err != nil {
			log.Printf("Unmarshal error: %s\n", err)
			continue
		}
		log.Println("REDIS<>ESS", message)

		_, err = u.pgdb.ExecContext(ctx, "INSERT INTO messages (content, author_id, podchannel_id, created_at) VALUES ($1, $2, $3, NOW())",
			message.Content, message.AuthorID, message.PodchannelID)
		if err != nil {
			log.Printf("DB insert error: %s\n", err)
		}
	}
}

// @Summary		Create Channel
// @Description	Create Channel
// @Tags			Channel
// @ID				create-channel
// @Accept			json
// @Produce		json
// @Param			name	query		string	false	"Name of Channel"
// @Success		200		{object}	models.Channel
// @Router			/channel/create [post]
func (u *UserH) CreateChannel(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	channel := models.Channel{}

	query := `INSERT INTO channels (name, created_at, updated_at)
	 VALUES ($1, NOW(), NOW()) RETURNING id, name, created_at, updated_at`
	err := u.pgdb.Get(&channel, query, name)
	if err != nil {
		utils.WriteError(w, 500, "Create channel err:", err)
		return
	}

	if err := utils.WriteJSON(w, 200, channel); err != nil {
		utils.WriteError(w, 500, "Create channel write", err)
		return
	}
}

// @Summary		Get channels
// @Description	Get channels
// @Tags			Channel
// @ID				get-channels
// @Accept			json
// @Produce		json
// @Success		200	{array}	models.Channel
// @Router			/channels [get]
func (u *UserH) GetChannels(w http.ResponseWriter, r *http.Request) {
	var channels []models.Channel

	query := `SELECT * FROM channels`
	err := u.pgdb.Select(&channels, query)
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

// @Summary		Get one channel
// @Description	Get one channel
// @Tags			Channel
// @ID				get-channel
// @Accept			json
// @Produce		json
// @Param			id	query		string	true	"ID of the channel"
// @Success		200			{object}	ChannelWithPodchannels
// @Router			/channel [get]
func (u *UserH) GetChannel(w http.ResponseWriter, r *http.Request) {
	channelID := r.URL.Query().Get("id")

	var channel models.Channel
	var podchannels []models.Podchannel

	query := `SELECT * FROM channels WHERE id=$1`
	err := u.pgdb.Get(&channel, query, channelID)
	if err != nil {
		utils.WriteError(w, 500, "Get channel err:", err)
		return
	}

	podquery := `SELECT * FROM podchannels WHERE channel_id=$1`
	err = u.pgdb.Select(&podchannels, podquery, channel.ID)
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

// @Summary		Get podchannel
// @Description	Get podchannel
// @Tags			PodChannel
// @ID				get-podchannel
// @Accept			json
// @Produce		json
// @Param			channelId	query		int	true	"ID of the podchannel"
// @Success		200				{array}	models.Podchannel
// @Router			/podchannels [get]
func (u *UserH) GetPodchannels(w http.ResponseWriter, r *http.Request) {
	podchannelID := r.URL.Query().Get("channelId")

	var podchannel []models.Podchannel
	query := `SELECT * FROM podchannels WHERE channel_id=$1`
	err := u.pgdb.Select(&podchannel, query, podchannelID)

	if err != nil {
		utils.WriteError(w, 500, "Get podchannel err:", err)
		return
	}

	if err := utils.WriteJSON(w, 200, podchannel); err != nil {
		utils.WriteError(w, 500, "Get podchannel write", err)
		return
	}
}

// @Summary		Create one podchannel
// @Description	Create one podchannel
// @Tags			PodChannel
// @ID				create-podchannel
// @Accept			json
// @Produce		json
// @Param			name		query		string	true	"Name of the podchannel"
// @Param			type		query		string	true	"type of the podchannel"
// @Param			id	query		int		true	"channel of the podchannel"
// @Success		200			{object}	ChannelWithPodchannels
// @Router			/podchannel/create [post]
func (u *UserH) CreatePodchannel(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	types := r.URL.Query().Get("types")
	channelID := r.URL.Query().Get("id")
	var podchannel models.Podchannel

	query := `INSERT INTO podchannels (name, types, channel_id, created_at, updated_at)
	 VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id, name, created_at, updated_at`
	err := u.pgdb.Get(&podchannel, query, name, types, channelID)
	if err != nil {
		utils.WriteError(w, 500, "Create podchannel err:", err)
		return
	}

	if err := utils.WriteJSON(w, 200, podchannel); err != nil {
		utils.WriteError(w, 500, "Create podchannel write", err)
		return
	}
}
