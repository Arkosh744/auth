package model

import (
	"database/sql"
	"time"
)

const (
	RoleUnknown = "unknown"
	RoleAdmin   = "admin"
	RoleUser    = "user"

	SpecializationEngineer = "engineer"
	SpecializationManager  = "manager"
)

const (
	StatusNone           ExistsStatus = iota // 0
	StatusUsernameExists                     // 1
	StatusEmailExists                        // 2
	StatusBothExist                          // 3
)

type ExistsStatus int

type User struct {
	ID       int64
	Username string
	Email    string
	Password string
	Role     string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type UserIdentifier struct {
	Username sql.NullString
	Email    sql.NullString
}

func (u *UserIdentifier) Set(username, email string) {
	if username != "" {
		u.Username.String = username
		u.Username.Valid = true
	}

	if email != "" {
		u.Email.String = email
		u.Email.Valid = true
	}
}

type UpdateUser struct {
	UserIdentifier
	Password sql.NullString
	Role     sql.NullString
}
