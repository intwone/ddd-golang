package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	"github.com/intwone/ddd-golang/internal/presentation/errors"
)

type DefaultDeleteQuestionByIDControllerInterface struct {
	DeleteQuestionByIDUseCase uc.DeleteQuestionByIDUseCaseInterface
}

func NewDefaultDeleteQuestionByIDController(deleteQuestionByIDUseCase uc.DeleteQuestionByIDUseCaseInterface) *DefaultDeleteQuestionByIDControllerInterface {
	return &DefaultDeleteQuestionByIDControllerInterface{
		DeleteQuestionByIDUseCase: deleteQuestionByIDUseCase,
	}
}

func (dqc *DefaultDeleteQuestionByIDControllerInterface) Handle(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		restErr := errors.NewBadRequestError("id must be a valid UUID")
		c.JSON(restErr.Code, restErr)
		return
	}

	// TODO: Pegar o userID através do header para colocar no input.AuthorID do usecase
	err := dqc.DeleteQuestionByIDUseCase.Execute(uc.DeleteQuestionByIDUseCaseInput{ID: id})

	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			restErr := errors.NewNotFoundError("question not found")
			c.JSON(restErr.Code, restErr)
			return
		}

		if strings.Contains(err.Error(), "not allowed") {
			restErr := errors.NewUnauthorizedError("you cant delete this question")
			c.JSON(restErr.Code, restErr)
			return
		}

		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
