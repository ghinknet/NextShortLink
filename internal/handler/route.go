package handler

import (
	"NextShortLink/internal/controller/link"

	"github.com/gofiber/fiber/v3"
)

// Register global routes
func Register(app *fiber.App) {
	app.Get("/:linkID", link.Redirect)
}
