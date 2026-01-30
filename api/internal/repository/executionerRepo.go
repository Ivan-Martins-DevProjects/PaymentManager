package repository

import (
	"context"
	"fmt"
	"strings"
)

func (r *MainRepo) InsertGatewayInfoDB(ctx context.Context, input []*InputDBGateway) error {
	pool, err := r.DB.GetPool(context.Background())
	if err != nil {
		return err
	}

	var sb strings.Builder
	sb.WriteString("INSERT INTO gateways (name, api_url, api_key, timeout, retries, createdAt, expiresAt) VALUES ")

	args := make([]any, 0, len(input)*7)
	for index, gateways := range input {
		p := index * 7
		fmt.Fprintf(&sb, "($%d, $%d, $%d, $%d, $%d, $%d, $%d)",
			p+1, p+2, p+3, p+4, p+5, p+6, p+7)

		if index < len(input)-1 {
			sb.WriteString(", ")
		}
		args = append(
			args,
			gateways.ID, gateways.Api_URL, gateways.Api_Key, gateways.Timeout,
			gateways.Retries, gateways.CreatedAt, gateways.ExpireAt,
		)
	}

	_, err = pool.Exec(ctx, sb.String(), args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *MainRepo) CheckIfGatewayExists(ctx context.Context, name string) (bool, error) {
	pool, err := r.DB.GetPool(ctx)
	if err != nil {
		return false, err
	}

	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM gateways WHERE name = $1)"

	err = pool.QueryRow(ctx, query, name).Scan(&exists)
	return exists, err
}
