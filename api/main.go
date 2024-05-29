package api

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/m0lson84/story-tracker/api/root"
	"github.com/m0lson84/story-tracker/api/stories"
	"github.com/m0lson84/story-tracker/api/users"
	"github.com/m0lson84/story-tracker/internal/database"
	"github.com/m0lson84/story-tracker/internal/logging"
)

// Setup configures the REST API routes for the application.
func Setup(app *fiber.App, db database.Service) {
	logger := logging.NewNamed("Routes")

	logger.Debug("Configuring OpenAPI docs...")
	config := huma.DefaultConfig("Story Tracker API", "0.1")
	config.DocsPath = ""
	api := humafiber.New(app, config)

	root.Routes(api, db)
	stories.Routes(api, db)
	users.Routes(api, db)
}
