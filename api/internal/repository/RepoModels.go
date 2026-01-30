package repository

import (
	"sync"
	"time"

	e "github.com/Ivan-Martins-DevProjects/PayHub/internal/appErrors"
	"github.com/Ivan-Martins-DevProjects/PayHub/internal/models"
	"github.com/Ivan-Martins-DevProjects/PayHub/internal/security"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDb struct {
	pool  *pgxpool.Pool
	mutex sync.RWMutex
	dsn   string
}

type MainRepo struct {
	DB *PostgresDb
}

type InputDBGateway struct {
	ID        string
	Api_URL   string
	Api_Key   string
	Timeout   int16
	Retries   int16
	CreatedAt time.Time
	ExpireAt  time.Time
}

func GetInputDB(models []*models.Config, secret string) ([]*InputDBGateway, error) {
	var response []*InputDBGateway
	configs, err := SetSecretAndDecodeAPIKey(models, secret)
	if err != nil {
		return nil, e.GenerateError(*InternalDBError, err)
	}

	for _, config := range configs {
		for name, g := range config.Gateways {
			item := &InputDBGateway{
				ID:        name,
				Api_URL:   g.Info.Api_URL,
				Api_Key:   g.Secrets.Api_Key,
				Timeout:   g.Retries.Timeout,
				Retries:   g.Retries.Retries,
				CreatedAt: time.Now(),
				ExpireAt:  time.Now().Add(168 * time.Hour),
			}

			response = append(response, item)
		}
	}

	return response, nil
}

func SetSecretAndDecodeAPIKey(configs []*models.Config, secret string) ([]*models.Config, error) {
	for _, config := range configs {
		for name, gateway := range config.Gateways {
			EncryptKey, err := security.EncryptKey(gateway.Secrets.Api_Key, secret)
			if err != nil {
				return nil, err
			}

			config.Gateways[name].Secrets.Api_Key = EncryptKey
			config.Gateways[name].Secrets.Secret = secret
		}
	}

	return configs, nil
}
