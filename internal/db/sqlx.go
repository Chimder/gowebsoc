package db

import (
	"goSql/internal/config"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// type ctx context.Context
func DBConnection() (*sqlx.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	db, err := sqlx.Connect("postgres", config.LoadEnv().DB_URL)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
		return nil, err
	}

	return db, nil
}
