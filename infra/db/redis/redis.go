package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

// NewRedisClient khởi tạo kết nối Redis
func NewRedisClient(host string, port int, password string, db int) (*RedisClient, error) {
	addr := fmt.Sprintf("%s:%d", host, port)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // để trống nếu không có
		DB:       db,       // mặc định 0
	})

	// Test kết nối với timeout 5s
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return &RedisClient{Client: rdb}, nil
}
