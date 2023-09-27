package use_cases

import (
	"fmt"

	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

type CreateUserUseCaseInput struct {
	Name string
	Role string
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
	newUser := enterprise.NewUser(input.Name, input.Role)

	fmt.Println(newUser.GetID())

	err := uc.UserRepository.Create(newUser)

	if err != nil {
		return nil, err
	}

	return newUser, nil
}
