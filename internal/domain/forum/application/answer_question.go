package application

import (
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

type AnswerQuestionUseCaseInput struct {
	InstructorID string
	QuestionID   string
	Content      string
}

type AnswerQuestionUseCaseInterface interface {
	Execute(input AnswerQuestionUseCaseInput) (enterprise.Answer, error)
}

type DefaultAnswerQuestionUseCase struct {
	AnswersRepository repositories.RepositoryInterface
}

func NewDefaultAnswerQuestionUseCase(answersRepository repositories.RepositoryInterface) *DefaultAnswerQuestionUseCase {
	return &DefaultAnswerQuestionUseCase{
		AnswersRepository: answersRepository,
	}
}

func (uc *DefaultAnswerQuestionUseCase) Execute(input AnswerQuestionUseCaseInput) (enterprise.Answer, error) {
	newAnswer := enterprise.NewAnswer(input.Content, input.InstructorID, input.QuestionID)

	uc.AnswersRepository.Create(newAnswer)

	return *newAnswer, nil
}
