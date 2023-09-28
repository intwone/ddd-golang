package value_objects

import (
	"errors"
	"regexp"

	"github.com/intwone/ddd-golang/internal/constants"
)

type Email struct {
	Value string
}

func NewEmail(value string) (*Email, error) {
	if !IsValidEmail(value) {
		return nil, errors.New(constants.InvalidEmailError)
	}

	email := Email{Value: value}

	return &email, nil
}

func IsValidEmail(value string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, value)

	return match
}
