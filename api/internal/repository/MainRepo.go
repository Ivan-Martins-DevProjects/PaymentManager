package repository

import (
	"context"
	"fmt"

	"github.com/Ivan-Martins-DevProjects/PayHub/internal/models"
)

func InitRepo(config []*models.Config, secret string) error {
	GatewaysConfig, err := GetInputDB(config, secret)
	if err != nil {
		return err
	}

	ctx := context.Background()

	dbInstance := PostgresDb{}
	repo := MainRepo{
		db: &dbInstance,
	}

	for _, item := range GatewaysConfig {
		exists, err := repo.CheckIfGatewayExists(ctx, item.ID)
		if err != nil {
			return err
		}

		if exists == true {
			return fmt.Errorf("Gateway já cadastrado: %v", item.ID)
		}
	}

	err = repo.InsertGatewayInfo(ctx, GatewaysConfig)
	if err != nil {
		return err
	}

	return nil
}
