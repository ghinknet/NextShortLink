package link

import (
	"NextShortLink/internal/infra/cache"
	"context"
	"crypto/rand"
	"math/big"
	"strings"
	"time"
)

// IssueToken issue a token
func IssueToken(secretID string, secretKey string) (string, error) {
	var token string
	for {
		randomBytes := make([]byte, 24)
		if _, err := rand.Read(randomBytes); err != nil {
			return "", err
		}

		const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

		var num big.Int
		num.SetBytes(randomBytes)

		var result []byte
		base := big.NewInt(62)
		zero := big.NewInt(0)

		for num.Cmp(zero) > 0 {
			mod := new(big.Int)
			num.DivMod(&num, base, mod)
			result = append([]byte{base62Chars[mod.Int64()]}, result...)
		}

		if len(result) == 0 {
			result = []byte{base62Chars[0]}
		}

		token = string(result)
		tokenKey := cache.GenKey("token", token)
		tokenValue := strings.Join([]string{secretID, secretKey}, ":")

		// Use SET with NX and EX options for atomic operation to avoid race condition
		// Only set if key doesn't exist
		cmd := cache.R.Do(
			context.Background(),
			"SET",
			tokenKey,
			tokenValue,
			"NX",
			"EX",
			time.Hour.Seconds(),
		)
		if cmd.Err() != nil {
			return "", cmd.Err()
		}

		// Result returns "OK" if SET was successful (when NX condition is met)
		value, err := cmd.Result()
		if err != nil {
			return "", err
		}
		if val, ok := value.(string); ok && val == "OK" {
			// Successfully set the token
			break
		}
		// Token already exists, regenerate and retry
	}

	return token, nil
}
