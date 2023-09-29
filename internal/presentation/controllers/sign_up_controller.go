package controllers

import (
	"net/http"

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

	user, _ := cuc.GetUserByEmailUseCase.Execute(uc.GetUserByEmailUseCaseInput{
		Email: signUpRequestDTO.Email,
	})

	if user != nil {
		causes := []errors.Cause{
			{Field: "email", Message: constants.EmailAlreadyTakenError},
		}

		restErr := errors.NewConflictError(constants.InvalidFieldsError, causes)
		c.JSON(http.StatusConflict, restErr)
		return
	}

	_, createUserUseCaseErrs := cuc.CreateUserUseCase.Execute(uc.CreateUserUseCaseInput{
		Name:     signUpRequestDTO.Name,
		Email:    signUpRequestDTO.Email,
		Password: signUpRequestDTO.Password,
		Role:     signUpRequestDTO.Role,
	})

	if len(createUserUseCaseErrs) > 0 {
		var causes = []errors.Cause{}

		for _, err := range createUserUseCaseErrs {
			switch err.Error() {
			case constants.InvalidEmailError:
				cause := errors.Cause{Field: "email", Message: constants.InvalidEmailError}
				causes = append(causes, cause)
			case constants.NotContainMinimumCaracteresPasswordError:
				cause := errors.Cause{Field: "password", Message: constants.NotContainMinimumCaracteresPasswordError}
				causes = append(causes, cause)
			case constants.NotContainUpperCaseCharacterePasswordError:
				cause := errors.Cause{Field: "password", Message: constants.NotContainUpperCaseCharacterePasswordError}
				causes = append(causes, cause)
			case constants.NotContainSpecialCharacterePasswordError:
				cause := errors.Cause{Field: "password", Message: constants.NotContainSpecialCharacterePasswordError}
				causes = append(causes, cause)
			case constants.InvalidRoleError:
				cause := errors.Cause{Field: "role", Message: constants.InvalidRoleError}
				causes = append(causes, cause)
			}
		}

		if len(causes) > 0 {
			restErr := errors.NewBadRequestError(constants.OccurredSameErrorsError, causes)
			c.JSON(restErr.Code, restErr)
			return
		}

		restErr := errors.NewInternalServerError(constants.UnexpectedError)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
