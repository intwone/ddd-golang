package use_cases

import (
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
	"github.com/intwone/ddd-golang/internal/infra/hasher"
)

type CreateUserUseCaseInput struct {
	Name     string
	Email    string
	Password string
	Role     string
}

type CreateUserUseCaseInterface interface {
	Execute(input CreateUserUseCaseInput) (*enterprise.User, []error)
}

type DefaultCreateUserUseCase struct {
	UserRepository repositories.UserRepositoryInterface
	Hasher         hasher.HasherInterface
}

func NewDefaultCreateUserUseCase(userRepository repositories.UserRepositoryInterface, hasher hasher.HasherInterface) *DefaultCreateUserUseCase {
	return &DefaultCreateUserUseCase{
		UserRepository: userRepository,
		Hasher:         hasher,
	}
}

func (uc *DefaultCreateUserUseCase) Execute(input CreateUserUseCaseInput) (*enterprise.User, []error) {
	newUser, newUserErrs := enterprise.NewUser(input.Name, input.Email, input.Password, input.Role)

	if len(newUserErrs) > 0 {
		return nil, newUserErrs
	}

	hashedPassword, hashErr := uc.Hasher.Hash(input.Password)

	if hashErr != nil {
		return nil, []error{hashErr}
	}

	newUser.SetPassword(*hashedPassword)

	createUserRepoErr := uc.UserRepository.Create(newUser)

	if createUserRepoErr != nil {
		return nil, []error{createUserRepoErr}
	}

	return newUser, nil
}
