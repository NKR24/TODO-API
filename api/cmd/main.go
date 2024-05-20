package main

import (
	"context"

	"github.com/nitishm/go-rejson/v4"
	"github.com/redis/go-redis/v9"
)

var (
	ctx context.Context
	rbd *redis.Client
	rh *rejson.Handler
)
