package use_cases

import (
	"errors"

	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

type UpdateAnswerByIDUseCaseInput struct {
	ID       string
	AuthorID string
	Content  string
}

type UpdateAnswerByIDUseCaseInterface interface {
	Execute(input UpdateAnswerByIDUseCaseInput) (enterprise.Answer, error)
}

type DefaultUpdateAnswerByIDUseCase struct {
	AnswerRepository repositories.AnswerRepositoryInterface
}

func NewDefaultUpdateAnswerByIDUseCase(answerRepository repositories.AnswerRepositoryInterface) *DefaultUpdateAnswerByIDUseCase {
	return &DefaultUpdateAnswerByIDUseCase{
		AnswerRepository: answerRepository,
	}
}

func (uc *DefaultUpdateAnswerByIDUseCase) Execute(input UpdateAnswerByIDUseCaseInput) (enterprise.Answer, error) {
	answer, err := uc.AnswerRepository.GetByID(input.ID)

	if err != nil {
		return enterprise.Answer{}, err
	}

	if input.AuthorID != *answer.GetAuthorID().Value {
		return enterprise.Answer{}, errors.New("not allowed")
	}

	answer.SetContent(input.Content)

	err = uc.AnswerRepository.Save(answer)

	if err != nil {
		return enterprise.Answer{}, nil
	}

	return *answer, nil
}
