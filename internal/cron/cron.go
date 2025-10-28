package cron

import (
	"NextShortLink/internal/logger"

	"github.com/robfig/cron/v3"
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
	_, err := C.AddFunc("@every 30s", deleteExpired)
	if err != nil {
		logger.L.Fatal(err.Error())
	}
}
