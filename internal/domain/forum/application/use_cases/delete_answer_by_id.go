package use_cases

import (
	"errors"

	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
)

type DeleteAnswerByIDUseCaseInput struct {
	ID       string
	AuthorID string
}

type DeleteAnswerByIDUseCaseInterface interface {
	Execute(input DeleteAnswerByIDUseCaseInput) error
}

type DefaultDeleteAnswerByIDUseCase struct {
	AnswerRepository repositories.AnswerRepositoryInterface
}

func NewDefaultDeleteAnswerByIDUseCase(answerRepository repositories.AnswerRepositoryInterface) *DefaultDeleteAnswerByIDUseCase {
	return &DefaultDeleteAnswerByIDUseCase{
		AnswerRepository: answerRepository,
	}
}

func (uc *DefaultDeleteAnswerByIDUseCase) Execute(input DeleteAnswerByIDUseCaseInput) error {
	answer, err := uc.AnswerRepository.GetByID(input.ID)

	if err != nil {
		return err
	}

	if input.AuthorID != answer.GetAuthorID() {
		return errors.New("not allowed")
	}

	uc.AnswerRepository.DeleteByID(input.ID)

	return nil
}
