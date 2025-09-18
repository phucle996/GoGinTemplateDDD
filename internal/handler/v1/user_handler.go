package v1handler

import (
	service_interface "authen_service/internal/domain/service"

	"github.com/gin-gonic/gin"
)

// user handler for adin managment api

type UserHandler struct {
	userService service_interface.UserService
}

func NewUserHandler(userService service_interface.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// implementation code

// POST /api/v1/users
func (u *UserHandler) CreateUser(c *gin.Context) {
	// imple code
}

// GET /api/v1/users
func (u *UserHandler) ListUsers(c *gin.Context) {
	// imple code
}

// GET /api/v1/users/:username
func (u *UserHandler) GetUserByUsername(c *gin.Context) {
	// imple code
}

// GET /api/v1/users/:email
func (u *UserHandler) GetUserByEmail(c *gin.Context) {
	// imple code
}

// PUT /api/v1/users/
func (u *UserHandler) UpdateUser(c *gin.Context) {
	// imple code
}

// PUT /api/v1/users/
func (u *UserHandler) SoftDeleteUser(c *gin.Context) {
	// imple code
}
