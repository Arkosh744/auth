package model

import (
	"encoding/json"
	"fmt"
)

type UserSpec struct {
	User
	Specialization Specializer
}

type UserSpecSerialized struct {
	User
	Specialization Specialization
}

type Specializer interface {
	Type() string
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

func (e *Engineer) Type() string { return SpecializationEngineer }
func (m *Manager) Type() string  { return SpecializationManager }

func (u *UserSpec) ToUserSpecSerialized() (*UserSpecSerialized, error) {
	var (
		specialization json.RawMessage
		err            error
	)

	switch u.Specialization.Type() {
	case SpecializationEngineer:
		specialization, err = json.Marshal(u.Specialization.(*Engineer))
	case SpecializationManager:
		specialization, err = json.Marshal(u.Specialization.(*Manager))
	}

	if err != nil {
		return nil, fmt.Errorf("failed to marshal specialization: %w", err)
	}

	return &UserSpecSerialized{
		User:           u.User,
		Specialization: Specialization{Type: u.Specialization.Type(), Attributes: specialization},
	}, nil
}

type Specialization struct {
	Type       string          `json:"type"`
	Attributes json.RawMessage `json:"attributes"`
}

func (u *UserSpecSerialized) ToUserSpec() (*UserSpec, error) {
	var (
		specializer Specializer
		err         error
	)

	switch u.Specialization.Type {
	case SpecializationEngineer:
		var engineer Engineer
		err = json.Unmarshal(u.Specialization.Attributes, &engineer)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal specialization to Engineer: %w", err)
		}
		specializer = &engineer

	case SpecializationManager:
		var manager Manager
		err = json.Unmarshal(u.Specialization.Attributes, &manager)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal specialization to Manager: %w", err)
		}
		specializer = &manager
	}

	return &UserSpec{
		User:           u.User,
		Specialization: specializer,
	}, nil
}
