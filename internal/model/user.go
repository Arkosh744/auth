package model

import (
	"database/sql"
	"time"
)

type Role int

const (
	RoleUnknown Role = iota // 0
	RoleAdmin               // 1
	RoleUser                // 2
)

type User struct {
	ID       int64
	Username string
	Email    string
	Password string
	Role     Role

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type UserIdentifier struct {
	Username sql.NullString
	Email    sql.NullString
}

type UpdateUser struct {
	UserIdentifier
	Password sql.NullString
	Role     sql.NullString
}

func (r Role) String() string {
	switch r {
	case RoleAdmin:
		return "admin"
	case RoleUser:
		return "user"
	default:
		return ""
	}
}

func StringToRole(roleStr string) Role {
	switch roleStr {
	case "admin":
		return RoleAdmin
	case "user":
		return RoleUser
	default:
		return RoleUnknown
	}
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

type ExistsStatus int

const (
	StatusNone ExistsStatus = iota
	StatusUsernameExists
	StatusEmailExists
	StatusBothExist
)
