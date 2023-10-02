package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/intwone/ddd-golang/internal/constants"
	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	"github.com/intwone/ddd-golang/internal/presentation/dtos"
	er "github.com/intwone/ddd-golang/internal/presentation/errors"
	"github.com/intwone/ddd-golang/internal/presentation/mappers"
)

type DefaultGetRecentQuestionsControllerInterface struct {
	GetRecentQuestionsUseCase uc.GetRecentQuestionsUseCaseInterface
}

func NewDefaultGetRecentQuestionsController(getRecentQuestionsUseCase uc.GetRecentQuestionsUseCaseInterface) *DefaultGetRecentQuestionsControllerInterface {
	return &DefaultGetRecentQuestionsControllerInterface{
		GetRecentQuestionsUseCase: getRecentQuestionsUseCase,
	}
}

func (cqc *DefaultGetRecentQuestionsControllerInterface) Handle(c *gin.Context) {
	page, err := strconv.ParseInt(c.Query("page"), 10, 64)

	if err != nil {
		cause := er.NewCause("page", constants.InvalidQueryParam)
		restErr := er.NewBadRequestError(constants.InvalidQueryParam, []er.Cause{*cause})
		c.JSON(restErr.Code, restErr)
		return
	}

	offset := (page - 1) * 20

	questions, err := cqc.GetRecentQuestionsUseCase.Execute(uc.GetRecentQuestionsUseCaseInput{Page: offset})

	if err != nil {
		if strings.Contains(err.Error(), constants.NoRowsFound) {
			restErr := er.NewNotFoundError(constants.QuestionNotFoundError)
			c.JSON(restErr.Code, restErr)
			return
		}

		restErr := er.NewInternalServerError(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	questionMapped := mappers.QuestionsDTOMapper(*questions)

	if len(questionMapped) == 0 {
		c.JSON(http.StatusOK, dtos.ResponseDTO{"data": []dtos.QuestionDTO{}})
		return
	}

	c.JSON(http.StatusOK, dtos.ResponseDTO{"data": questionMapped})
	return
}
