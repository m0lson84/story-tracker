package stories

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/m0lson84/story-tracker/internal/database"
	"github.com/m0lson84/story-tracker/internal/logging"
)

// Routes registers the story routes with the application.
func Routes(api huma.API, db database.Service) {
	logger := logging.NewNamed("Routes.Stories")

	logger.Info("Registering story routes...")
	h := NewHandler(logger, db)
	basePath := "/api/v1/stories"

	// Create a new story
	huma.Register(api, huma.Operation{
		OperationID:   "create-story",
		Method:        http.MethodPost,
		Path:          basePath + "/",
		Summary:       "Create Story",
		Description:   "Create a new user story.",
		Tags:          []string{"Stories"},
		DefaultStatus: 201,
	}, h.CreateStory)

	// Delete a story
	huma.Register(api, huma.Operation{
		OperationID:   "delete-story",
		Method:        http.MethodDelete,
		Path:          basePath + "/{id}",
		Summary:       "Delete Story",
		Description:   "Delete a user story by ID.",
		Tags:          []string{"Stories"},
		DefaultStatus: 204,
	}, h.DeleteStory)

	// Get a story by ID
	huma.Register(api, huma.Operation{
		OperationID:   "get-story",
		Method:        http.MethodGet,
		Path:          basePath + "/{id}",
		Summary:       "Get Story",
		Description:   "Get a user story by ID.",
		Tags:          []string{"Stories"},
		DefaultStatus: 200,
	}, h.GetStory)

	// List all stories
	huma.Register(api, huma.Operation{
		OperationID:   "get-story",
		Method:        http.MethodGet,
		Path:          basePath + "/",
		Summary:       "List Stories",
		Description:   "List all user stories.",
		Tags:          []string{"Stories"},
		DefaultStatus: 200,
	}, h.ListStories)

	// Update a story
	huma.Register(api, huma.Operation{
		OperationID:   "update-story",
		Method:        http.MethodPut,
		Path:          basePath + "/{id}",
		Summary:       "Update Story",
		Description:   "Update a user story by ID.",
		Tags:          []string{"Stories"},
		DefaultStatus: 204,
	}, h.UpdateStory)
}
