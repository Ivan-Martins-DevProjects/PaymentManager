package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func InsertTokenAndApiInfo(ctx context.Context, rdb *MainCache, info *CacheInput) error {
	jsonBytes, err := json.Marshal(info)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("token:%s", info.Token)

	_, err = rdb.conn.Set(ctx, key, jsonBytes, 1*time.Hour).Result()
	if err != nil {
		return err
	}

	return nil
}
