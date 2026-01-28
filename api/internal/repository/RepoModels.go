package repository

import (
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDb struct {
	pool  *pgxpool.Pool
	mutex sync.RWMutex
	dsn   string
}
