package tokenjwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type JwtInput struct {
	ID      string
	ApiURL  string
	Secret  []byte
	Timeout int
	Retries int
	Expires time.Time
}

func (r *JwtInput) getSecret() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	secret := os.Getenv("SECRET_KEY")
	r.Secret = []byte(secret)
	return nil
}

func SetSecretJwtInput(Gateway []*JwtInput) ([]*JwtInput, error) {
	for i, item := range Gateway {
		if item == nil {
			return nil, fmt.Errorf("Gateway[%d] é nil", i)
		}

		err := item.getSecret()
		if err != nil {
			return nil, err
		}
	}

	return Gateway, nil
}

type claims struct {
	ID      string `json:"api_id"`
	ApiURL  string `json:"api_url"`
	ApiKey  string `json:"api_key"`
	Timeout int    `json:"timeout"`
	Retries int    `json:"retries"`
	jwt.RegisteredClaims
}
