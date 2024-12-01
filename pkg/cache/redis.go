package cache

import (
	"Ecommerce/config"
	"context"

	"github.com/redis/go-redis/v9"
)

func InitCache(cfg *config.Config) (*redis.Client, error) {
	// 創建 Redis 客戶端
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPass,
		DB:       0, // 使用默認 DB
	})

	// 測試連接
	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
