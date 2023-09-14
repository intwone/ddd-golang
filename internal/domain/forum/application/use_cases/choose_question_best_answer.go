package use_cases

import (
	"errors"

	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
	vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"
)

type ChooseQuestionBestAnswerUseCaseInput struct {
	AnswerID string
	AuthorID string
}

type ChooseQuestionBestAnswerUseCaseInterface interface {
	Execute(input ChooseQuestionBestAnswerUseCaseInput) (enterprise.Question, error)
}

type DefaultChooseQuestionBestAnswerUseCase struct {
	QuestionsRepository repositories.QuestionRepositoryInterface
	AnswersRepository   repositories.AnswerRepositoryInterface
}

func NewDefaultChooseQuestionBestAnswerUseCase(questionRepository repositories.QuestionRepositoryInterface, answersRepository repositories.AnswerRepositoryInterface) *DefaultChooseQuestionBestAnswerUseCase {
	return &DefaultChooseQuestionBestAnswerUseCase{
		QuestionsRepository: questionRepository,
		AnswersRepository:   answersRepository,
	}
}

func (uc *DefaultChooseQuestionBestAnswerUseCase) Execute(input ChooseQuestionBestAnswerUseCaseInput) (enterprise.Question, error) {
	answer, answerGetByIDErr := uc.AnswersRepository.GetByID(input.AnswerID)

	if answerGetByIDErr != nil {
		return enterprise.Question{}, answerGetByIDErr
	}

	question, questionrGetByIDErr := uc.QuestionsRepository.GetByID(*answer.GetQuestionID().Value)

	if questionrGetByIDErr != nil {
		return enterprise.Question{}, questionrGetByIDErr
	}

	if input.AuthorID != *question.GetAuthorID().Value {
		return enterprise.Question{}, errors.New("not allowed")
	}

	answerID := vo.NewUniqueID(*question.GetAuthorID().Value)

	question.SetBestAnswerID(*answerID)

	saveErr := uc.QuestionsRepository.Save(question)

	if saveErr != nil {
		return enterprise.Question{}, saveErr
	}

	return *question, nil
}
