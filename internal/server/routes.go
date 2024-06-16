package server

import (
	"goSql/internal/handler"
	"goSql/sqlc/tutorial"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type Router struct {
	// db *sql.DB
	db *pgxpool.Pool
}

func NewRouter(db *pgxpool.Pool) http.Handler {
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
