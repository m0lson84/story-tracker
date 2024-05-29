package users

import "github.com/m0lson84/story-tracker/db"

// user represents a user in the system.
type user struct {
	// The name of the user.
	Username string `json:"username" example:"john.doe"`
	// The unique identifier of the user record.
	ID int `json:"id" example:"1"`
}

// createUser represents the request for creating a new user.
type createUser struct {
	// The body of the request.
	Body struct {
		// The name of the user.
		Username string `json:"username" example:"john.doe"`
	}
}

// deleteUser represents the request for deleting a user.
type deleteUser struct {
	// The unique identifier of the user to delete.
	ID int `path:"id" doc:"The unique identifier of the user to delete"`
}

// getUser represents the request for retrieving a user.
type getUser struct {
	// The unique identifier of the user to retrieve.
	ID int `path:"id" doc:"The unique identifier of the user to retrieve"`
}

// updateUser represents the request for updating an existing user.
type updateUser struct {
	// The body of the request.
	Body struct {
		// The name of the user.
		Username string `json:"username" example:"john.doe"`
	}
	// The unique identifier of the user to update.
	ID int `path:"id" doc:"The unique identifier of the user to update"`
}

// userResponse represents the response for a single user.
type userResponse struct {
	// The user.
	Body user `json:"body" doc:"The user."`
}

// FromDB converts a database user record to a user model and populates the response.
func (resp *userResponse) FromDB(u db.User) {
	resp.Body = user{
		ID:       int(u.ID),
		Username: u.Username,
	}
}
