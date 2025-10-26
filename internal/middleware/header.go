package middleware

import (
	"NextShortLink/internal/config"
	"fmt"

	"github.com/gofiber/fiber/v3"
)

// CustomHeader sets custom header
func CustomHeader(c fiber.Ctx) error {
	c.Set("X-Powered-By", fmt.Sprintf("%s %s %s", config.ENName, config.Version))
	c.Set("X-Tech-Support", "Ghink Universe")
	c.Set("X-Tech-Contact", "service@ghink.net")

	return c.Next()
}
