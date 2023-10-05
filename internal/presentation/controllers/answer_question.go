package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	"github.com/intwone/ddd-golang/internal/presentation/dtos"
	er "github.com/intwone/ddd-golang/internal/presentation/errors"
	"github.com/intwone/ddd-golang/internal/presentation/validations"
)

type DefaultAnswerQuestionControllerInterface struct {
	AnswerQuestionUseCase uc.AnswerQuestionUseCaseInterface
}

func NewDefaultAnswerQuestionController(answerQuestionUseCase uc.AnswerQuestionUseCaseInterface) *DefaultAnswerQuestionControllerInterface {
	return &DefaultAnswerQuestionControllerInterface{
		AnswerQuestionUseCase: answerQuestionUseCase,
	}
}

func (aqc *DefaultAnswerQuestionControllerInterface) Handle(c *gin.Context) {
	var answerQuestionRequestDTO dtos.AnswerQuestionRequestDTO

	jsonBindErr := c.ShouldBindJSON(&answerQuestionRequestDTO)

	if jsonBindErr != nil {
		restError := validations.ErrorValidation(jsonBindErr)

		c.JSON(restError.Code, restError)
		return
	}

	_, err := aqc.AnswerQuestionUseCase.Execute(uc.AnswerQuestionUseCaseInput{
		AuthorID:   answerQuestionRequestDTO.AuthorID,
		QuestionID: answerQuestionRequestDTO.QuestionID,
		Content:    answerQuestionRequestDTO.Content,
	})

	if err != nil {
		restErr := er.NewInternalServerError(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusCreated, nil)
}
