package model

import "time"

type UserRole string

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
)

type User struct {
	ID       int64
	Username string
	Email    string
	Password string
	Role     UserRole

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (ur UserRole) String() string {
	return string(ur)
}
