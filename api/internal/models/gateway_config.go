package models

import (
	"time"

	"github.com/Ivan-Martins-DevProjects/PayHub/internal/repository"
)

type Config struct {
	Gateways map[string]Gateway `yaml:",inline" validate:"required"`
}

type Gateway struct {
	Info    InfoConfig    `yaml:"info"`
	Secrets SecretsConfig `yaml:"secrets"`
	Retries RetriesConfig `yaml:"retries"`
}

type InfoConfig struct {
	Api_URL string `yaml:"api_url" validate:"required"`
}

type SecretsConfig struct {
	Api_Key string `yaml:"api_key" validate:"required"`
}

type RetriesConfig struct {
	Timeout int16 `yaml:"timeout"`
	Retries int16 `yaml:"retries"`
}

func (g *Config) GetInputDB() (*repository.InputDBGateway, error) {
	var response *repository.InputDBGateway
	for name, gateways := range g.Gateways {
		response = &repository.InputDBGateway{
			ID:        name,
			Api_URL:   gateways.Info.Api_URL,
			Api_Key:   gateways.Secrets.Api_Key,
			Timeout:   gateways.Retries.Timeout,
			Retries:   gateways.Retries.Retries,
			CreatedAt: time.Now(),
			ExpireAt:  time.Now().Add(168 * time.Hour),
		}
	}
	return response, nil
}
