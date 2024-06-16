package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"

	// _ "goSql/bun/migrations"

	"github.com/uptrace/bun/migrate"
)

func main() {
	// Open a PostgreSQL database.
	dsn := "postgres://postgres.donkqrnmvlejhyhlcddu:7DsShp1,l8Gx@aws-0-eu-central-1.pooler.supabase.com:6543/postgres"
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Create a Bun db on top of it.
	db := bun.NewDB(pgdb, pgdialect.New())

	// Print all queries to stdout.
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	// Create a new migrator.
	migrator := migrate.NewMigrator(db)

	// Run the migrations.
	ctx := context.Background()
	if err := migrator.Init(ctx); err != nil {
		log.Fatalf("could not initialize migrator: %v", err)
	}
	if _, err := migrator.Migrate(ctx); err != nil {
		log.Fatalf("could not run migrations: %v", err)
	}

	fmt.Println("Migrations successfully applied")
}
