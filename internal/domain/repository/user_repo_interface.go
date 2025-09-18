package repository_interface

import (
	"authen_service/internal/models"
	"context"
)

// UserRepository defines the interface for user data operations.
// It abstracts the underlying data storage and provides a clean API for the service layer.
type UserRepository interface {

	// Create inserts a new user record into the database.
	Create(ctx context.Context, user *models.User) error
}
