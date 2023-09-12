package usecases

import (
	"github.com/intwone/ddd-golang/internal/domain/entities"
	"github.com/intwone/ddd-golang/internal/domain/repositories"
)

type AnswerQuestionUseCaseInput struct {
	InstructorID string
	QuestionID   string
	Content      string
}

type AnswerQuestionUseCaseInterface interface {
	Execute(input AnswerQuestionUseCaseInput) (entities.Answer, error)
}

type DefaultAnswerQuestionUseCase struct {
	AnswersRepository repositories.RepositoryInterface
}

func NewDefaultAnswerQuestionUseCase(answersRepository repositories.RepositoryInterface) *DefaultAnswerQuestionUseCase {
	return &DefaultAnswerQuestionUseCase{
		AnswersRepository: answersRepository,
	}
}

func (uc *DefaultAnswerQuestionUseCase) Execute(input AnswerQuestionUseCaseInput) (entities.Answer, error) {
	newAnswer := entities.NewAnswer(input.Content, input.InstructorID, input.QuestionID)

	uc.AnswersRepository.Create(*newAnswer)

	return *newAnswer, nil
}
