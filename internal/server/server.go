package server

import (
	"context"

	_ "goSql/docs"
	"goSql/internal/db"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	// _ "github.com/lib/pq"
)

type Server struct {
	httpServer *http.Server
	db         *sqlx.DB
}

func NewServer() *Server {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "4000"
	}

	db, err := db.DBConn()
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	httpServer := &http.Server{
		Addr:         ":" + PORT,
		Handler:      NewRouter(db),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}

	return &Server{
		httpServer: httpServer,
		db:         db,
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
