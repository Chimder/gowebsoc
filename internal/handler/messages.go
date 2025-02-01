package handler

import (
	"context"
	"encoding/json"
	"goSql/internal/queries"
	"goSql/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type MessagesH struct {
	sqlc *queries.Queries
	rdb  *redis.Client
}

func NewMessage(sqlc *queries.Queries, rdb *redis.Client) *MessagesH {
	return &MessagesH{sqlc: sqlc, rdb: rdb}
}

// @Summary		Get Messages PodChannel
// @Description	mess podchannel
// @Tags			PodChannel
// @ID				get-podchannel-message
// @Accept			json
// @Produce		json
// @Param			podchannel_id	query		int	true	"podchannel id"
// @Param			limit	query		int	true	"limit"
// @Param			page	query		int	true	"page"
// @Success		200		{array}	queries.Message
// @Router			/podchannel/message [get]
func (m *MessagesH) GetPodchannelsMessages(w http.ResponseWriter, r *http.Request) {
	op := "handler GetPodchannelsMessages"
	podchannelId, err := strconv.Atoi(r.URL.Query().Get("podchannel_id"))
	if err != nil {
		utils.WriteError(w, 500, op+"Atoi", err)
		return
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit <= 0 {
		limit = 20
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit

	messages, err := m.sqlc.GetMessages(r.Context(),
		queries.GetMessagesParams{PodchannelID: int32(podchannelId), Limit: int32(limit), Offset: int32(offset)})
	// query := `SELECT * FROM messages WHERE podchannel_id=$1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
	// query := `SELECT * FROM messages WHERE podchannel_id=$1 LIMIT $2 OFFSET $3`
	if err != nil {
		utils.WriteError(w, 500, op+"GM", err)
		return
	}

	utils.WriteJSON(w, 200, messages)
	return

}

func ProcessMessages(sqlc *queries.Queries, rdb *redis.Client) {
	ctx := context.Background()
	for {
		// time := 5 * time.Second
		msgData, err := rdb.BLPop(ctx, 0, "messageQueue").Result()
		if err != nil {
			log.Printf("Redis pop error: %s\n", err)
			continue
		}

		if len(msgData) < 2 {
			continue
		}

		var data *EventMessage
		if err := json.Unmarshal([]byte(msgData[1]), &data); err != nil {
			log.Printf("Unmarshal error: %s\n", err)
			continue
		}

		message, ok := data.Message.(string)
		if !ok {
			log.Printf("Data is not a string: %v\n", data.Message)
			continue
		}

		err = sqlc.CreateMessage(ctx, queries.CreateMessageParams{
			ID: data.MessageID, Message: message, AuthorID: data.AuthorID, PodchannelID: int32(data.PodchannelID), CreatedAt: data.CreatedAt,
		})
		if err != nil {
			log.Printf("DB insert error: %s\n", err)
		}
	}
}
