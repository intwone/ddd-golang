package use_cases

import (
	"errors"

	"github.com/intwone/ddd-golang/internal/constants"
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"
	"github.com/intwone/ddd-golang/internal/infra/cryptography"
	"github.com/intwone/ddd-golang/internal/infra/hasher"
)

type AuthenticateUseCaseInput struct {
	Email    string
	Password string
}

type AuthenticateUseCaseInterface interface {
	Execute(input AuthenticateUseCaseInput) (*string, []error)
}

type DefaultAuthenticateUseCase struct {
	UserRepository repositories.UserRepositoryInterface
	Hasher         hasher.HasherInterface
	Cryptography   cryptography.CryptographyInterface
}

func NewDefaulAuthenticateUseCase(userRepository repositories.UserRepositoryInterface, hasher hasher.HasherInterface, crytography cryptography.CryptographyInterface) *DefaultAuthenticateUseCase {
	return &DefaultAuthenticateUseCase{
		UserRepository: userRepository,
		Hasher:         hasher,
		Cryptography:   crytography,
	}
}

func (uc *DefaultAuthenticateUseCase) Execute(input AuthenticateUseCaseInput) (*string, []error) {
	email, newEmailErr := vo.NewEmail(input.Email)

	if newEmailErr != nil {
		return nil, []error{newEmailErr}
	}

	password, newPasswordErrs := vo.NewPassword(input.Password)

	if len(newPasswordErrs) > 0 {
		return nil, newPasswordErrs
	}

	user, getByEmailErr := uc.UserRepository.GetByEmail(email.Value)

	if getByEmailErr != nil {
		return nil, []error{getByEmailErr}
	}

	isValid := uc.Hasher.Compare(password.Value, user.GetPassword())

	if !isValid {
		return nil, []error{errors.New(constants.PasswordAreNotTheSame)}
	}

	token, encryptErr := uc.Cryptography.Encrypt(user.GetID())

	if encryptErr != nil {
		return nil, []error{encryptErr}
	}

	return token, nil
}
