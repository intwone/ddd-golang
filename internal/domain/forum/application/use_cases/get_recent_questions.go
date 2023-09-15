package use_cases

import (
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

type GetRecentQuestionsUseCaseInput struct {
	Page int64
}

type GetRecentQuestionsUseCaseInterface interface {
	Execute(input GetRecentQuestionsUseCaseInput) (enterprise.Question, error)
}

type DefaultGetRecentQuestionsUseCase struct {
	QuestionRepository repositories.QuestionRepositoryInterface
}

func NewDefaulGetRecentQuestionsUseCase(questionRepository repositories.QuestionRepositoryInterface) *DefaultGetRecentQuestionsUseCase {
	return &DefaultGetRecentQuestionsUseCase{
		QuestionRepository: questionRepository,
	}
}

func (uc *DefaultGetRecentQuestionsUseCase) Execute(input GetRecentQuestionsUseCaseInput) ([]enterprise.Question, error) {
	questions, err := uc.QuestionRepository.GetManyRecent(input.Page)

	if err != nil {
		return []enterprise.Question{}, err
	}

	return questions, nil
}
