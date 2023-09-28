package value_objects

import "regexp"

type Email struct {
	Value string
}

func NewEmail(value string) *Email {
	if !IsValidEmail(value) {
		return nil
	}

	email := Email{Value: value}

	return &email
}

func IsValidEmail(value string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, value)

	return match
}
