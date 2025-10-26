package handler

import (
	"NextShortLink/internal/config"
	"NextShortLink/internal/logger"
	"NextShortLink/internal/middleware"
	"NextShortLink/internal/model"
	"fmt"
	"net/http"
	"time"

	"github.com/bytedance/sonic"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/gofiber/utils/v2"
	"go.uber.org/zap"
)

var app *fiber.App

// structValidator struct implementation
type structValidator struct {
	validate *validator.Validate
}

// Validate method implementation
func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

// fiberAPP provides a fiber app
func fiberAPP() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:     sonic.Marshal,
		JSONDecoder:     sonic.Unmarshal,
		ProxyHeader:     fiber.HeaderXForwardedFor,
		StructValidator: &structValidator{validate: validator.New()},
	})

	// Use requestID middleware
	app.Use(requestid.New(requestid.Config{
		Next:      nil,
		Header:    fiber.HeaderXRequestID,
		Generator: utils.UUIDv4,
	}))

	// Use global logger
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger.L,
		FieldsFunc: func(c fiber.Ctx) []zap.Field {
			requestID := requestid.FromContext(c)
			return []zap.Field{
				zap.String("requestID", requestID),
			}
		},
	}))

	// Use customer header middleware
	app.Use(middleware.CustomHeader)

	// Ping test router handler
	app.All("/ping", func(c fiber.Ctx) error {
		return model.RespSuccess(c, map[string]any{
			"msg":   "pong",
			"stamp": float64(time.Now().UnixNano()) / 1e9,
		})
	})

	// Root info router handler
	app.All("/", func(c fiber.Ctx) error {
		return model.RespSuccess(c, map[string]any{
			"poweredBy":   fmt.Sprintf("%s %s", config.ENName, config.Version),
			"techSupport": "Ghink Universe",
			"techContact": "service@ghink.net",
		})
	})

	// Register global routes
	Register(app)

	// Not found router handler
	app.Use(func(c fiber.Ctx) error {
		return model.RespNotFound(c)
	})

	return app
}

// RunHTTPServer runs a HTTP server
func RunHTTPServer() {
	// Create Fiber app
	app = fiberAPP()

	// Use fiber as handler
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.C.GetString("server.host"), config.C.GetInt("server.port")),
		Handler: adaptor.FiberApp(app),
	}

	server.SetKeepAlivesEnabled(true)

	// Start HTTP server
	go func() {
		logger.L.Fatal(server.ListenAndServe().Error())
	}()

	if config.C.GetBool("debug") {
		host := config.C.GetString("server.host")
		if host == "" {
			host = "0.0.0.0"
		}
		visit := host
		if host == "0.0.0.0" {
			visit = "localhost"
		}

		logger.L.Info(fmt.Sprintf("Server is running on %s:%d", host, config.C.GetInt("server.port")))
		logger.L.Debug(fmt.Sprintf("Visit by http://%s:%d", visit, config.C.GetInt("server.port")))
	}
}
