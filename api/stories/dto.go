package stories

import (
	"github.com/m0lson84/story-tracker/db"
)

// createStory represents the data transfer object for creating a story.
type createStory struct {
	// The body of the request.
	Body struct {
		// The title of the story.
		Title string `json:"title" example:"Do Work"`
		// The type of story.
		Type string `json:"type" example:"feature"`
		// The current status of the story.
		Status string `json:"status" enum:"unstarted,started,finished,delivered,rejected,accepted" example:"unstarted"`
		// The number of points associated with the story.
		Points string `json:"points" enum:"none,one,two,three,five,eight,thirteen" example:"three"`
		// The description of the story.
		Description string `json:"description" example:"As a user, I want to do work so that I can be productive."`
		// The identifier of the user tasked with the story.
		UserID int32 `json:"userId" example:"1"`
	}
}

// ToParams converts the DTO to a set of database parameters.
func (c createStory) ToParams() (db.CreateStoryParams, error) {
	params := db.CreateStoryParams{
		UserID:      &c.Body.UserID,
		Title:       c.Body.Title,
		Description: c.Body.Description,
	}

	if err := params.Type.Scan(c.Body.Type); err != nil {
		return db.CreateStoryParams{}, err
	}

	if err := params.Status.Scan(c.Body.Status); err != nil {
		return db.CreateStoryParams{}, err
	}

	if err := params.Points.Scan(c.Body.Points); err != nil {
		return db.CreateStoryParams{}, err
	}

	return params, nil
}

// deleteStory represents the request for deleting a story.
type deleteStory struct {
	// The unique identifier of the story to delete.
	ID int `path:"id" doc:"The unique identifier of the story to delete"`
}

// getStory represents the request for retrieving a story.
type getStory struct {
	// The unique identifier of the story to retrieve.
	ID int `path:"id" doc:"The unique identifier of the story to retrieve"`
}

// updateStory represents the request for updating an existing story.
type updateStory struct {
	// The body of the request.
	Body struct {
		// The identifier of the user tasked with the story.
		UserID *int32 `json:"userId,omitempty" example:"1"`
		// The title of the story.
		Title *string `json:"title,omitempty" example:"Do Work"`
		// The type of story.
		Type *string `json:"type,omitempty" example:"feature"`
		// The current status of the story.
		Status *string `json:"status,omitempty" enum:"unstarted,started,finished,delivered,rejected,accepted" example:"unstarted"`
		// The number of points associated with the story.
		Points *string `json:"points,omitempty" enum:"none,one,two,three,five,eight,thirteen" example:"three"`
		// The description of the story.
		Description *string `json:"description,omitempty" example:"As a user, I want to do work so that I can be productive."`
	}
	// The unique identifier of the story to update.
	ID int32 `path:"id" doc:"The unique identifier of the story to update"`
}

// ToParams converts the DTO to a set of database parameters.
func (u updateStory) ToParams() (db.UpdateStoryParams, error) {
	params := db.UpdateStoryParams{
		ID:          int32(u.ID),
		Title:       u.Body.Title,
		Description: u.Body.Description,
		UserID:      u.Body.UserID,
	}

	if u.Body.Type != nil {
		if err := params.Type.Scan(*u.Body.Type); err != nil {
			return db.UpdateStoryParams{}, err
		}
	}

	if u.Body.Status != nil {
		if err := params.Status.Scan(*u.Body.Status); err != nil {
			return db.UpdateStoryParams{}, err
		}
	}

	if u.Body.Points != nil {
		if err := params.Points.Scan(*u.Body.Points); err != nil {
			return db.UpdateStoryParams{}, err
		}
	}

	return params, nil
}

// storyResponse represents the response for a single story.
type storyResponse struct {
	// The user story.
	Body db.Story `json:"body" doc:"The story."`
}

// storiesResponse represents the response for a list of stories.
type storiesResponse struct {
	// The list of stories.
	Body []db.Story `json:"body" doc:"The list of stories."`
}
