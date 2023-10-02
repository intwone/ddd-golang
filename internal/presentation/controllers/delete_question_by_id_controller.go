package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/intwone/ddd-golang/internal/constants"
	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	er "github.com/intwone/ddd-golang/internal/presentation/errors"
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
	questionID := c.Param("questionID")

	userID := c.MustGet("userID")

	if _, err := uuid.Parse(questionID); err != nil {
		cause := er.NewCause("id", constants.InvalidUUIDError)
		restErr := er.NewBadRequestError(constants.OccurredSameErrorsError, []er.Cause{*cause})
		c.JSON(restErr.Code, restErr)
		return
	}

	err := dqc.DeleteQuestionByIDUseCase.Execute(uc.DeleteQuestionByIDUseCaseInput{ID: questionID, AuthorID: userID.(string)})

	if err != nil {
		if strings.Contains(err.Error(), constants.NoRowsFound) {
			restErr := er.NewNotFoundError(constants.QuestionNotFoundError)
			c.JSON(restErr.Code, restErr)
			return
		}

		if strings.Contains(err.Error(), constants.NotAllowedError) {
			restErr := er.NewUnauthorizedError(constants.NotAllowedError)
			c.JSON(restErr.Code, restErr)
			return
		}

		restErr := er.NewInternalServerError(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
