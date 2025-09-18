package service_interface

import (
	"authen_service/internal/domain/entity"
	"context"
)

type UserService interface {
	// Create a new user

	Create(ctx context.Context, user *entity.User) error
}
