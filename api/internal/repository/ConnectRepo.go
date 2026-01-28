package repository

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func (db *PostgresDb) GetPool(ctx context.Context) (*pgxpool.Pool, error) {
	if db.dsn == "" {
		dsn := os.Getenv("POSTGRES_URL")

		db.dsn = dsn
	}

	db.mutex.RLock()
	if db.pool != nil {
		defer db.mutex.RUnlock()
		return db.pool, nil
	}
	db.mutex.RUnlock()

	db.mutex.Lock()
	defer db.mutex.Unlock()

	if db.pool != nil {
		return db.pool, nil
	}

	config, err := pgxpool.ParseConfig(db.dsn)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	db.pool = pool
	return db.pool, nil
}

func (db *PostgresDb) Close() {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if db.pool != nil {
		db.pool.Close()
		db.pool = nil
	}
}
