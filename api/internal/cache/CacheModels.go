package cache

import (
	"time"

	"github.com/Ivan-Martins-DevProjects/PayHub/internal/security"
	"github.com/Ivan-Martins-DevProjects/PayHub/internal/system/files"
)

type GatewayToRepo struct {
	Api_Key string
	Info    InfoGatewayToRepo
}

type InfoGatewayToRepo struct {
	ID      string
	Api_URL string
	Timeout int16
	Retries int16
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

func (g *GatewayToRepo) DecodeApiKey(secret string) error {
	EncryptKey, err := files.ReadFile(".keys", g.Info.ID)
	if err != nil {
		return err
	}

	DecryptKey, err := security.DecryptKey(EncryptKey, secret)
	if err != nil {
		return err
	}

	g.Api_Key = DecryptKey
	return nil
}

func (g *GatewayToRepo) GetInputDB(secret string) *InputDBGateway {
	err := g.DecodeApiKey(secret)
	if err != nil {
		return nil
	}

	response := &InputDBGateway{
		ID:        g.Info.ID,
		Api_URL:   g.Info.Api_URL,
		Timeout:   g.Info.Timeout,
		Retries:   g.Info.Retries,
		CreatedAt: time.Now(),
		ExpireAt:  time.Now().Add(168 * time.Hour),
	}
	return response
}
