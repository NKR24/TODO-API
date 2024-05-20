package main

import (
	"context"

	"github.com/nitishm/go-rejson/v4"
	"github.com/redis/go-redis/v9"
)

var (
	ctx context.Context
	rbd *redis.Client
	rh  *rejson.Handler
)

func init() {
	ctx := context.Background()
	rbd := redis.NewClient(&redis.Options{
		DB:   0,
		Addr: "localhost:6379",
	})
	rh := rejson.NewReJSONHandler()
	rh.SetGoRedisClientWithContext(ctx, rbd)
}
