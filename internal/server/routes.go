package server

import (
	"goSql/internal/handler"
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

	WebsocketHandler := handler.NewWsHandler()
	r.Get("/ws", WebsocketHandler.WsConnections)

	return r
}
