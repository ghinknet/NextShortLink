package link

import (
	"NextShortLink/internal/cache"
	"NextShortLink/internal/config"
	"NextShortLink/internal/database"
	"NextShortLink/internal/model"
	"NextShortLink/internal/repository"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ghinknet/toolbox/expr"
	"github.com/ghinknet/toolbox/pointer"
)

// AddLink adds a link
func AddLink(link string, validity *int64) (string, error) {
	// Check link
	if !(strings.HasPrefix(link, "https://") || strings.HasPrefix(link, "http://")) {
		return "", model.ErrLinkInvalid
	}

	// Check validity
	if validity != nil && *validity < time.Now().Unix() {
		return "", model.ErrValidityInvalid
	}

	// Construct link
	linkInsert := &model.DatabaseLink{
		Link:     link,
		Validity: validity,
	}

	// Get database session
	databaseSession := database.E.NewSession()
	defer database.Close(databaseSession)

	// Get repository session
	linkRepo := repository.NewLinkRepository(databaseSession)

	// Insert link
	if err := linkRepo.Insert(linkInsert); err != nil {
		return "", err
	}

	// Make reversed field
	fieldMap := make(map[int64]rune)
	for k, v := range config.Field {
		fieldMap[v] = k
	}

	// Encode linkID
	linkIDSlice := make([]rune, 0)
	remaining := linkInsert.ID

	for remaining > 0 {
		remainder := remaining % 62
		linkIDSlice = append(linkIDSlice, fieldMap[remainder])
		remaining = remaining / 62
	}

	// Reverse slice
	for i, j := 0, len(linkIDSlice)-1; i < j; i, j = i+1, j-1 {
		linkIDSlice[i], linkIDSlice[j] = linkIDSlice[j], linkIDSlice[i]
	}

	linkID := string(linkIDSlice)

	// Record redis cache
	cache.R.Set(
		context.Background(),
		fmt.Sprintf("nextShortLink:links:%s", linkID),
		link,
		expr.Ternary(validity != nil, time.Until(time.Unix(pointer.SafeDeref(validity), 0)), 1*time.Hour),
	)

	return linkID, nil
}
