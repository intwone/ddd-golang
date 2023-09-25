package use_cases

import (
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

type GetAnswerCommentsByIDUseCaseInput struct {
	ID   string
	Page int64
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

func (uc *DefaultGetAnswerCommentsByIDUseCase) Execute(input GetAnswerCommentsByIDUseCaseInput) (*[]enterprise.AnswerComment, error) {
	answerComments, err := uc.AnswerCommentsRepository.GetManyByID(input.Page, input.ID)

	if err != nil {
		return nil, err
	}

	return answerComments, nil
}
