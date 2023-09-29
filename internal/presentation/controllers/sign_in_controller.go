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

type DefaultSignInControllerInterface struct {
	AuthenticateUseCase uc.AuthenticateUseCaseInterface
}

func NewDefaultSignInController(authenticateUseCase uc.AuthenticateUseCaseInterface) *DefaultSignInControllerInterface {
	return &DefaultSignInControllerInterface{
		AuthenticateUseCase: authenticateUseCase,
	}
}

func (ac *DefaultSignInControllerInterface) Handle(c *gin.Context) {
	var SignInRequestDTO dtos.SignInRequestDTO

	jsonBindErr := c.ShouldBindJSON(&SignInRequestDTO)

	if jsonBindErr != nil {
		restError := validations.ErrorValidation(jsonBindErr)

		c.JSON(restError.Code, restError)
		return
	}

	token, useCaseErrs := ac.AuthenticateUseCase.Execute(uc.AuthenticateUseCaseInput{
		Email:    SignInRequestDTO.Email,
		Password: SignInRequestDTO.Password,
	})

	if len(useCaseErrs) > 0 {
		var causes = []errors.Cause{}

		for _, err := range useCaseErrs {
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
			case constants.NoRowsFound:
				cause := errors.Cause{Field: "email/password", Message: constants.EmailOrPasswordIncorrectError}
				causes = append(causes, cause)
			case constants.PasswordAreNotTheSame:
				cause := errors.Cause{Field: "email/password", Message: constants.EmailOrPasswordIncorrectError}
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

	c.JSON(http.StatusOK, dtos.ResponseDTO{"token": token})
}
