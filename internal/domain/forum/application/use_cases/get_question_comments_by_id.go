package use_cases

import (
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

type GetQuestionCommentsByIDUseCaseInput struct {
	ID   string
	Page int64
}

type GetQuestionCommentsByIDUseCaseInterface interface {
	Execute(input GetQuestionCommentsByIDUseCaseInput) ([]enterprise.QuestionComment, error)
}

type DefaultGetQuestionCommentsByIDUseCase struct {
	QuestionCommentsRepository repositories.QuestionCommentsRepositoryInterface
}

func NewDefaulGetQuestionCommentsByIDUseCase(questionCommentsRepository repositories.QuestionCommentsRepositoryInterface) *DefaultGetQuestionCommentsByIDUseCase {
	return &DefaultGetQuestionCommentsByIDUseCase{
		QuestionCommentsRepository: questionCommentsRepository,
	}
}

func (uc *DefaultGetQuestionCommentsByIDUseCase) Execute(input GetQuestionCommentsByIDUseCaseInput) (*[]enterprise.QuestionComment, error) {
	questionComments, err := uc.QuestionCommentsRepository.GetManyByID(input.Page, input.ID)

	if err != nil {
		return nil, err
	}

	return questionComments, nil
}
