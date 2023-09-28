package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	"github.com/intwone/ddd-golang/internal/presentation/dtos"
	"github.com/intwone/ddd-golang/internal/presentation/errors"
	"github.com/intwone/ddd-golang/internal/presentation/validations"
)

type DefaultSignUpControllerInterface struct {
	CreateUserUseCase uc.CreateUserUseCaseInterface
}

func NewDefaultSignUpController(createUserUseCase uc.CreateUserUseCaseInterface) *DefaultSignUpControllerInterface {
	return &DefaultSignUpControllerInterface{
		CreateUserUseCase: createUserUseCase,
	}
}

func (cuc *DefaultSignUpControllerInterface) Handle(c *gin.Context) {
	var signUpRequestDTO dtos.SignUpRequestDTO

	jsonBindErr := c.ShouldBindJSON(&signUpRequestDTO)

	if jsonBindErr != nil {
		restError := validations.ErrorValidation(jsonBindErr)

		c.JSON(restError.Code, restError)
		return
	}

	if !signUpRequestDTO.Role.RoleValidation() {
		restErr := errors.NewBadRequestError("role must be student or instructor")
		c.JSON(restErr.Code, restErr)
		return
	}

	// finduserbyemail

	_, useCaseErr := cuc.CreateUserUseCase.Execute(uc.CreateUserUseCaseInput{
		Name:     signUpRequestDTO.Name,
		Email:    signUpRequestDTO.Email,
		Password: signUpRequestDTO.Password,
		Role:     string(signUpRequestDTO.Role),
	})

	if useCaseErr != nil {
		c.JSON(http.StatusInternalServerError, useCaseErr)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
