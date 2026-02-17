package main

import (
	"NextShortLink/internal/cache"
	"NextShortLink/internal/config"
	"NextShortLink/internal/cron"
	"NextShortLink/internal/database"
	"NextShortLink/internal/handler"
	"NextShortLink/internal/logger"
	"NextShortLink/internal/repository"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"xorm.io/xorm"
)

func main() {
	// Load static config
	config.LoadStatic()

	// Init logger
	logger.InitLogger()
	defer func(L *zap.Logger) {
		_ = L.Sync()
	}(logger.L)

	// Init redis
	cache.InitRedis()
	defer func(R *redis.Client) {
		_ = R.Close()
	}(cache.R)

	// Init database
	database.InitDB()
	defer func(E *xorm.Engine) {
		_ = E.Close()
	}(database.E)

	// Init dynamic config
	databaseSession := database.E.NewSession()
	defer database.Close(databaseSession)
	configRepo := repository.NewConfigRepository(databaseSession)
	if err := configRepo.Init(); err != nil {
		logger.L.Fatal("failed to init dynamic config", zap.Error(err))
	}

	// Init cron
	cron.InitCron()

	// Start HTTP server
	handler.RunHTTPServer()

	select {}
}
