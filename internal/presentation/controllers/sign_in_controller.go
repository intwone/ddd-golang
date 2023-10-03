package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/intwone/ddd-golang/internal/constants"
	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	"github.com/intwone/ddd-golang/internal/presentation/dtos"
	er "github.com/intwone/ddd-golang/internal/presentation/errors"
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

func (sic *DefaultSignInControllerInterface) Handle(c *gin.Context) {
	var SignInRequestDTO dtos.SignInRequestDTO

	jsonBindErr := c.ShouldBindJSON(&SignInRequestDTO)

	if jsonBindErr != nil {
		restError := validations.ErrorValidation(jsonBindErr)

		c.JSON(restError.Code, restError)
		return
	}

	token, useCaseErrs := sic.AuthenticateUseCase.Execute(uc.AuthenticateUseCaseInput{
		Email:    SignInRequestDTO.Email,
		Password: SignInRequestDTO.Password,
	})

	if len(useCaseErrs) > 0 {
		causes := handleSignInErrorCauses(useCaseErrs)

		if len(causes) > 0 {
			restErr := er.NewBadRequestError(constants.OccurredSameErrorsError, causes)
			c.JSON(restErr.Code, restErr)
			return
		}

		restErr := er.NewInternalServerError(constants.UnexpectedError)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, dtos.ResponseDTO{"token": token})
}

func handleSignInErrorCauses(errs []error) []er.Cause {
	var causes = []er.Cause{}

	for _, err := range errs {
		switch err.Error() {
		case constants.InvalidEmailError:
			cause := er.Cause{Field: "email", Message: constants.InvalidEmailError}
			causes = append(causes, cause)

		case constants.NotContainMinimumCaracteresPasswordError:
			cause := er.Cause{Field: "password", Message: constants.NotContainMinimumCaracteresPasswordError}
			causes = append(causes, cause)

		case constants.NotContainUpperCaseCharacterePasswordError:
			cause := er.Cause{Field: "password", Message: constants.NotContainUpperCaseCharacterePasswordError}
			causes = append(causes, cause)

		case constants.NotContainSpecialCharacterePasswordError:
			cause := er.Cause{Field: "password", Message: constants.NotContainSpecialCharacterePasswordError}
			causes = append(causes, cause)

		case constants.NoRowsFound:
			cause := er.Cause{Field: "email/password", Message: constants.EmailOrPasswordIncorrectError}
			causes = append(causes, cause)

		case constants.PasswordAreNotTheSame:
			cause := er.Cause{Field: "email/password", Message: constants.EmailOrPasswordIncorrectError}
			causes = append(causes, cause)
		}
	}

	return causes
}
