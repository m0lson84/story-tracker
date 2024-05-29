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
func (h Handler) Root(ctx context.Context, req *struct{}) (*rootResponse, error) {
	h.logger.Info("Root API called...")
	resp := rootResponse{}

	// Prepare the response
	resp.Status = 200
	resp.Body.Message = "Hello world!"

	return &resp, nil
}

// CheckHealth Health check for the application.
func (h Handler) CheckHealth(ctx context.Context, req *struct{}) (*healthResponse, error) {
	h.logger.Info("Health check called...")
	resp := healthResponse{}

	// Check the database health
	health := h.db.Health()

	// Prepare the response
	resp.Status = 200
	resp.Body.DB = health

	return &resp, nil
}
