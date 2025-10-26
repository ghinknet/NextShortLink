package web

import (
	"strings"

	"github.com/gofiber/fiber/v3"
)

// ClientIP returns real client IP
func ClientIP(c fiber.Ctx) string {
	return strings.Split(strings.TrimSpace(c.IP()), ",")[0]
}
