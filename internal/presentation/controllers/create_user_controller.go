package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	"github.com/intwone/ddd-golang/internal/presentation/dtos"
	"github.com/intwone/ddd-golang/internal/presentation/errors"
)

type DefaultCreateUserControllerInterface struct {
	CreateUserUseCase uc.CreateUserUseCaseInterface
}

func NewDefaultCreateUserController(createUserUseCase uc.CreateUserUseCaseInterface) *DefaultCreateUserControllerInterface {
	return &DefaultCreateUserControllerInterface{
		CreateUserUseCase: createUserUseCase,
	}
}

func (cuc *DefaultCreateUserControllerInterface) Handle(c *gin.Context) {
	var createUserRequestDTO dtos.CreateUserRequestDTO

	jsonBindErr := c.ShouldBindJSON(&createUserRequestDTO)

	if jsonBindErr != nil {
		restError := errors.ValidateError(jsonBindErr)

		c.JSON(restError.Code, restError)
		return
	}

	if !createUserRequestDTO.Role.Validate() {
		restErr := errors.NewBadRequestError("role should be student or instructor")
		c.JSON(restErr.Code, restErr)
		return
	}

	_, useCaseErr := cuc.CreateUserUseCase.Execute(uc.CreateUserUseCaseInput{
		Name: createUserRequestDTO.Name,
		Role: string(createUserRequestDTO.Role),
	})

	if useCaseErr != nil {
		c.JSON(http.StatusInternalServerError, useCaseErr)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
