package server

import (
	"context"
	"database/sql"
	"goSql/internal/db"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	httpServer *http.Server
	db         *sql.DB
}

func NewServer() *Server {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "4000"
	}

	database, err := db.DBConn()
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer database.Close()

	httpServer := &http.Server{
		Addr:         ":" + PORT,
		Handler:      NewRouter(database),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}

	return &Server{
		httpServer: httpServer,
		db:         database,
	}
}

func (s *Server) Close() {
	if s.db != nil {
		s.db.Close()
	}
}

func (s *Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) Addr() string {
	return s.httpServer.Addr
}
