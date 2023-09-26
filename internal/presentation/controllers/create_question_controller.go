package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	"github.com/intwone/ddd-golang/internal/presentation/dtos"
	"github.com/intwone/ddd-golang/internal/presentation/errors"
)

type DefaultCreateQuestionControllerInterface struct {
	CreateQuestionUseCase uc.CreateQuestionUseCaseInterface
}

func NewDefaultCreateQuestionController(createQuestionUseCase uc.CreateQuestionUseCaseInterface) *DefaultCreateQuestionControllerInterface {
	return &DefaultCreateQuestionControllerInterface{
		CreateQuestionUseCase: createQuestionUseCase,
	}
}

func (cqc *DefaultCreateQuestionControllerInterface) Handle(c *gin.Context) {
	var questionRequestDTO dtos.CreateQuestionRequestDTO

	jsonBindErr := c.ShouldBindJSON(&questionRequestDTO)

	if jsonBindErr != nil {
		restError := errors.ValidateError(jsonBindErr)

		c.JSON(restError.Code, restError)
		return
	}

	_, useCaseErr := cqc.CreateQuestionUseCase.Execute(uc.CreateQuestionUseCaseInput{
		AuthorID: questionRequestDTO.AuthorID,
		Title:    questionRequestDTO.Title,
		Content:  questionRequestDTO.Content,
	})

	if useCaseErr != nil {
		c.JSON(http.StatusInternalServerError, useCaseErr)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
