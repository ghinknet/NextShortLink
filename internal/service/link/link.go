package link

import (
	"NextShortLink/internal/cache"
	"NextShortLink/internal/config"
	"NextShortLink/internal/database"
	"NextShortLink/internal/repository"
	"context"
	"fmt"
	"time"
)

// intPow get pow of a int num
func intPow(base, exp int64) int64 {
	var result int64 = 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}

// GetLink reads link of linkID
func GetLink(linkID string) (string, error) {
	// TODO: Check charset to return 404

	// Check redis cache
	exists, err := cache.R.Exists(
		context.Background(),
		fmt.Sprintf("nextShortLink:links:%s", linkID),
	).Result()
	if err != nil {
		return "", err
	}
	if exists > 0 {
		link, err := cache.R.Get(
			context.Background(),
			fmt.Sprintf("nextShortLink:links:%s", linkID),
		).Result()
		if err != nil {
			return "", err
		}
		return link, nil
	}

	// Convert linkID to ID
	var linkIDConverted int64 = 0
	for i := 0; i < len(linkID); i++ {
		char := rune(linkID[len(linkID)-1-i])
		linkIDConverted += config.Field[char] * intPow(int64(len(config.Field)), int64(i))
	}

	// Get database session
	databaseSession := database.E.NewSession()
	defer database.Close(databaseSession)

	// Get repository session
	linkRepo := repository.NewLinkRepository(databaseSession)

	// Check database
	link, validity, err := linkRepo.Read(linkIDConverted)
	if err != nil {
		return "", err
	}

	// Calc ttl
	var ttl time.Duration = 0
	if validity != nil {
		ttl = time.Until(time.Unix(*validity, 0))
	}

	// Record redis cache
	cache.R.Set(
		context.Background(),
		fmt.Sprintf("nextShortLink:links:%s", linkID),
		link,
		ttl,
	)

	return link, nil
}
