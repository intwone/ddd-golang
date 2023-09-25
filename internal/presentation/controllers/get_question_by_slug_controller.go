package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	"github.com/intwone/ddd-golang/internal/presentation/mappers"
)

type GetQuestionBySlugInterface interface {
	Handle(c *gin.Context)
}

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
		c.JSON(http.StatusInternalServerError, err)
	}

	questionMapped := mappers.QuestionDTOMapper(question)

	c.JSON(http.StatusOK, questionMapped)
}
