package cache

import (
	"context"

	"github.com/Ivan-Martins-DevProjects/PayHub/internal/repository"
	"github.com/Ivan-Martins-DevProjects/PayHub/internal/security"
)

func (c *MainCache) InsertGatewayCacheInfo(ctx context.Context, info []*repository.InputDBGateway) error {
	for _, item := range info {
		token := security.GenerateToken()
		element := &CacheInput{
			ID:     item.ID,
			Token:  token,
			ApiKey: item.Api_Key,
			ApiURL: item.Api_URL,
		}

		err := InsertTokenAndApiInfo(ctx, c, element)
		if err != nil {
			return err
		}
	}

	return nil
}
