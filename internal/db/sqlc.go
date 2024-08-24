package db

import (
	"context"
	"goSql/internal/config"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

func DBConn(ctx context.Context) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(config.LoadEnv().DB_URL)
	if err != nil {
		log.Fatalf("Unable to parse config: %v", err)
		return nil, err
	}
	// cfg.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
	// 	_, err := conn.Exec(ctx, "DEALLOCATE ALL")
	// 	return err
	// }
	// cfg.ConnConfig.RuntimeParams["prefer_simple_protocol"] = "1"

	sqlcPool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
		return nil, err
	}

	return sqlcPool, nil
}
