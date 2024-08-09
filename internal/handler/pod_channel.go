package handler

import (
	"goSql/internal/queries"
	"goSql/utils"
	"net/http"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type PodChannelH struct {
	sqlc *queries.Queries
	rdb  *redis.Client
}

func NewPodChannel(sqlc *queries.Queries, rdb *redis.Client) *PodChannelH {
	return &PodChannelH{sqlc: sqlc, rdb: rdb}
}

// @Summary		Get podchannels
// @Description	Get podchannels
// @Tags			PodChannel
// @ID				get-podchannels
// @Accept			json
// @Produce		json
// @Param			channelId	query		int	true	"ID of the podchannel"
// @Success		200				{array}	queries.Podchannel
// @Router			/podchannels [get]
func (p *PodChannelH) GetPodchannels(w http.ResponseWriter, r *http.Request) {
	op := "handler GetPodchannels"
	podchannelID, err := strconv.Atoi(r.URL.Query().Get("channelId"))
	if err != nil {
		utils.WriteError(w, 500, op+"Atoi", err)
		return
	}

	podchannel, err := p.sqlc.GetPodchannels(r.Context(), int32(podchannelID))
	if err != nil {
		utils.WriteError(w, 500, op+"GP", err)
		return
	}

	if err := utils.WriteJSON(w, 200, podchannel); err != nil {
		utils.WriteError(w, 500, op+"WJ", err)
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
// @Param			types		query		string	true	"type of the podchannel"
// @Param			id	query		int		true	"channel of the podchannel"
// @Success		200			{object}	ChannelWithPodchannels
// @Router			/podchannel/create [post]
func (p *PodChannelH) CreatePodchannel(w http.ResponseWriter, r *http.Request) {
	op := "heandler CreatePodchannel"
	name := r.URL.Query().Get("name")
	types := r.URL.Query().Get("types")
	channelID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		utils.WriteError(w, 500, op+"Atoi", err)
		return
	}

	podChannels, err := p.sqlc.CreatePodChannel(r.Context(), queries.CreatePodChannelParams{
		Name: name, Types: types, ChannelID: int32(channelID)})
	if err != nil {
		utils.WriteError(w, 500, op+"CPC", err)
		return
	}

	if err := utils.WriteJSON(w, 200, podChannels); err != nil {
		utils.WriteError(w, 500, op+"WJ", err)
		return
	}
}
