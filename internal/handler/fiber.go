package handler

import (
	"NextShortLink/internal/infra/config"
	"NextShortLink/internal/infra/logger"
	"NextShortLink/internal/middleware"
	"NextShortLink/internal/model"
	"fmt"
	"net/http"

	"github.com/ghinknet/json"
	"github.com/ghinknet/toolbox/expr"
	"github.com/go-playground/validator/v10"
	fiberzap "github.com/gofiber/contrib/v3/zap"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/gofiber/utils/v2"
	"go.uber.org/zap"
)

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
		JSONEncoder:     json.Marshal,
		JSONDecoder:     json.Unmarshal,
		ProxyHeader:     fiber.HeaderXForwardedFor,
		StructValidator: &structValidator{validate: validator.New()},
		ErrorHandler:    model.RespInternalServerError,
	})
	
	// Use recoverer
	app.Use(recoverer.New())

	// Use requestID middleware
	app.Use(requestid.New(requestid.Config{
		Next:      nil,
		Header:    fiber.HeaderXRequestID,
		Generator: utils.UUIDv4,
	}))

	// Use global logger
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger.L,
		Fields: []string{"latency", "status", "method", "url", "requestId", "ua"},
		FieldsFunc: func(c fiber.Ctx) []zap.Field {
			return []zap.Field{
				zap.String("ip", expr.Ternary(len(c.IPs()) > 0, c.IPs(), []string{c.IP()})[0]),
			}
		},
	}))

	// Use customer header middleware
	app.Use(middleware.CustomHeader)

	// Root info router handler
	app.Get("/", func(c fiber.Ctx) error {
		return c.Redirect().Status(http.StatusFound).To(config.Get().Index)
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
	app := fiberAPP()

	// Use Fiber as handler
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Get().Server.Host, config.Get().Server.Port),
		Handler: adaptor.FiberApp(app),
	}

	server.SetKeepAlivesEnabled(true)

	// Start HTTP server
	go func() {
		logger.L.Fatal("Failed to start main http service", zap.Error(server.ListenAndServe()))
	}()

	if config.Debug {
		host := config.Get().Server.Host
		if host == "" {
			host = "[::]"
		}
		visit := host
		if host == "[::]" || host == "0.0.0.0" {
			visit = "localhost"
		}

		logger.L.Info(fmt.Sprintf("Server is running on %s:%d", host, config.Get().Server.Port))
		logger.L.Debug(fmt.Sprintf("Visit by %s:%d", visit, config.Get().Server.Port))
	}
}
