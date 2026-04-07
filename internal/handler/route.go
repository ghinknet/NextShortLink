package handler

import (
	"NextShortLink/internal/controller/link"
	"NextShortLink/internal/middleware"

	"github.com/gofiber/fiber/v3"
)

// Register global routes
func Register(app *fiber.App) {
	app.Get("/t", middleware.CheckPermissionApplication, link.IssueToken)
	app.Post("/a", middleware.CheckPermissionApplication, link.AddLink)
	app.Get("/:linkID", link.Redirect)
}
