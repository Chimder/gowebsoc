package server

import (
	"goSql/internal/handler"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Router struct {
	db *sqlx.DB
}

func NewRouter(db *sqlx.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))
	//////////////////////
	wsServer := handler.NewWebServer()
	wsServer.Run()

	//////////////////////

	userHandler := handler.NewUser(db)
	/////////////////////
	r.Get("/ws", wsServer.WsConnections)
	/////////////////////////
	r.HandleFunc("/yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "docs/swagger.yaml")
	})
	r.Get("/swag", httpSwagger.WrapHandler)
	//////////////////////////
	r.Post("/channel/create", userHandler.CreateChannel)
	r.Get("/channel/get", userHandler.GetChannels)
	r.Get("/channel/{channelID}", userHandler.GetChannel)
	r.Post("/podchannel/create", userHandler.CreatePodchannel)
	r.Get("/podchannel/{podchannelID}", userHandler.GetPodchannel)
	r.Get("/podchannels", userHandler.GetPodchannels)

	return r
}
