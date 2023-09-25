package use_cases

import (
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

type CommentOnAnswerUseCaseInput struct {
	AuthorID string
	AnswerID string
	Content  string
}

type CommentOnAnswerUseCaseInterface interface {
	Execute(input CommentOnAnswerUseCaseInput) (enterprise.QuestionComment, error)
}

type DefaultCommentOnAnswerUseCase struct {
	AnswerRepository        repositories.AnswerRepositoryInterface
	AnswerCommentRepository repositories.AnswerCommentsRepositoryInterface
}

func NewDefaultCommentOnAnswerUseCase(answerRepository repositories.AnswerRepositoryInterface, answerComentsRepository repositories.AnswerCommentsRepositoryInterface) *DefaultCommentOnAnswerUseCase {
	return &DefaultCommentOnAnswerUseCase{
		AnswerRepository:        answerRepository,
		AnswerCommentRepository: answerComentsRepository,
	}
}

func (uc *DefaultCommentOnAnswerUseCase) Execute(input CommentOnAnswerUseCaseInput) (*enterprise.AnswerComment, error) {
	_, err := uc.AnswerRepository.GetByID(input.AnswerID)

	if err != nil {
		return nil, err
	}

	answerComment := enterprise.NewAnswerComment(input.Content, input.AuthorID, input.AnswerID)

	createErr := uc.AnswerCommentRepository.Create(answerComment)

	if createErr != nil {
		return nil, createErr
	}

	return answerComment, nil
}
