package db

import (
	"database/sql"
	"goSql/internal/config"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func BunConnection() (*bun.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	dbURL := os.Getenv("DB_BUN")
	if dbURL == "" {
		dbURL = config.LoadEnv().DB_URL
	}
	bundb, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
		return nil, err
	}

	db := bun.NewDB(bundb, sqlitedialect.New())

	return db, nil
}
