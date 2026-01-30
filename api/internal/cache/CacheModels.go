package cache

import (
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type CacheInput struct {
	ID     string    `json:"id"`
	Token  uuid.UUID `json:"token"`
	ApiKey string    `json:"api_key"`
	ApiURL string    `json:"api_url"`
}

type MainCache struct {
	conn *redis.Client
}
