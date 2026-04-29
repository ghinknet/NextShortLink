package main

import (
	"NextShortLink/internal/cron"
	"NextShortLink/internal/handler"
	"NextShortLink/internal/infra/cache"
	"NextShortLink/internal/infra/config"
	"NextShortLink/internal/infra/database"
	"NextShortLink/internal/infra/logger"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"xorm.io/xorm"
)

func main() {
	// Load static config
	config.Load()

	// Init logger
	logger.InitLogger()
	defer func(L *zap.Logger) { _ = L.Sync() }(logger.L)

	// Init redis
	cache.InitRedis()
	defer func(R *redis.Client) { _ = R.Close() }(cache.R)

	// Init database
	database.InitDB()
	defer func(E *xorm.Engine) { _ = E.Close() }(database.E)

	// Init cron
	cron.InitCron()

	// Start HTTP server
	handler.RunHTTPServer()

	select {}
}
