package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	re "github.com/intwone/ddd-golang/internal/presentation/errors"
	"github.com/intwone/ddd-golang/internal/presentation/mappers"
)

type DefaultGetQuestionBySlugInterface struct {
	GetQuestionBySlugUseCase uc.GetQuestionBySlugUseCaseInterface
}

func NewDefaultGetQuestionBySlug(getQuestionBySlugUseCase uc.GetQuestionBySlugUseCaseInterface) *DefaultGetQuestionBySlugInterface {
	return &DefaultGetQuestionBySlugInterface{
		GetQuestionBySlugUseCase: getQuestionBySlugUseCase,
	}
}

func (cqc *DefaultGetQuestionBySlugInterface) Handle(c *gin.Context) {
	slug := c.Param("slug")

	question, err := cqc.GetQuestionBySlugUseCase.Execute(uc.GetQuestionBySlugUseCaseInput{Slug: slug})

	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			restErr := re.NewNotFoundError("question not found")
			c.JSON(restErr.Code, restErr)
			return
		}

		c.JSON(http.StatusInternalServerError, err)
		return
	}

	questionMapped := mappers.QuestionDTOMapper(question)

	c.JSON(http.StatusOK, questionMapped)
}
