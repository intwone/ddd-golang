package value_objects

import (
	"errors"

	"github.com/intwone/ddd-golang/internal/constants"
)

type Role struct {
	Value string
}

func NewRole(value string) (*Role, error) {
	if !IsValidRole(value) {
		return nil, errors.New(constants.InvalidRoleError)
	}

	role := Role{Value: value}

	return &role, nil
}

func IsValidRole(value string) bool {
	switch value {
	case "student", "instructor":
		return true
	default:
		return false
	}
}
