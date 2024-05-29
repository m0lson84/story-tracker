package users

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/m0lson84/story-tracker/internal/database"
	"github.com/m0lson84/story-tracker/internal/logging"
)

// Routes registers the users routes with the application.
func Routes(api huma.API, db database.Service) {
	logger := logging.NewNamed("Routes.Users")

	logger.Info("Registering users routes...")
	h := NewHandler(logger, db)
	basePath := "/api/v1/users"

	// Create a new user
	huma.Register(api, huma.Operation{
		OperationID:   "create-story",
		Method:        http.MethodPost,
		Path:          basePath + "/",
		Summary:       "Create User",
		Description:   "Create a new user.",
		Tags:          []string{"Users"},
		DefaultStatus: 201,
	}, h.CreateUser)

	// Delete a user
	huma.Register(api, huma.Operation{
		OperationID:   "delete-user",
		Method:        http.MethodDelete,
		Path:          basePath + "/{id}",
		Summary:       "Delete User",
		Description:   "Delete a user by ID.",
		Tags:          []string{"Users"},
		DefaultStatus: 204,
	}, h.DeleteUser)

	// Get a user by ID
	huma.Register(api, huma.Operation{
		OperationID:   "get-user",
		Method:        http.MethodGet,
		Path:          basePath + "/{id}",
		Summary:       "Get User",
		Description:   "Get a user by ID.",
		Tags:          []string{"Users"},
		DefaultStatus: 200,
	}, h.GetUser)

	// Update a user
	huma.Register(api, huma.Operation{
		OperationID:   "update-user",
		Method:        http.MethodPut,
		Path:          basePath + "/{id}",
		Summary:       "Update User",
		Description:   "Update a user by ID.",
		Tags:          []string{"Users"},
		DefaultStatus: 204,
	}, h.UpdateUser)
}
