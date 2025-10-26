package middleware

import (
	"NextShortLink/internal/cache"
	"NextShortLink/internal/database"
	"NextShortLink/internal/model"
	"NextShortLink/internal/repository"
	"NextShortLink/pkg/web"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
)

var requestHistory = make([]model.RequestHistory, 0)

// CheckPermissionApplication provides a method to check the permission of an application
func CheckPermissionApplication(c fiber.Ctx) error {
	// Get time now
	now := time.Now().Unix()

	// Get client IP
	IP := web.ClientIP(c)
	ParsedIP := net.ParseIP(IP)

	// Get credential
	authorization := c.Get("Authorization")
	var ID, Key, Token string
	if authorization == "" {
		ID = c.Get("SecretID")
		Key = fmt.Sprintf("%x", sha256.Sum256([]byte(c.Get("SecretKey"))))
	} else if strings.HasPrefix(authorization, "Bearer ") {
		Token = strings.TrimPrefix(authorization, "Bearer ")
	} else {
		parts := strings.Split(authorization, ":")
		if len(parts) == 2 {
			ID = parts[0]
			Key = fmt.Sprintf("%x", sha256.Sum256([]byte(parts[1])))
		}
	}

	// Get database session
	databaseSession := database.E.NewSession()
	defer database.Close(databaseSession)

	// Get repository
	ApplicationRepo := repository.NewApplicationRepository(databaseSession)
	PermissionRepo := repository.NewPermissionRepository(databaseSession)
	PackageRepo := repository.NewPackageRepository(databaseSession)

	if Token != "" {
		exists, err := cache.R.Exists(
			context.Background(),
			fmt.Sprintf("nextShortLink:token:%s", Token),
		).Result()
		if err != nil {
			return model.RespInternalServerError(c, err)
		}
		if exists > 0 {
			authorization, err := cache.R.Get(
				context.Background(),
				fmt.Sprintf("nextShortLink:token:%s", Token),
			).Result()
			if err != nil {
				return model.RespInternalServerError(c, err)
			}
			parts := strings.Split(authorization, ":")
			if len(parts) == 2 {
				ID = parts[0]
				Key = parts[1]
			}
		} else {
			return model.RespPermissionDenied(c)
		}
	}

	// Save ID and Key in context
	c.Locals("SecretID", ID)
	c.Locals("SecretKey", Key)

	// Check credential
	id, err := ApplicationRepo.Get(ID, Key)
	if err != nil {
		if errors.Is(err, model.ErrApplicationNotFound) {
			return model.RespApplicationNotFound(c)
		} else {
			return model.RespInternalServerError(c, err)
		}
	}
	if id == 0 {
		return model.RespPermissionDenied(c)
	}

	// Request speed limit
	path := c.Path()
	disableKey, disableToken,
		qps, qpm,
		blacklist, whitelist,
		err := PermissionRepo.Check(id, path)
	if err != nil {
		if errors.Is(err, model.ErrPermissionDenied) {
			return model.RespPermissionDenied(c)
		} else {
			return model.RespInternalServerError(c, err)
		}
	}

	// Check authentication method limit
	if Token != "" && disableToken {
		return model.RespPermissionDenied(c)
	} else if Token == "" && disableKey {
		return model.RespPermissionDenied(c)
	}

	// Check IP blacklist and whitelist
	if whitelist != nil && len(whitelist) > 0 {
		flag := false
		for _, v := range whitelist {
			_, IPNet, err := net.ParseCIDR(v)
			if err == nil && IPNet.Contains(ParsedIP) {
				flag = true
				break
			}
		}
		if !flag {
			return model.RespPermissionDenied(c)
		}
	}
	for _, v := range blacklist {
		_, IPNet, err := net.ParseCIDR(v)
		if err == nil && IPNet.Contains(ParsedIP) {
			return model.RespPermissionDenied(c)
		}
	}

	// Check request speed limit
	var qpsCount int64 = 0
	var qpmCount int64 = 0
	requestHistoryCopy := make([]model.RequestHistory, len(requestHistory))
	copy(requestHistoryCopy, requestHistory)
	for _, history := range requestHistoryCopy {
		if history.Interface == path && history.AppID == id && now == history.Stamp {
			qpsCount++
		}
		if history.Interface == path && history.AppID == id && (now-history.Stamp) <= 60 {
			qpmCount++
		}
	}
	if (qpsCount >= qps && qps != -1) || (qpmCount >= qpm && qpm != -1) {
		return model.RespTooManyRequests(c)
	}

	// Append request history
	requestHistory = append(requestHistory, model.RequestHistory{
		Interface: path,
		AppID:     id,
		Stamp:     now,
	})

	// Start a clean task
	go func() {
		requestHistoryClean := make([]model.RequestHistory, 0)
		for _, history := range requestHistory {
			if now-history.Stamp < 5*60 {
				requestHistoryClean = append(requestHistoryClean, history)
			}
		}
		requestHistory = requestHistoryClean
	}()

	// Check package remaining
	err = PackageRepo.Take(id, path)
	if err != nil {
		if errors.Is(err, model.ErrNoPackageAvailable) {
			return model.RespNoPackageAvailable(c)
		} else {
			return model.RespInternalServerError(c, err)
		}
	}

	return c.Next()
}
