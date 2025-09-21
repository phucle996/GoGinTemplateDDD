package dto

// RegisterRequest defines the structure for a user registration request.
type RegisterRequest struct {
	Username   string `json:"username" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Fullname   string `json:"fullname"`
	Password   string `json:"password" binding:"required,min=8"`
	RePassword string `json:"re_password" binding:"required"`
}

// LoginRequest defines the structure for a user login request.
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserResponse defines the structure for a user response.
type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Status   string `json:"status"`
}
