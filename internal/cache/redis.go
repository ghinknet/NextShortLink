package cache

import (
	"NextShortLink/internal/config"
	"NextShortLink/internal/env"
	"NextShortLink/internal/logger"
	"context"
	"fmt"
	"strings"

	"github.com/redis/go-redis/v9"
)

var R *redis.Client

type NOOPLogger struct{}

func (NOOPLogger) Printf(ctx context.Context, format string, v ...any) {
	// Do nothing here
}

func InitRedis() {
	redis.SetLogger(NOOPLogger{})

	R = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(
			"%s:%d",
			config.Get().Redis.Host,
			config.Get().Redis.Port,
		),
		Password: config.Get().Redis.Password,
		DB:       config.Get().Redis.DB,
	})

	logger.L.Debug("Redis initialized")
}

func GenKey(keys ...string) string {
	return fmt.Sprintf(
		"%s:%s",
		env.ENName,
		strings.Join(keys, ":"),
	)
}
