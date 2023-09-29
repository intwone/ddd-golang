package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/intwone/ddd-golang/internal/constants"
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
		causes := []errors.Cause{
			{Field: "id", Message: constants.InvalidUUIDError},
		}

		restErr := errors.NewBadRequestError(constants.OccurredSameErrorsError, causes)
		c.JSON(restErr.Code, restErr)
		return
	}

	// TODO: Pegar o userID através do header para colocar no input.AuthorID do usecase
	err := dqc.DeleteQuestionByIDUseCase.Execute(uc.DeleteQuestionByIDUseCaseInput{ID: id})

	if err != nil {
		if strings.Contains(err.Error(), constants.NoRowsFound) {
			restErr := errors.NewNotFoundError(constants.QuestionNotFoundError)
			c.JSON(restErr.Code, restErr)
			return
		}

		if strings.Contains(err.Error(), constants.NotAllowedError) {
			restErr := errors.NewUnauthorizedError(constants.NotAllowedError)
			c.JSON(restErr.Code, restErr)
			return
		}

		restErr := errors.NewInternalServerError(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
