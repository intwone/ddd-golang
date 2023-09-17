package use_cases

import (
	"errors"

	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
)

type DeleteQuestionCommentByIDUseCaseInput struct {
	ID       string
	AuthorID string
}

type DeleteQuestionCommentByIDUseCaseInterface interface {
	Execute(input DeleteQuestionCommentByIDUseCaseInput) error
}

type DefaultDeleteQuestionCommentByIDUseCase struct {
	QuestionCommentRepository repositories.QuestionCommentsRepositoryInterface
}

func NewDefaultDeleteQuestionCommentByIDUseCase(questionCommentRepository repositories.QuestionCommentsRepositoryInterface) *DefaultDeleteQuestionCommentByIDUseCase {
	return &DefaultDeleteQuestionCommentByIDUseCase{
		QuestionCommentRepository: questionCommentRepository,
	}
}

func (uc *DefaultDeleteQuestionCommentByIDUseCase) Execute(input DeleteQuestionCommentByIDUseCaseInput) error {
	questionComment, err := uc.QuestionCommentRepository.GetByID(input.ID)

	if err != nil {
		return err
	}

	if input.AuthorID != *questionComment.GetAuthorID().Value {
		return errors.New("not allowed")
	}

	uc.QuestionCommentRepository.DeleteByID(input.ID)

	return nil
}
