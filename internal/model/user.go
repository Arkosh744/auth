package model

import (
	"database/sql"
	"fmt"
	"strings"
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
	switch strings.ToLower(roleStr) {
	case "admin":
		return RoleAdmin
	case "user":
		return RoleUser
	default:
		return RoleUnknown
	}
}

func (r *Role) Scan(value interface{}) error {
	if value == nil {
		*r = RoleUnknown
		return nil
	}

	roleStr, ok := value.(string)
	if !ok {
		return fmt.Errorf("cannot convert %v to string", value)
	}

	*r = StringToRole(roleStr)

	return nil
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
	StatusNone           ExistsStatus = iota // 0
	StatusUsernameExists                     // 1
	StatusEmailExists                        // 2
	StatusBothExist                          // 3
)
