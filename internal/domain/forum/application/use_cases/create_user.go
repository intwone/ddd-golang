package use_cases

import (
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

type CreateUserUseCaseInput struct {
	Name     string
	Email    string
	Password string
	Role     string
}

type CreateUserUseCaseInterface interface {
	Execute(input CreateUserUseCaseInput) (*enterprise.User, error)
}

type DefaultCreateUserUseCase struct {
	UserRepository repositories.UserRepositoryInterface
}

func NewDefaultCreateUserUseCase(userRepository repositories.UserRepositoryInterface) *DefaultCreateUserUseCase {
	return &DefaultCreateUserUseCase{
		UserRepository: userRepository,
	}
}

func (uc *DefaultCreateUserUseCase) Execute(input CreateUserUseCaseInput) (*enterprise.User, error) {
	newUser, err := enterprise.NewUser(input.Name, input.Email, input.Password, input.Role)

	if err != nil {
		return nil, err
	}

	createUserRepoErr := uc.UserRepository.Create(newUser)

	if createUserRepoErr != nil {
		return nil, createUserRepoErr
	}

	return newUser, nil
}
