package repository

import (
	"context"

	e "github.com/Ivan-Martins-DevProjects/PayHub/internal/appErrors"
	"github.com/Ivan-Martins-DevProjects/PayHub/internal/models"
)

func (db *MainRepo) InsertGatewayInfo(ctx context.Context, config []*models.Config, secret string) ([]*InputDBGateway, error) {
	GatewaysConfig, err := GetInputDB(config, secret)
	if err != nil {
		return nil, err
	}

	for _, item := range GatewaysConfig {
		exists, err := db.CheckIfGatewayExists(ctx, item.ID)
		if err != nil {
			return nil, e.GenerateError(*InternalDBError, err)
		}

		if exists == true {
			return nil, e.GenerateError(*GatewayAlreadyExists, err)
		}
	}

	err = db.InsertGatewayInfoDB(ctx, GatewaysConfig)
	if err != nil {
		return nil, e.GenerateError(*InternalDBError, err)
	}

	return GatewaysConfig, nil
}
