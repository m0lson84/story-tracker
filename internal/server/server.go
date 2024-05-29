package server

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/redirect"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/m0lson84/story-tracker/internal/config"
	"github.com/m0lson84/story-tracker/internal/logging"
)

// FiberServer represents a fiber server.
type FiberServer struct {
	// The fiber application instance.
	*fiber.App
}

// NewServer creates a new fiber server with the given configuration.
func NewServer(config config.App) *FiberServer {
	zap := fiberzap.NewLogger(fiberzap.LoggerConfig{
		SetLogger: logging.New().Desugar(),
		ExtraKeys: []string{"requestid"},
	})
	fiberlog.SetLogger(zap)

	app := fiber.New(fiber.Config{
		ServerHeader: "story-tracker",
		AppName:      "story-tracker",
	})

	// Add middleware
	app.Use(requestid.New())
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(logger.New(logger.Config{
		Format:        "${time} [DEBUG] ${status} - ${method} ${path} (RequestID: ${locals:requestid})\n",
		DisableColors: config.Env == "production",
	}))

	// Configure docs
	app.Use(redirect.New(redirect.Config{
		StatusCode: 301,
		Rules: map[string]string{
			"/docs": "/docs/index.html",
		},
	}))
	app.Get("/docs/index.html", func(ctx *fiber.Ctx) error {
		ctx.Response().Header.Set("Content-Type", "text/html")
		ctx.Response().BodyWriter().Write([]byte(`<!doctype html>
<html>
  <head>
    <title>API Reference</title>
    <meta charset="utf-8" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1" />
  </head>
  <body>
    <script
      id="api-reference"
      data-url="/openapi.json"></script>
    <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
  </body>
</html>`))
		return nil
	})

	return &FiberServer{
		App: app,
	}
}
