package root

import (
	"context"

	"github.com/m0lson84/story-tracker/internal/database"
	"go.uber.org/zap"
)

// Handler The handler for the root API endpoints.
type Handler struct {
	// Application logging instance
	logger *zap.SugaredLogger
	// The service for interacting with the database
	db database.Service
}

// NewHandler Creates a new handler for the root API endpoints.
func NewHandler(logger *zap.SugaredLogger, db database.Service) *Handler {
	return &Handler{
		logger: logger,
		db:     db,
	}
}

// Root The root endpoint for the application.
func (h Handler) Root(_ context.Context, _ *struct{}) (*Response, error) {
	h.logger.Info("Root API called...")
	resp := Response{}

	// Prepare the response
	resp.Status = 200
	resp.Body.Message = "Hello world!"

	return &resp, nil
}

// CheckHealth Health check for the application.
func (h Handler) CheckHealth(_ context.Context, _ *struct{}) (*HealthResponse, error) {
	h.logger.Info("Health check called...")
	resp := HealthResponse{}

	// Check the database health
	health := h.db.Health()

	// Prepare the response
	resp.Status = 200
	resp.Body.DB = health

	return &resp, nil
}
