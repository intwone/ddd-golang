package use_cases

import (
	"errors"

	"github.com/intwone/ddd-golang/internal/constants"
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

func (uc *DefaultChooseQuestionBestAnswerUseCase) Execute(input ChooseQuestionBestAnswerUseCaseInput) (*enterprise.Question, error) {
	answer, answerGetByIDErr := uc.AnswersRepository.GetByID(input.AnswerID)

	if answerGetByIDErr != nil {
		return nil, answerGetByIDErr
	}

	question, questionrGetByIDErr := uc.QuestionsRepository.GetByID(answer.GetQuestionID())

	if questionrGetByIDErr != nil {
		return nil, questionrGetByIDErr
	}

	if input.AuthorID != question.GetAuthorID() {
		return nil, errors.New(constants.NotAllowedError)
	}

	answerID := vo.NewUniqueID(question.GetAuthorID())

	question.SetBestAnswerID(*answerID)

	saveErr := uc.QuestionsRepository.Save(question)

	if saveErr != nil {
		return nil, saveErr
	}

	return question, nil
}
