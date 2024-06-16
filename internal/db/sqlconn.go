package db

import (
	"context"
	"goSql/internal/config"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// DBConn establishes a connection to the database
func DBConn() (*pgxpool.Pool, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = config.LoadEnv().DB_URL
	}

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatal("Unable to parse database URL:", err)
		return nil, err
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Unable to create database pool:", err)
		return nil, err
	}

	return dbpool, nil
}
