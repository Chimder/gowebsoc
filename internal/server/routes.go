package server

import (
	"goSql/internal/handler"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
	log.Println("Server Run")

	r.Get("/ws", wsServer.WsConnections)

	r.Get("/channel/create", userHandler.Create)
	r.Get("/channel/get", userHandler.GetChannel)
	r.Get("/podchannel/create", userHandler.CreatePodchannel)
	r.Get("/podchannel/get", userHandler.GetPodchannel)

	return r
}
