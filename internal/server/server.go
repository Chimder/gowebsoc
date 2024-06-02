package server

import (
	"goSql/internal/handler"
	"net/http"
	"os"
	"time"
)

func NewServer() *http.Server {
	wsServer := handler.NewWebServer()
	wsServer.Run()

	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "4000"
	}
	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      NewRouter(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}

	return server
}
