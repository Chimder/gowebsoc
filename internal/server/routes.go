package server

import (
	"goSql/internal/handler"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

type Router struct {
}

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))
	///
	wsServer := handler.NewWebServer()
	wsServer.Run()
	///
	log.Println("Server Run")
	r.Get("/ws", wsServer.WsConnections)

	return r
}
