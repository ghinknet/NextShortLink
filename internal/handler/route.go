package handler

import (
	"NextShortLink/internal/controller/link"
	"NextShortLink/internal/middleware"

	"github.com/gofiber/fiber/v3"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
)

// Register global routes
func Register(app *fiber.App) {
	app.Use(recoverer.New())
	app.Get("/t", middleware.CheckPermissionApplication, link.IssueToken)
	app.Post("/a", middleware.CheckPermissionApplication, link.AddLink)
	app.Get("/:linkID", link.Redirect)
}
