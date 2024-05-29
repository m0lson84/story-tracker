package users

import (
	"context"

	"github.com/m0lson84/story-tracker/db"
	"github.com/m0lson84/story-tracker/internal/database"
	"github.com/m0lson84/story-tracker/internal/logging"
	"go.uber.org/zap"
)

// Service represents a service that interacts with user records.
type Service interface {
	// CreateUser creates a new user record.
	// It returns the user record and an error if the operation failed.
	CreateUser(ctx context.Context, username string) (db.User, error)
	// DeleteUser deletes a user record.
	// It returns an error if the operation failed.
	DeleteUser(ctx context.Context, id int) error
	// GetUser retrieves a user record.
	// It returns the user record and an error if the operation failed.
	GetUser(ctx context.Context, id int) (db.User, error)
	// UpdateUser updates a user record.
	// It returns the updated user record and an error if the operation failed.
	UpdateUser(ctx context.Context, id int, username string) error
}

// UserService represents a service that interacts with user records.
type UserService struct {
	// Application logging instance.
	logger *zap.SugaredLogger
	// Service that interacts with the database.
	db database.Service
}

// NewService creates a new instance of the user service.
func NewService(db database.Service) Service {
	logger := logging.NewNamed("Services.Users")
	return &UserService{
		db:     db,
		logger: logger,
	}
}

func (s UserService) CreateUser(ctx context.Context, username string) (db.User, error) {
	s.logger.Infow("Creating new user record...", "username", username)
	user, err := s.db.Execute().CreateUser(ctx, username)
	if err != nil {
		s.logger.Errorw("Error creating user record...", "error", err)
		return db.User{}, err
	}
	return user, nil
}

func (s UserService) DeleteUser(ctx context.Context, id int) error {
	s.logger.Infow("Deleting user record...", "id", id)
	err := s.db.Execute().DeleteUser(ctx, int32(id))
	if err != nil {
		s.logger.Errorw("Error deleting user record...", "error", err)
		return err
	}
	return nil
}

func (s UserService) GetUser(ctx context.Context, id int) (db.User, error) {
	s.logger.Infow("Getting user record...", "id", id)
	user, err := s.db.Execute().GetUser(ctx, int32(id))
	if err != nil {
		s.logger.Errorw("Error getting user record...", "error", err)
		return db.User{}, err
	}
	return user, nil
}

func (s UserService) UpdateUser(ctx context.Context, id int, username string) error {
	s.logger.Infow("Updating user record...", "id", id, "username", username)
	err := s.db.Execute().UpdateUser(ctx, db.UpdateUserParams{ID: int32(id), Username: username})
	if err != nil {
		s.logger.Errorw("Error updating user record...", "error", err)
		return err
	}
	return nil
}
