package use_cases

import (
	"errors"

	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
)

type UpdateQuestionByIDUseCaseInput struct {
	ID       string
	AuthorID string
	Title    string
	Content  string
}

type UpdateQuestionByIDUseCaseInterface interface {
	Execute(input UpdateQuestionByIDUseCaseInput) error
}

type DefaultUpdateQuestionByIDUseCase struct {
	QuestionRepository repositories.QuestionRepositoryInterface
}

func NewDefaultUpdateQuestionByIDUseCase(questionRepository repositories.QuestionRepositoryInterface) *DefaultUpdateQuestionByIDUseCase {
	return &DefaultUpdateQuestionByIDUseCase{
		QuestionRepository: questionRepository,
	}
}

func (uc *DefaultUpdateQuestionByIDUseCase) Execute(input UpdateQuestionByIDUseCaseInput) error {
	question, err := uc.QuestionRepository.GetByID(input.ID)

	if err != nil {
		return err
	}

	if input.AuthorID != *question.GetAuthorID().Value {
		return errors.New("not allowed")
	}

	question.SetTitle(input.Title)
	question.SetContent(input.Content)

	uc.QuestionRepository.Save(&question)

	return nil
}
