package use_cases

import (
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

type GetAnswerCommentsByIDUseCaseInput struct {
	Page int64
	ID   string
}

type GetAnswerCommentsByIDUseCaseInterface interface {
	Execute(input GetAnswerCommentsByIDUseCaseInput) ([]enterprise.AnswerComment, error)
}

type DefaultGetAnswerCommentsByIDUseCase struct {
	AnswerCommentsRepository repositories.AnswerCommentsRepositoryInterface
}

func NewDefaulGetAnswerCommentsByIDUseCase(answerCommentsRepository repositories.AnswerCommentsRepositoryInterface) *DefaultGetAnswerCommentsByIDUseCase {
	return &DefaultGetAnswerCommentsByIDUseCase{
		AnswerCommentsRepository: answerCommentsRepository,
	}
}

func (uc *DefaultGetAnswerCommentsByIDUseCase) Execute(input GetAnswerCommentsByIDUseCaseInput) ([]enterprise.AnswerComment, error) {
	answerComments, err := uc.AnswerCommentsRepository.GetManyByID(input.Page, input.ID)

	if err != nil {
		return []enterprise.AnswerComment{}, err
	}

	return answerComments, nil
}
