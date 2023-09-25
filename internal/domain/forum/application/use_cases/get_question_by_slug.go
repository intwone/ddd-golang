package use_cases

import (
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

type GetQuestionBySlugUseCaseInput struct {
	Slug string
}

type GetQuestionBySlugUseCaseInterface interface {
	Execute(input GetQuestionBySlugUseCaseInput) (*enterprise.Question, error)
}

type DefaultGetQuestionBySlugUseCase struct {
	QuestionRepository repositories.QuestionRepositoryInterface
}

func NewDefaulGetQuestionBySlugUseCase(questionRepository repositories.QuestionRepositoryInterface) *DefaultGetQuestionBySlugUseCase {
	return &DefaultGetQuestionBySlugUseCase{
		QuestionRepository: questionRepository,
	}
}

func (uc *DefaultGetQuestionBySlugUseCase) Execute(input GetQuestionBySlugUseCaseInput) (*enterprise.Question, error) {
	question, err := uc.QuestionRepository.GetBySlug(input.Slug)

	if err != nil {
		return nil, err
	}

	return question, nil
}
