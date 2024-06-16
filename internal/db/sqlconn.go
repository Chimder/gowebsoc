package db

import (
	"database/sql"
	"goSql/internal/config"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// DBConn establishes a connection to the database
func DBConn() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	db, err := sql.Open("postgres", config.LoadEnv().DB_URL)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Unable to ping the database:", err)
		return nil, err
	}

	return db, nil
}
