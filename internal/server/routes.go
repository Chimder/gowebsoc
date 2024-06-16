package server

import (
	"database/sql"
	"goSql/internal/handler"
	"goSql/sql/tutorial"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

type Router struct {
	db *sql.DB
}

func NewRouter(db *sql.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))
	//////////////////////
	wsServer := handler.NewWebServer()
	wsServer.Run()

	//////////////////////

	query := tutorial.New(db)
	userHandler := handler.NewUser(query)
	/////////////////////
	log.Println("Server Run")

	r.Get("/ws", wsServer.WsConnections)

	r.Get("/channel/create", userHandler.Create)
	r.Get("/channel/get", userHandler.GetChannel)
	r.Get("/podchannel/create", userHandler.CreatePodchannel)
	r.Get("/podchannel/get", userHandler.GetPodchannel)

	return r
}
