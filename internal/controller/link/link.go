package link

import (
	"NextShortLink/internal/model"
	"NextShortLink/internal/service/link"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

// Redirect to the destination
func Redirect(c fiber.Ctx) error {
	// Get link
	dest, err := link.GetLink(c.Params("linkID"))
	if err != nil {
		switch {
		case errors.Is(err, model.ErrLinkNotExist):
			return model.RespLinkNotExist(c)
		default:
			return model.RespInternalServerError(c, err)
		}
	}
	return c.Redirect().Status(http.StatusFound).To(dest)
}
