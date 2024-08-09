package server

import (
	"goSql/internal/handler"
	"goSql/internal/queries"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(sqlc *queries.Queries, rdb *redis.Client) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))

	//////////////////////

	wsServer := handler.NewWebServer(sqlc, rdb)
	wsServer.Run()

	//////////////////////
	channelHandler := handler.NewChannel(sqlc, rdb)
	podchannelHandler := handler.NewPodChannel(sqlc, rdb)
	messageHandler := handler.NewMessage(sqlc, rdb)

	r.Get("/ws", wsServer.WsConnections)

	r.Get("/yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "docs/swagger.yaml")
	})
	r.Mount("/swagger/", httpSwagger.WrapHandler)

	r.Get("/channel", channelHandler.GetChannel)
	r.Get("/channels", channelHandler.GetChannels)
	r.Post("/channel/create", channelHandler.CreateChannel)

	r.Get("/podchannels", podchannelHandler.GetPodchannels)
	r.Post("/podchannel/create", podchannelHandler.CreatePodchannel)

	r.Get("/podchannel/message", messageHandler.GetPodchannelsMessages)

	return r
}
