package entity

import "time"

// pure entity for using in service interface and imple
type User struct {
	UserId     int
	Username   string
	Email      string
	Fullname   string
	PasswdHash string
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}
