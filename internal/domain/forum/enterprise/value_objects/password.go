package value_objects

import "unicode"

type Password struct {
	Value string
}

func NewPassword(value string) (*Password, []string) {
	isValid, errs := IsValidPassword(value)

	if !isValid {
		return nil, errs
	}

	password := Password{Value: value}

	return &password, nil
}

func IsValidPassword(value string) (bool, []string) {
	errors := []string{}
	valid := true

	if !HasMinimumCaracteres(value) {
		valid = false
		errors = append(errors, "the password must contain at least eight characters long")
	}

	if !HasOneUpperCaseCaractere(value) {
		valid = false
		errors = append(errors, "the password must contain at least one uppercase character")
	}

	if !HasOneSpecialCaractere(value) {
		valid = false
		errors = append(errors, "the password must contain at least one special character")
	}

	return valid, errors
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