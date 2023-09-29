package hasher

import (
	"errors"

	"github.com/intwone/ddd-golang/internal/constants"
	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct{}

func NewBcryptHasher() *BcryptHasher {
	return &BcryptHasher{}
}

func (b *BcryptHasher) Hash(value string) (*string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)

	if err != nil {
		return nil, errors.New(constants.GenerateHashError)
	}

	hashedString := string(hashed)

	return &hashedString, nil
}

func (b *BcryptHasher) Compare(value string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))

	return err == nil
}
