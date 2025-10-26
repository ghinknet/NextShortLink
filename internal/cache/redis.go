package cache

import (
	"NextShortLink/internal/config"
	"NextShortLink/internal/logger"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var R *redis.Client

type NOOPLogger struct{}

func (NOOPLogger) Printf(ctx context.Context, format string, v ...interface{}) {
	// Do nothing here
}

func InitRedis() {
	redis.SetLogger(NOOPLogger{})

	R = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(
			"%s:%d",
			config.C.GetString("redis.host"),
			config.C.GetInt("redis.port"),
		),
		Password: config.C.GetString("redis.password"),
		DB:       config.C.GetInt("redis.db"),
	})

	logger.L.Debug("Redis initialized")
}
