package cron

import (
	"NextShortLink/internal/infra/logger"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

var C *cron.Cron

// InitCron inits global cron object
func InitCron() {
	C = cron.New()

	registerDefault()

	C.Start()
}

// registerDefault registers default cron tasks
func registerDefault() {
	if _, err := C.AddFunc("@every 30s", deleteExpired); err != nil {
		logger.L.Fatal("failed to register default cron 'deleteExpired'", zap.Error(err))
	}
}
