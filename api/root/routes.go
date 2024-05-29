package root

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/m0lson84/story-tracker/internal/database"
	"github.com/m0lson84/story-tracker/internal/logging"
)

// Routes registers the root API routes with the application.
func Routes(api huma.API, db database.Service) {
	logger := logging.NewNamed("Routes.Root")

	logger.Info("Registering story routes...")
	h := NewHandler(logger, db)

	huma.Register(api, huma.Operation{
		OperationID: "root",
		Method:      http.MethodGet,
		Path:        "/",
		Summary:     "Root API",
		Description: "The root endpoint for the application.",
		Tags:        []string{"Root"},
	}, h.Root)
	huma.Register(api, huma.Operation{
		OperationID: "check-health",
		Method:      http.MethodGet,
		Path:        "/health",
		Summary:     "Check Health",
		Description: "Health check for the application.",
		Tags:        []string{"Root"},
	}, h.CheckHealth)
}
