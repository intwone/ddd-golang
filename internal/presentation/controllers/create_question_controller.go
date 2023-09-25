package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	"github.com/intwone/ddd-golang/internal/presentation/dtos"
	"github.com/intwone/ddd-golang/internal/presentation/errors"
)

type DefaultCreateQuestionInterface struct {
	CreateQuestionUseCase uc.CreateQuestionUseCaseInterface
}

func NewDefaultCreateQuestion(createQuestionUseCase uc.CreateQuestionUseCaseInterface) *DefaultCreateQuestionInterface {
	return &DefaultCreateQuestionInterface{
		CreateQuestionUseCase: createQuestionUseCase,
	}
}

func (cqc *DefaultCreateQuestionInterface) Handle(c *gin.Context) {
	var questionRequestDTO dtos.CreateQuestionRequestDTO

	jsonBindErr := c.ShouldBindJSON(questionRequestDTO)

	if jsonBindErr != nil {
		restError := errors.ValidateError(jsonBindErr)

		c.JSON(restError.Code, restError)
	}

	_, useCaseErr := cqc.CreateQuestionUseCase.Execute(uc.CreateQuestionUseCaseInput{Title: questionRequestDTO.Title, Content: questionRequestDTO.Content})

	if useCaseErr != nil {
		c.JSON(http.StatusInternalServerError, useCaseErr)
	}

	c.JSON(http.StatusNoContent, nil)
}
