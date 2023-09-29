package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/intwone/ddd-golang/internal/constants"
	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	"github.com/intwone/ddd-golang/internal/presentation/dtos"
	er "github.com/intwone/ddd-golang/internal/presentation/errors"
	"github.com/intwone/ddd-golang/internal/presentation/mappers"
)

type DefaultGetQuestionBySlugControllerInterface struct {
	GetQuestionBySlugUseCase uc.GetQuestionBySlugUseCaseInterface
}

func NewDefaultGetQuestionBySlugController(getQuestionBySlugUseCase uc.GetQuestionBySlugUseCaseInterface) *DefaultGetQuestionBySlugControllerInterface {
	return &DefaultGetQuestionBySlugControllerInterface{
		GetQuestionBySlugUseCase: getQuestionBySlugUseCase,
	}
}

func (cqc *DefaultGetQuestionBySlugControllerInterface) Handle(c *gin.Context) {
	slug := c.Param("slug")

	question, err := cqc.GetQuestionBySlugUseCase.Execute(uc.GetQuestionBySlugUseCaseInput{Slug: slug})

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

	questionMapped := mappers.QuestionDTOMapper(question)

	c.JSON(http.StatusOK, dtos.ResponseDTO{"data": questionMapped})
}
