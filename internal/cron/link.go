package cron

import (
	"NextShortLink/internal/database"
	"NextShortLink/internal/logger"
	"NextShortLink/internal/repository"

	"go.uber.org/zap"
)

// deleteExpired deletes expired links from database
func deleteExpired() {
	// Get database session
	databaseSession := database.E.NewSession()
	defer database.Close(databaseSession)

	// Get repository session
	linkRepo := repository.NewLinkRepository(databaseSession)

	// Delete expired link
	err := linkRepo.DeleteExpired()
	if err != nil {
		logger.L.Error("Delete expired links failed", zap.Error(err))
	}
}
