package link

import (
	"NextShortLink/internal/model"
	"NextShortLink/internal/service/link"

	"github.com/gofiber/fiber/v3"
)

// IssueToken issue a token
func IssueToken(c fiber.Ctx) error {
	// Issue a token
	token, err := link.IssueToken(
		c.Locals("SecretID").(string),
		c.Locals("SecretKey").(string),
	)
	if err != nil {
		return model.RespInternalServerError(c, err)
	}

	// Return the token
	return model.RespSuccess(c, model.ReturnToken{
		Token: token,
	})
}
