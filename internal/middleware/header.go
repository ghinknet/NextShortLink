package middleware

import (
	"NextShortLink/internal/env"

	"github.com/gofiber/fiber/v3"
)

// CustomHeader sets custom header
func CustomHeader(c fiber.Ctx) error {
	c.Set("X-Powered-By", env.PoweredByText)
	c.Set("X-Tech-Support", "Ghink Universe")
	c.Set("X-Tech-Contact", "service@ghink.net")

	return c.Next()
}
