package stories

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/m0lson84/story-tracker/internal/database"
	"github.com/m0lson84/story-tracker/internal/stories"
	"go.uber.org/zap"
)

// Handler The handler for the story API endpoints.
type Handler struct {
	// Application logging instance
	logger *zap.SugaredLogger
	// The service for interacting with stories
	stories stories.Service
}

// NewHandler Create a new handler for the story API endpoints.
func NewHandler(logger *zap.SugaredLogger, db database.Service) *Handler {
	return &Handler{
		stories: stories.NewService(db),
		logger:  logger,
	}
}

// CreateStory Create a new story.
func (h *Handler) CreateStory(ctx context.Context, req *createStory) (*storyResponse, error) {
	h.logger.Info("Creating a new user story...")

	h.logger.Debug("Converting DTO to DB params...")
	params, err := req.ToParams()
	if err != nil {
		h.logger.Error("Error converting DTO to DB params: ", err)
		return nil, huma.Error400BadRequest("Invalid request params", err)
	}

	story, err := h.stories.CreateStory(ctx, params)
	if err != nil {
		h.logger.Error("Error creating story: ", err)
		return nil, huma.Error500InternalServerError("Internal server error", err)
	}

	return &storyResponse{Body: story}, nil
}

// DeleteStory Delete a story by ID.
func (h *Handler) DeleteStory(ctx context.Context, req *deleteStory) (*struct{}, error) {
	h.logger.Info("Deleting story by ID...")

	err := h.stories.DeleteStory(ctx, req.ID)
	if err != nil {
		h.logger.Error("Error deleting story: ", err)
		return nil, huma.Error500InternalServerError("Internal server error", err)
	}

	return nil, nil
}

// GetStory Get a story by ID.
func (h *Handler) GetStory(ctx context.Context, req *getStory) (*storyResponse, error) {
	h.logger.Info("Getting story by ID...")

	story, err := h.stories.GetStory(ctx, req.ID)
	if err != nil {
		h.logger.Error("Error getting story: ", err)
		return nil, huma.Error500InternalServerError("Internal server error", err)
	}

	return &storyResponse{Body: story}, nil
}

// ListStories List all stories.
func (h *Handler) ListStories(ctx context.Context, req *struct{}) (*storiesResponse, error) {
	h.logger.Info("Listing all stories...")

	stories, err := h.stories.ListStories(ctx)
	if err != nil {
		h.logger.Error("Error listing stories: ", err)
		return nil, huma.Error500InternalServerError("Internal server error", err)
	}

	return &storiesResponse{Body: stories}, nil
}

// UpdateStory Update a story by ID.
func (h *Handler) UpdateStory(ctx context.Context, req *updateStory) (*struct{}, error) {
	h.logger.Info("Updating story by ID...")

	h.logger.Debug("Converting DTO to DB params...")
	params, err := req.ToParams()
	if err != nil {
		h.logger.Error("Error converting DTO to DB params: ", err)
		return nil, huma.Error400BadRequest("Invalid request params", err)
	}

	if err := h.stories.UpdateStory(ctx, params); err != nil {
		h.logger.Error("Error updating user: ", err)
		return nil, huma.Error500InternalServerError("Internal server error", err)
	}

	return nil, nil
}
