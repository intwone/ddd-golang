package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	"github.com/intwone/ddd-golang/internal/presentation/errors"
	"github.com/intwone/ddd-golang/internal/presentation/mappers"
)

type DefaultGetQuestionBySlugControllerInterface struct {
	GetQuestionBySlugUseCase uc.GetQuestionBySlugUseCaseInterface
}

func NewDefaultGetQuestionBySlug(getQuestionBySlugUseCase uc.GetQuestionBySlugUseCaseInterface) *DefaultGetQuestionBySlugControllerInterface {
	return &DefaultGetQuestionBySlugControllerInterface{
		GetQuestionBySlugUseCase: getQuestionBySlugUseCase,
	}
}

func (cqc *DefaultGetQuestionBySlugControllerInterface) Handle(c *gin.Context) {
	slug := c.Param("slug")

	question, err := cqc.GetQuestionBySlugUseCase.Execute(uc.GetQuestionBySlugUseCaseInput{Slug: slug})

	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			restErr := errors.NewNotFoundError("question not found")
			c.JSON(restErr.Code, restErr)
			return
		}

		c.JSON(http.StatusInternalServerError, err)
		return
	}

	questionMapped := mappers.QuestionDTOMapper(question)

	c.JSON(http.StatusOK, questionMapped)
}
