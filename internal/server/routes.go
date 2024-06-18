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

	r.Get("/ws", wsServer.WsConnections)

	r.Get("/yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "docs/swagger.yaml")
	})
	r.Mount("/swagger/", httpSwagger.WrapHandler)

	r.Get("/channel", userHandler.GetChannel)
	r.Get("/channels", userHandler.GetChannels)
	r.Post("/channel/create", userHandler.CreateChannel)

	r.Get("/podchannels", userHandler.GetPodchannels)
	r.Post("/podchannel/create", userHandler.CreatePodchannel)

	return r
}
