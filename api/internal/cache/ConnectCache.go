package cache

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	REDIS_HOST string
	REDIS_PASS string
	REDIS_DB   int
}

func createRedisConfig() (*RedisConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Erro ao carregar variáveis de ambiente")
	}
	config := &RedisConfig{}
	host := os.Getenv("REDIS_HOST")
	pass := os.Getenv("REDIS_PASS")
	db := os.Getenv("REDIS_DB")
	if host == "" || pass == "" || db == "" {
		return nil, fmt.Errorf("Variáveis sobre o Redis não encontradas!")
	}

	intDB, err := strconv.Atoi(db)
	if err != nil {
		return nil, err
	}

	config.REDIS_HOST = host
	config.REDIS_PASS = pass
	config.REDIS_DB = intDB

	return config, nil
}

func CreateRedisClient() (*MainCache, error) {
	config, err := createRedisConfig()
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_HOST,
		Password: config.REDIS_PASS,
		DB:       config.REDIS_DB,
		PoolSize: 10,
	})

	response := &MainCache{
		conn: rdb,
	}

	return response, nil
}
