package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/intwone/ddd-golang/internal/constants"
	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	"github.com/intwone/ddd-golang/internal/presentation/dtos"
	"github.com/intwone/ddd-golang/internal/presentation/errors"
	"github.com/intwone/ddd-golang/internal/presentation/validations"
)

type DefaultSignUpControllerInterface struct {
	CreateUserUseCase     uc.CreateUserUseCaseInterface
	GetUserByEmailUseCase uc.GetUserByEmailUseCaseInterface
}

func NewDefaultSignUpController(createUserUseCase uc.CreateUserUseCaseInterface, getUserByEmailUseCase uc.GetUserByEmailUseCaseInterface) *DefaultSignUpControllerInterface {
	return &DefaultSignUpControllerInterface{
		CreateUserUseCase:     createUserUseCase,
		GetUserByEmailUseCase: getUserByEmailUseCase,
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
		restErr := errors.NewBadRequestError(constants.InvalidRoleError)
		c.JSON(restErr.Code, restErr)
		return
	}

	user, _ := cuc.GetUserByEmailUseCase.Execute(uc.GetUserByEmailUseCaseInput{
		Email: signUpRequestDTO.Email,
	})

	if user != nil {
		restErr := errors.NewConflictError(constants.EmailAlreadyTakenError)
		c.JSON(http.StatusConflict, restErr)
		return
	}

	_, createUserCaseErr := cuc.CreateUserUseCase.Execute(uc.CreateUserUseCaseInput{
		Name:     signUpRequestDTO.Name,
		Email:    signUpRequestDTO.Email,
		Password: signUpRequestDTO.Password,
		Role:     string(signUpRequestDTO.Role),
	})

	if createUserCaseErr != nil {
		if strings.Contains(createUserCaseErr.Error(), constants.InvalidEmailError) {
			restErr := errors.NewBadRequestError(constants.InvalidEmailError)
			c.JSON(restErr.Code, restErr)
			return
		}

		messageErrors := [...]string{
			constants.NotContainMinimumCaracteresPasswordError,
			constants.NotContainUpperCaseCharacterePasswordError,
			constants.NotContainSpecialCharacterePasswordError,
		}

		e := messageErrors[:]
		messageErr := strings.Join(e, ",")

		if strings.Contains(messageErr, createUserCaseErr.Error()) {
			restErr := errors.NewBadRequestError(createUserCaseErr.Error())
			c.JSON(restErr.Code, restErr)
			return
		}

		c.JSON(http.StatusInternalServerError, createUserCaseErr)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
