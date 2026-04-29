package model

import (
	"NextShortLink/internal/infra/logger"
	"net/http"
	"reflect"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"go.gh.ink/json"
	"go.uber.org/zap"
)

type response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func init() {
	_ = json.Preheat(reflect.TypeOf(response[any]{}))
}

func Resp[T any](c fiber.Ctx, code int, data T, msg string) error {
	return c.Status(http.StatusOK).JSON(response[T]{code, msg, data})
}

// --------------- 200 ---------------

func RespSuccess[T any](c fiber.Ctx, data T) error {
	return Resp(c, http.StatusOK, data, "success")
}

// --------------- 400 ---------------

func RespBadRequest(c fiber.Ctx) error {
	return Resp(c, http.StatusBadRequest, any(nil), "bad request")
}

func RespNotFound(c fiber.Ctx) error {
	return Resp(c, http.StatusNotFound, any(nil), "not found")
}

func RespMethodNotAllowed(c fiber.Ctx) error {
	return Resp(c, http.StatusMethodNotAllowed, any(nil), "method not allowed")
}

func RespTeaPot[T any](c fiber.Ctx, data T) error {
	return Resp(c, http.StatusTeapot, data, "I'm a tea pot")
}

func RespTooManyRequests(c fiber.Ctx) error {
	return Resp(c, http.StatusTooManyRequests, any(nil), "too many requests")
}

// --------------- 500 ---------------

func RespInternalServerError(c fiber.Ctx, err error) error {
	requestID := requestid.FromContext(c)
	logger.L.Error(
		"internal server error happened",
		zap.Error(err),
		zap.String("requestID", requestID),
	)
	return Resp(c, http.StatusInternalServerError, any(nil), "internal server error")
}

// --------------- 800 ---------------

func RespMissingParameter[T any](c fiber.Ctx, params T) error {
	return Resp(c, 800, params, "missing parameter")
}

func RespPermissionDenied(c fiber.Ctx) error {
	return Resp(c, 801, any(nil), ErrPermissionDenied.Error())
}

func RespNoPackageAvailable(c fiber.Ctx) error {
	return Resp(c, 802, any(nil), ErrNoPackageAvailable.Error())
}

func RespApplicationNotFound(c fiber.Ctx) error {
	return Resp(c, 803, any(nil), ErrApplicationNotFound.Error())
}

func RespLinkNotExist(c fiber.Ctx) error {
	return Resp(c, 804, any(nil), ErrLinkNotExist.Error())
}

func RespLinkInvalid(c fiber.Ctx) error {
	return Resp(c, 805, any(nil), ErrLinkInvalid.Error())
}

func RespValidityInvalid(c fiber.Ctx) error {
	return Resp(c, 806, any(nil), ErrValidityInvalid.Error())
}
