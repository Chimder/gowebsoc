package db

import (
	"goSql/internal/config"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

// DBConn establishes a connection to the database

func DBConn() (*sqlx.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// dbURL := os.Getenv("DB_URL")
	// if dbURL == "" {
	// 	dbURL = config.LoadEnv().DB_URL
	// }

	db, err := sqlx.Connect("postgres", config.LoadEnv().DB_URL)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
		return nil, err
	}
	return db, nil
}
