package service_imple

import (
	"authen_service/internal/domain/entity"
	repository_interface "authen_service/internal/domain/repository"
	service_interface "authen_service/internal/domain/service"
	"context"
)

type UserService struct {
	userRepo repository_interface.UserRepository
}

func NewUserService(userRepo repository_interface.UserRepository) service_interface.UserService {
	return &UserService{userRepo: userRepo}
}

// Create a new user
func (s *UserService) Create(ctx context.Context, user *entity.User) error {

	// imple code
	return nil
}

// Another imple func
