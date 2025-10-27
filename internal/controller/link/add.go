package link

import (
	"NextShortLink/internal/model"
	"NextShortLink/internal/service/link"
	"errors"

	"github.com/gofiber/fiber/v3"
)

// AddLink adds a link
func AddLink(c fiber.Ctx) error {
	// Get request params
	var req model.RequestAddLink
	if err := c.Bind().JSON(&req); err != nil {
		return model.RespMissingParameter(c, req)
	}

	// Add link
	linkID, err := link.AddLink(req.Link, req.Validity)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrLinkInvalid):
			return model.RespLinkInvalid(c)
		case errors.Is(err, model.ErrValidityInvalid):
			return model.RespValidityInvalid(c)
		default:
			return model.RespInternalServerError(c, err)
		}
	}

	// Return linkID
	return model.RespSuccess(
		c, model.ReturnLinkID{LinkID: linkID},
	)
}
