package handler

import (
	"goSql/internal/queries"
	"goSql/utils"
	"net/http"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type ChannelH struct {
	sqlc *queries.Queries
	rdb  *redis.Client
}

func NewChannel(sqlc *queries.Queries, rdb *redis.Client) *ChannelH {
	return &ChannelH{sqlc: sqlc, rdb: rdb}
}

// @Summary		Create Channel
// @Description	Create Channel
// @Tags			Channel
// @ID				create-channel
// @Accept			json
// @Produce		json
// @Param			name	query		string	false	"Name of Channel"
// @Success		200		{object}	queries.Channel
// @Router			/channel/create [post]
func (c *ChannelH) CreateChannel(w http.ResponseWriter, r *http.Request) {
	op := "handler Create Channel"
	name := r.URL.Query().Get("name")

	channel, err := c.sqlc.CreateChannel(r.Context(), name)
	if err != nil {
		utils.WriteError(w, 500, op+"CC", err)
		return
	}

	utils.WriteJSON(w, 200, channel)
	return

}

// @Summary		Get channels
// @Description	Get channels
// @Tags			Channel
// @ID				get-channels
// @Accept			json
// @Produce		json
// @Success		200	{array}	queries.Channel
// @Router			/channels [get]
func (c *ChannelH) GetChannels(w http.ResponseWriter, r *http.Request) {
	op := "handler GetChannels"

	channels, err := c.sqlc.GetChannels(r.Context())
	if err != nil {
		utils.WriteError(w, 500, op+"GC", err)
		return
	}

	utils.WriteJSON(w, 200, channels)
	return

}

type ChannelWithPodchannels struct {
	Channel     queries.Channel      `json:"channel"`
	Podchannels []queries.Podchannel `json:"podchannels"`
}

// func (c *ChannelH) GetChannels(w http.ResponseWriter, r *Request) {
// 	op := "handler GetChannels"

// 	channels, err := c.sqlc.GetChannels(r.Context())
// 	if err != nil {
// 		utils.WriteError(w, 500, op+"GC", err)
// 		return
// 	}

// 	type ChannelWithPodchannels struct {
// 		Channel    queries.Channel     `json:"channel"`
// 		Podchannels []queries.Podchannel `json:"podchannels"`
// 	}

// 	var result []ChannelWithPodchannels

// 	for _, channel := range channels {
// 		podchannels, err := c.sqlc.GetPodchannels(r.Context(), channel.ID)
// 		if err != nil {
// 			utils.WriteError(w, 500, op+"GPC", err)
// 			return
// 		}

// 		channelWithPodchannels := ChannelWithPodchannels{
// 			Channel:    channel,
// 			Podchannels: podchannels,
// 		}

// 		result = append(result, channelWithPodchannels)
// 	}

// 	if err := utils.WriteJSON(w, 200, result); err != nil {
// 		utils.WriteError(w, 500, op+"WJ", err)
// 		return
// 	}
// }

// @Summary		Get one channel
// @Description	Get one channel
// @Tags			Channel
// @ID				get-channel
// @Accept			json
// @Produce		json
// @Param			id	query		string	true	"ID of the channel"
// @Success		200			{object}	ChannelWithPodchannels
// @Router			/channel [get]
func (c *ChannelH) GetChannel(w http.ResponseWriter, r *http.Request) {
	op := "handler GetChannel"
	channelID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		utils.WriteError(w, 500, op+"Atoi", err)
		return
	}

	channel, err := c.sqlc.GetChannel(r.Context(), int32(channelID))
	if err != nil {
		utils.WriteError(w, 500, "Get channel err:", err)
		return
	}

	podchannels, err := c.sqlc.GetPodchannels(r.Context(), channel.ID)
	if err != nil {
		utils.WriteError(w, 500, "Get podchannels err:", err)
		return
	}

	channelWithPod := ChannelWithPodchannels{Channel: channel, Podchannels: podchannels}

	utils.WriteJSON(w, 200, channelWithPod)

}
