package server

import (
	"goSql/internal/db"
	"log"
	"net/http"
	"os"
	"time"
)

func NewServer() *http.Server {

	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "4000"

	}

	db, err := db.DBConnection()
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer db.Close()

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      NewRouter(db),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}

	return server
}
