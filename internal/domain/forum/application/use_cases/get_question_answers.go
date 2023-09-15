package use_cases

import (
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

type GetQuestionAnswersUseCaseInput struct {
	Page       int64
	QuestionID string
}

type GetQuestionAnswersUseCaseInterface interface {
	Execute(input GetQuestionAnswersUseCaseInput) (enterprise.Answer, error)
}

type DefaultGetQuestionAnswersUseCase struct {
	AnswerRepository repositories.AnswerRepositoryInterface
}

func NewDefaulGetQuestionAnswersUseCase(answerRepository repositories.AnswerRepositoryInterface) *DefaultGetQuestionAnswersUseCase {
	return &DefaultGetQuestionAnswersUseCase{
		AnswerRepository: answerRepository,
	}
}

func (uc *DefaultGetQuestionAnswersUseCase) Execute(input GetQuestionAnswersUseCaseInput) ([]enterprise.Answer, error) {
	answers, err := uc.AnswerRepository.GetManyByQuestionID(input.Page, input.QuestionID)

	if err != nil {
		return []enterprise.Answer{}, err
	}

	return answers, nil
}
