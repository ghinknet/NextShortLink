package model

import (
	"NextShortLink/internal/logger"
	"net/http"

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
	return c.Status(http.StatusOK).JSON(response{code, data, msg})
}

// --------------- 200 ---------------

func RespSuccess(c fiber.Ctx, data any) error {
	return Resp(c, http.StatusOK, data, "success")
}

// --------------- 400 ---------------

func RespBadRequest(c fiber.Ctx) error {
	return Resp(c, http.StatusBadRequest, nil, "bad request")
}

func RespNotFound(c fiber.Ctx) error {
	return Resp(c, http.StatusNotFound, nil, "not found")
}

func RespMethodNotAllowed(c fiber.Ctx) error {
	return Resp(c, http.StatusMethodNotAllowed, nil, "method not allowed")
}

func RespTeaPot(c fiber.Ctx, data any) error {
	return Resp(c, http.StatusTeapot, data, "I'm a tea pot")
}

func RespTooManyRequests(c fiber.Ctx) error {
	return Resp(c, http.StatusTooManyRequests, nil, "too many requests")
}

// --------------- 500 ---------------

func RespInternalServerError(c fiber.Ctx, err error) error {
	requestID := requestid.FromContext(c)
	logger.L.Error(
		"internal server error happened",
		zap.Error(err),
		zap.String("requestID", requestID),
	)
	return Resp(c, http.StatusInternalServerError, nil, "internal server error")
}

// --------------- 800 ---------------

func RespMissingParameter(c fiber.Ctx, params any) error {
	return Resp(c, 800, params, "missing parameter")
}

func RespPermissionDenied(c fiber.Ctx) error {
	return Resp(c, 801, nil, ErrPermissionDenied.Error())
}

func RespNoPackageAvailable(c fiber.Ctx) error {
	return Resp(c, 802, nil, ErrNoPackageAvailable.Error())
}

func RespApplicationNotFound(c fiber.Ctx) error {
	return Resp(c, 803, nil, ErrApplicationNotFound.Error())
}

func RespLinkNotExist(c fiber.Ctx) error {
	return Resp(c, 804, nil, ErrLinkNotExist.Error())
}

func RespLinkInvalid(c fiber.Ctx) error {
	return Resp(c, 805, nil, ErrLinkInvalid.Error())
}

func RespValidityInvalid(c fiber.Ctx) error {
	return Resp(c, 806, nil, ErrValidityInvalid.Error())
}
