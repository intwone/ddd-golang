package use_cases

import (
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

type GetUserByEmailUseCaseInput struct {
	Email string
}

type GetUserByEmailUseCaseInterface interface {
	Execute(input GetUserByEmailUseCaseInput) (*enterprise.User, error)
}

type DefaultGetUserByEmailUseCase struct {
	UserRepository repositories.UserRepositoryInterface
}

func NewDefaulGetUserByEmailUseCase(userRepository repositories.UserRepositoryInterface) *DefaultGetUserByEmailUseCase {
	return &DefaultGetUserByEmailUseCase{
		UserRepository: userRepository,
	}
}

func (uc *DefaultGetUserByEmailUseCase) Execute(input GetUserByEmailUseCaseInput) (*enterprise.User, error) {
	user, err := uc.UserRepository.GetByEmail(input.Email)

	if err != nil {
		return nil, err
	}

	return user, nil
}
