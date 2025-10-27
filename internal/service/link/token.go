package link

import (
	"NextShortLink/internal/cache"
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
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

		// Check whether the token exists
		exists, err := cache.R.Exists(
			context.Background(),
			fmt.Sprintf("nextShortLink:token:%s", token),
		).Result()
		if err != nil {
			return "", err
		}
		if exists == 0 {
			token = string(result)
			break
		}
	}

	// Record token
	err := cache.R.Set(
		context.Background(),
		fmt.Sprintf("nextShortLink:token:%s", token),
		fmt.Sprintf("%s:%s", secretID, secretKey),
		time.Hour*1,
	).Err()
	if err != nil {
		return "", err
	}

	return token, nil
}
