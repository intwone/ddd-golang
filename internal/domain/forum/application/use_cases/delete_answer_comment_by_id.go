package use_cases

import (
	"errors"

	"github.com/intwone/ddd-golang/internal/constants"
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
)

type DeleteAnswerCommentByIDUseCaseInput struct {
	ID       string
	AuthorID string
}

type DeleteAnswerCommentByIDUseCaseInterface interface {
	Execute(input DeleteAnswerCommentByIDUseCaseInput) error
}

type DefaultDeleteAnswerCommentByIDUseCase struct {
	AnswerCommentRepository repositories.AnswerCommentsRepositoryInterface
}

func NewDefaultDeleteAnswerCommentByIDUseCase(AnswerCommentRepository repositories.AnswerCommentsRepositoryInterface) *DefaultDeleteAnswerCommentByIDUseCase {
	return &DefaultDeleteAnswerCommentByIDUseCase{
		AnswerCommentRepository: AnswerCommentRepository,
	}
}

func (uc *DefaultDeleteAnswerCommentByIDUseCase) Execute(input DeleteAnswerCommentByIDUseCaseInput) error {
	answerComment, err := uc.AnswerCommentRepository.GetByID(input.ID)

	if err != nil {
		return err
	}

	if input.AuthorID != answerComment.GetAuthorID() {
		return errors.New(constants.NotAllowedError)
	}

	uc.AnswerCommentRepository.DeleteByID(input.ID)

	return nil
}
