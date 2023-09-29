package value_objects

import (
	"errors"
	"unicode"

	"github.com/intwone/ddd-golang/internal/constants"
)

type Password struct {
	Value string
}

func NewPassword(value string) (*Password, []error) {
	errs := IsValidPassword(value)

	if len(errs) > 0 {
		return nil, errs
	}

	password := Password{Value: value}

	return &password, nil
}

func IsValidPassword(value string) []error {
	errs := []error{}

	if !HasMinimumCaracteres(value) {
		errs = append(errs, errors.New(constants.NotContainMinimumCaracteresPasswordError))
	}

	if !HasOneUpperCaseCaractere(value) {
		errs = append(errs, errors.New(constants.NotContainUpperCaseCharacterePasswordError))
	}

	if !HasOneSpecialCaractere(value) {
		errs = append(errs, errors.New(constants.NotContainSpecialCharacterePasswordError))
	}

	return errs
}

func HasMinimumCaracteres(value string) bool {
	miminumCaracteresInPassword := 8

	return len([]rune(value)) >= miminumCaracteresInPassword
}

func HasOneUpperCaseCaractere(value string) bool {
	for _, char := range value {
		if unicode.IsUpper(char) {
			return true
		}
	}

	return false
}

func HasOneSpecialCaractere(value string) bool {
	for _, char := range value {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return true
		}
	}

	return false
}

func (p *Password) ToStringPassword() string {
	return p.Value
}
