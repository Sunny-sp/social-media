package db

import (
	"context"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
)

// make connection/create connection-pool

func ConnectPool(dsn string) (pool *pgxpool.Pool, err error) {

	ctx := context.Background()

	pool, err = pgxpool.New(ctx, dsn)

	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)

	if err != nil {
		return nil, err
	}

	log.Println("Database connected successfully!")

	return pool, err
}
