package models

import ()

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
	Api_Key    string `yaml:"api_key" validate:"required"`
}

type RetriesConfig struct {
	Timeout int16 `yaml:"timeout"`
	Retries int16 `yaml:"retries"`
}
