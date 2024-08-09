package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	_ "goSql/docs"
	"goSql/internal/db"
	"goSql/internal/queries"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	httpServer *http.Server
	sqlc       *pgxpool.Pool
	rdb        *redis.Client
}

func NewServer() *Server {
	ctx := context.Background()

	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "4000"
	}

	sqlc, err := db.DBConn(ctx)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	sqlcQueries := queries.New(sqlc)

	opt, err := db.RedisCon()
	if err != nil {
		log.Fatal("Unable to connect to redis:", err)
	}

	rdb := redis.NewClient(opt)

	httpServer := &http.Server{
		Addr:         ":" + PORT,
		Handler:      NewRouter(sqlcQueries, rdb),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}

	return &Server{
		httpServer: httpServer,
		sqlc:       sqlc,
		rdb:        rdb,
	}
}

func (s *Server) Close() {
	if s.sqlc != nil {
		s.sqlc.Close()
	}
	if s.rdb != nil {
		s.rdb.Close()
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
