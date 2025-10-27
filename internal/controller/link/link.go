package link

import (
	"NextShortLink/internal/model"
	"NextShortLink/internal/service/link"
	"errors"

	"github.com/gofiber/fiber/v3"
)

// Redirect to the destination
func Redirect(c fiber.Ctx) error {
	// Get link
	dest, err := link.GetLink(c.Params("linkID"))
	if err != nil {
		if errors.Is(err, model.ErrLinkNotExist) {
			return model.RespLinkNotExist(c)
		}
		return model.RespInternalServerError(c, err)
	}
	return c.Redirect().Status(model.CodeFound).To(dest)
}
