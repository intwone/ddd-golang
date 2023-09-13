package use_cases

import (
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
)

type DeleteQuestionByIDUseCaseInput struct {
	ID string
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
	_, err := uc.QuestionRepository.GetByID(input.ID)

	if err != nil {
		return err
	}

	uc.QuestionRepository.DeleteByID(input.ID)

	return nil
}
