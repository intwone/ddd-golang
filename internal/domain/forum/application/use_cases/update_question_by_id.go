package use_cases

import (
	"errors"

	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

type UpdateQuestionByIDUseCaseInput struct {
	ID       string
	AuthorID string
	Title    string
	Content  string
}

type UpdateQuestionByIDUseCaseInterface interface {
	Execute(input UpdateQuestionByIDUseCaseInput) (enterprise.Question, error)
}

type DefaultUpdateQuestionByIDUseCase struct {
	QuestionRepository repositories.QuestionRepositoryInterface
}

func NewDefaultUpdateQuestionByIDUseCase(questionRepository repositories.QuestionRepositoryInterface) *DefaultUpdateQuestionByIDUseCase {
	return &DefaultUpdateQuestionByIDUseCase{
		QuestionRepository: questionRepository,
	}
}

func (uc *DefaultUpdateQuestionByIDUseCase) Execute(input UpdateQuestionByIDUseCaseInput) (enterprise.Question, error) {
	question, err := uc.QuestionRepository.GetByID(input.ID)

	if err != nil {
		return enterprise.Question{}, err
	}

	if input.AuthorID != *question.GetAuthorID().Value {
		return enterprise.Question{}, errors.New("not allowed")
	}

	question.SetTitle(input.Title)
	question.SetContent(input.Content)

	err = uc.QuestionRepository.Save(&question)

	if err != nil {
		return enterprise.Question{}, err
	}

	return question, nil
}
