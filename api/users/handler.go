package users

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/m0lson84/story-tracker/internal/database"
	"github.com/m0lson84/story-tracker/internal/users"
	"go.uber.org/zap"
)

// Handler is the HTTP handler for the users API.
type Handler struct {
	// Application logging instance.
	logger *zap.SugaredLogger
	// Service for managing users.
	users users.Service
}

// NewHandler creates a new handler for the users API.
func NewHandler(logger *zap.SugaredLogger, db database.Service) *Handler {
	return &Handler{
		users:  users.NewService(db),
		logger: logger,
	}
}

// CreateUser Create a new user.
func (h *Handler) CreateUser(ctx context.Context, req *createUser) (*userResponse, error) {
	h.logger.Info("Creating a new user...")

	user, err := h.users.CreateUser(ctx, req.Body.Username)
	if err != nil {
		h.logger.Error("Error creating user: ", err)
		return nil, huma.Error500InternalServerError("Internal server error", err)
	}

	resp := userResponse{}
	resp.FromDB(user)

	return &resp, nil
}

// DeleteUser Delete a user.
func (h *Handler) DeleteUser(ctx context.Context, req *deleteUser) (*struct{}, error) {
	h.logger.Info("Deleting user by ID...")

	err := h.users.DeleteUser(ctx, req.ID)
	if err != nil {
		h.logger.Error("Error deleting user: ", err)
		return nil, huma.Error500InternalServerError("Internal server error", err)
	}

	return nil, nil
}

// GetUser Get a user by ID.
func (h *Handler) GetUser(ctx context.Context, req *getUser) (*userResponse, error) {
	h.logger.Info("Getting user by ID...")

	user, err := h.users.GetUser(ctx, req.ID)
	if err != nil {
		h.logger.Error("Error getting user: ", err)
		return nil, huma.Error500InternalServerError("Internal server error", err)
	}

	resp := userResponse{}
	resp.FromDB(user)

	return &resp, nil
}

// UpdateUser Update a user by ID.
func (h *Handler) UpdateUser(ctx context.Context, req *updateUser) (*struct{}, error) {
	h.logger.Info("Updating user by ID...")

	if err := h.users.UpdateUser(ctx, req.ID, req.Body.Username); err != nil {
		h.logger.Error("Error updating user: ", err)
		return nil, huma.Error500InternalServerError("Internal server error", err)
	}

	return nil, nil
}
