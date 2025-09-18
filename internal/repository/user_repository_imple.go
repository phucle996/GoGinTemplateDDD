package repository_imple

import (
	"authen_service/internal/models"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepoImple struct {
	db *sqlx.DB
}

func NewUserRepoImple(db *sqlx.DB) *UserRepoImple {
	return &UserRepoImple{db: db}
}

// Create inserts a new user record into the database.
func (r *UserRepoImple) Create(ctx context.Context, user *models.User) error {

	// query
	query := `
		INSERT INTO users (
			user_id, username, email, fullname, passwd_hash, status
		) VALUES (
			:user_id, :username, :email, :fullname, :passwd_hash, :status
		)
	`
	// apply parameter in model for value in query
	_, err := r.db.NamedExecContext(ctx, query, user)
	if err != nil {
		return fmt.Errorf("user create failed: %w", err)
	}
	return nil
}

// Another imple
