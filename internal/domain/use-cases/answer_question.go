package usecases

import "github.com/intwone/ddd-golang/internal/domain/entities"

type AnswerQuestionUseCaseInput struct {
	InstructorID string
	QuestionID   string
	Content      string
}

type AnswerQuestionUseCaseInterface interface {
	Execute(input AnswerQuestionUseCaseInput) (entities.Answer, error)
}

type DefaultAnswerQuestionUseCase struct{}

func (uc *DefaultAnswerQuestionUseCase) Execute(input AnswerQuestionUseCaseInput) (entities.Answer, error) {
	newAnswer := entities.NewAnswer(input.Content)

	return *newAnswer, nil
}
