package model

import (
	"time"
)

type Role int

const (
	RoleAdmin   Role = iota + 1 // 1
	RoleUser                    // 2
	RoleUnknown                 // 3
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
