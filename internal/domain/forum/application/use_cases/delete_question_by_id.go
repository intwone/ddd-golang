package use_cases

import (
	"errors"

	"github.com/intwone/ddd-golang/internal/constants"
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
)

type DeleteQuestionByIDUseCaseInput struct {
	ID       string
	AuthorID string
}

type DeleteQuestionByIDUseCaseInterface interface {
	Execute(input DeleteQuestionByIDUseCaseInput) error
}

type DefaultDeleteQuestionByIDUseCase struct {
	QuestionRepository repositories.QuestionRepositoryInterface
}

func NewDefaultDeleteQuestionByIDUseCase(questionRepository repositories.QuestionRepositoryInterface) *DefaultDeleteQuestionByIDUseCase {
	return &DefaultDeleteQuestionByIDUseCase{
		QuestionRepository: questionRepository,
	}
}

func (uc *DefaultDeleteQuestionByIDUseCase) Execute(input DeleteQuestionByIDUseCaseInput) error {
	question, err := uc.QuestionRepository.GetByID(input.ID)

	if err != nil {
		return err
	}

	if input.AuthorID != question.GetAuthorID() {
		return errors.New(constants.NotAllowedError)
	}

	uc.QuestionRepository.DeleteByID(input.ID)

	return nil
}
