package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/Ivan-Martins-DevProjects/PayHub/internal/cache"
)

func InsertGatewayInfo(db *PostgresDb, input []*cache.InputDBGateway) error {
	pool, err := db.GetPool(context.Background())
	if err != nil {
		return err
	}

	var sb strings.Builder
	sb.WriteString("INSERT INTO gateways (name, api_url, api_key, timeout, retries, createdAt, expireAt) VALUES ")

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

	_, err = pool.Exec(context.Background(), sb.String(), args...)
	if err != nil {
		return err
	}

	return nil
}
