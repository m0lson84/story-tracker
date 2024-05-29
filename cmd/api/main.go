package main

import (
	"fmt"
	"os"

	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/m0lson84/story-tracker/api"
	"github.com/m0lson84/story-tracker/internal/config"
	"github.com/m0lson84/story-tracker/internal/database"
	"github.com/m0lson84/story-tracker/internal/logging"
	"github.com/m0lson84/story-tracker/internal/server"
)

var logger = logging.New()

type Options struct{}

func main() {
	logger.Info("Initializing application...")

	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		// Load application configuration
		env := os.Getenv("APP_ENV")
		config := config.NewConfig(env)

		// Initialize the Fiber server
		server := server.NewServer(config.App)

		// Initialize the database connection
		db := database.NewDatabase(config.DB)
		db.Connect()

		// Configure the API routes
		api.Setup(server.App, db)

		// Start the REST API server
		hooks.OnStart(func() {
			logger.Info("Starting application...")
			logger.Infof("Listening at http://localhost:%d", config.Port)
			err := server.Listen(fmt.Sprintf(":%d", config.Port))
			if err != nil {
				logger.Panicf("Cannot start server: %s", err)
			}
		})

		// Gracefully shutdown the server
		hooks.OnStop(func() {
			logger.Info("Shutting down application...")
			server.Shutdown()
			db.Close()
		})
	})

	cmd := cli.Root()
	cmd.Use = "story-tracker"
	cmd.Version = "0.1.0"

	cli.Run()
}
