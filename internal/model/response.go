package model

import (
	"NextShortLink/internal/logger"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"go.uber.org/zap"
)

func Resp(c fiber.Ctx, code int, data any, msg string) error {
	type response struct {
		Code int    `json:"code"`
		Data any    `json:"data"`
		Msg  string `json:"msg"`
	}
	return c.Status(fiber.StatusOK).JSON(response{code, data, msg})
}

func RespSuccess(c fiber.Ctx, data any) error {
	return Resp(c, CodeOK, data, "success")
}

func RespBadRequest(c fiber.Ctx) error {
	return Resp(c, CodeBadRequest, nil, "bad request")
}

func RespNotFound(c fiber.Ctx) error {
	return Resp(c, CodeNotFound, nil, "not found")
}

func RespMethodNotAllowed(c fiber.Ctx) error {
	return Resp(c, CodeMethodNotAllowed, nil, "method not allowed")
}

func RespTeaPot(c fiber.Ctx, data any) error {
	return Resp(c, CodeTeaPot, data, "I'm a tea pot")
}

func RespTooManyRequests(c fiber.Ctx) error {
	return Resp(c, CodeTooManyRequests, nil, "too many requests")
}

func RespInternalServerError(c fiber.Ctx, err error) error {
	requestID := requestid.FromContext(c)
	logger.L.Error(
		err.Error(),
		zap.String("requestID", requestID),
	)
	return Resp(c, CodeInternalServerError, nil, "internal server error")
}

func RespMissingParameter(c fiber.Ctx, params any) error {
	return Resp(c, CodeMissingParameter, params, "missing parameter")
}

func RespPermissionDenied(c fiber.Ctx) error {
	return Resp(c, CodePermissionDenied, nil, "permission denied")
}

func RespNoPackageAvailable(c fiber.Ctx) error {
	return Resp(c, CodeNoPackageAvailable, nil, "no package available")
}

func RespApplicationNotFound(c fiber.Ctx) error {
	return Resp(c, CodeApplicationNotFound, nil, "application not found")
}

func RespLinkNotExist(c fiber.Ctx) error {
	return Resp(c, CodeLinkNotExist, nil, "link does not exist")
}

func RespLinkInvalid(c fiber.Ctx) error {
	return Resp(c, CodeLinkInvalid, nil, "link invalid")
}

func RespValidityInvalid(c fiber.Ctx) error {
	return Resp(c, CodeValidityInvalid, nil, "validity invalid")
}
