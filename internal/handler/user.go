package handler

import (
	"context"
	"goSql/internal/db"
	"goSql/utils"
	"net/http"

	_ "github.com/lib/pq"
)

type UserH struct {
	q *db.Queries
}

func NewUser(q *db.Queries) *UserH {
	return &UserH{q: q}
}

// type UserH struct {
// 	q *
// }

// func NewUser(db *sql.DB) *UserH {
// 	return &UserH{db: db}
// }

func (u *UserH) Create(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	channel, err := u.q.CreateChannel(ctx, "Games")
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
	ctx := context.Background()
	channel, err := u.q.GetChannel(ctx, 4)
	if err != nil {
		utils.WriteError(w, 500, "Get channel err:", err)
		return
	}

	podchannels, err := u.q.GetPodchannels(ctx, 4)
	if err != nil {
		utils.WriteError(w, 500, "Get podchannels err:", err)
		return
	}
	channel.PodChannels = podchannels

	if err := utils.WriteJSON(w, 200, channel); err != nil {
		utils.WriteError(w, 500, "Get channel write", err)
		return
	}
}

func (u *UserH) CreatePodchannel(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	podchannel, err := u.q.CreatePodchannel(ctx, db.CreatePodchannelParams{
		Name:      "Persna",
		Type:      "ver 5",
		ChannelID: 4,
	})
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
	ctx := context.Background()
	podchannels, err := u.q.GetPodchannels(ctx, 4)
	if err != nil {
		utils.WriteError(w, 500, "Get podchannel err:", err)
		return
	}

	if err := utils.WriteJSON(w, 200, podchannels); err != nil {
		utils.WriteError(w, 500, "Get podchannel write", err)
		return
	}
}
