package use_cases

import (
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

type CreateQuestionUseCaseInput struct {
	AuthorID string
	Title    string
	Content  string
}

type CreateQuestionUseCaseInterface interface {
	Execute(input CreateQuestionUseCaseInput) (enterprise.Question, error)
}

type DefaultCreateQuestionUseCase struct {
	QuestionRepository repositories.QuestionRepositoryInterface
}

func NewDefaultCreateQuestionUseCase(questionRepository repositories.QuestionRepositoryInterface) *DefaultCreateQuestionUseCase {
	return &DefaultCreateQuestionUseCase{
		QuestionRepository: questionRepository,
	}
}

func (uc *DefaultCreateQuestionUseCase) Execute(input CreateQuestionUseCaseInput) (enterprise.Question, error) {
	newQuestion := enterprise.NewQuestion(input.Title, input.Content, input.AuthorID)

	err := uc.QuestionRepository.Create(newQuestion)

	if err != nil {
		return enterprise.Question{}, err
	}

	return *newQuestion, nil
}
