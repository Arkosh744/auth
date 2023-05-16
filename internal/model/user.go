package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

const (
	RoleUnknown = "unknown"
	RoleAdmin   = "admin"
	RoleUser    = "user"

	SpecializationEngineer = "engineer"
	SpecializationManager  = "manager"
)

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

type UpdateUser struct {
	UserIdentifier
	Password sql.NullString
	Role     sql.NullString
}

type UserSpec struct {
	User
	Specialization
}

type UserSpecSerialized struct {
	User
	SpecializationSerialized []byte `db:"specialization"`
}

type Specialization struct {
	Type string `json:"type"`

	Engineer *Engineer `json:"engineer,omitempty"`
	Manager  *Manager  `json:"manager,omitempty"`
}

type Engineer struct {
	Level    int64  `json:"level"`
	Company  string `json:"company"`
	Language string `json:"language"`
}

type Manager struct {
	Level      int64  `json:"level"`
	Company    string `json:"company"`
	Department string `json:"department"`
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

func (u *UserSpec) ToUserSpecSerialized() (*UserSpecSerialized, error) {
	if u.Specialization == (Specialization{}) {
		return &UserSpecSerialized{
			User:                     u.User,
			SpecializationSerialized: []byte{},
		}, nil
	}

	specializationJSON, err := json.Marshal(u.Specialization)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal specialization: %w", err)
	}

	return &UserSpecSerialized{
		User:                     u.User,
		SpecializationSerialized: specializationJSON,
	}, nil
}

func (u *UserSpecSerialized) ToUserSpec() (*UserSpec, error) {
	if len(u.SpecializationSerialized) == 0 {
		// for old users without specialization
		return &UserSpec{
			User:           u.User,
			Specialization: Specialization{},
		}, nil
	}

	var specialization Specialization
	err := json.Unmarshal(u.SpecializationSerialized, &specialization)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal specialization: %w", err)
	}

	return &UserSpec{
		User:           u.User,
		Specialization: specialization,
	}, nil
}
