package stories

import (
	"context"

	"github.com/m0lson84/story-tracker/db"
	"github.com/m0lson84/story-tracker/internal/database"
	"github.com/m0lson84/story-tracker/internal/logging"
	"go.uber.org/zap"
)

// Service represents a service that interacts with story records.
type Service interface {
	// CreateStory creates a new story record.
	// It returns the story record and an error if the operation failed.
	CreateStory(ctx context.Context, params db.CreateStoryParams) (db.Story, error)
	// DeleteStory deletes a story record.
	// It returns an error if the operation failed.
	DeleteStory(ctx context.Context, id int) error
	// GetStory retrieves a story record.
	// It returns the story record and an error if the operation failed.
	GetStory(ctx context.Context, id int) (db.Story, error)
	// ListStories retrieves all story records.
	// It returns a list of story records and an error if the operation failed.
	ListStories(ctx context.Context) ([]db.Story, error)
	// UpdateStory updates a story record.
	// It returns the updated story record and an error if the operation failed.
	UpdateStory(ctx context.Context, params db.UpdateStoryParams) error
}

// StoryService represents a service that interacts with story records.
type StoryService struct {
	// Application logging instance.
	logger *zap.SugaredLogger
	// Service that interacts with the database.
	db database.Service
}

// NewService creates a new instance of the story service.
func NewService(db database.Service) Service {
	logger := logging.NewNamed("Services.Stories")
	return &StoryService{
		db:     db,
		logger: logger,
	}
}

func (s *StoryService) CreateStory(ctx context.Context, params db.CreateStoryParams) (db.Story, error) {
	s.logger.Infow("Creating new story record...", "params", params)

	story, err := s.db.Execute().CreateStory(ctx, params)
	if err != nil {
		s.logger.Errorw("Error creating story record: ", err)
		return db.Story{}, err
	}

	return story, nil
}

func (s *StoryService) DeleteStory(ctx context.Context, id int) error {
	s.logger.Infow("Deleting story record...", "id", id)
	err := s.db.Execute().DeleteStory(ctx, int32(id))
	if err != nil {
		s.logger.Errorw("Error deleting story record: ", err)
		return err
	}
	return nil
}

func (s *StoryService) GetStory(ctx context.Context, id int) (db.Story, error) {
	s.logger.Infow("Getting story record...", "id", id)
	story, err := s.db.Execute().GetStory(ctx, int32(id))
	if err != nil {
		s.logger.Errorw("Error getting story record: ", err)
		return db.Story{}, err
	}
	return story, nil
}

func (s *StoryService) ListStories(ctx context.Context) ([]db.Story, error) {
	s.logger.Info("Listing story records...")
	stories, err := s.db.Execute().ListStories(ctx)
	if err != nil {
		s.logger.Errorw("Error listing story records: ", err)
		return nil, err
	}
	return stories, nil
}

func (s *StoryService) UpdateStory(ctx context.Context, params db.UpdateStoryParams) error {
	s.logger.Infow("Updating story record...", "params", params)

	err := s.db.Execute().UpdateStory(ctx, params)
	if err != nil {
		s.logger.Errorw("Error updating story record: ", err)
		return err
	}

	return nil
}
