package models

import (
	"time"

	"authen_service/internal/domain/entity"
)

type User struct {
	UserId     int       `db:"user_id"`
	Username   string    `db:"username"`
	Email      string    `db:"email"`
	Fullname   string    `db:"fullname"`
	PasswdHash string    `db:"passwd_hash"`
	Status     string    `db:"status"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

// Convert DB model → entity
func UserModelsToEntity(m *User) *entity.User {

	return &entity.User{
		UserId:     m.UserId,
		Username:   m.Username,
		Email:      m.Email,
		Fullname:   m.Fullname,
		PasswdHash: m.PasswdHash,
		Status:     m.Status,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
}

// Convert entity → DB model
func UserEntityToModels(e *entity.User) *User {

	return &User{
		UserId:     e.UserId,
		Username:   e.Username,
		Email:      e.Email,
		Fullname:   e.Fullname,
		PasswdHash: e.PasswdHash,
		Status:     e.Status,
		CreatedAt:  e.CreatedAt,
		UpdatedAt:  e.UpdatedAt,
	}
}
