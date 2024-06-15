package main

import (
	"context"
	"fmt"
	"goSql/model"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func init() {
	migrate.NewMigrations().MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")
		_, err := db.NewCreateTable().Model((*model.Channel)(nil)).Exec(ctx)
		if err != nil {
			return fmt.Errorf("error creating table: %w", err)
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] ")
		_, err := db.NewDropTable().Model((*model.Channel)(nil)).Exec(ctx)
		if err != nil {
			return fmt.Errorf("error dropping table: %w", err)
		}
		return nil
	})
}
