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
		cause := er.NewCause("email", constants.EmailAlreadyTakenError)
		restErr := er.NewConflictError(constants.InvalidFieldsError, []er.Cause{*cause})
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
		causes := handleSignInErrorCauses(createUserUseCaseErrs)

		if len(causes) > 0 {
			restErr := er.NewBadRequestError(constants.OccurredSameErrorsError, causes)
			c.JSON(restErr.Code, restErr)
			return
		}

		restErr := er.NewInternalServerError(constants.UnexpectedError)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func handleSignUpErrorCauses(errs []error) []er.Cause {
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

		case constants.InvalidRoleError:
			cause := er.Cause{Field: "role", Message: constants.InvalidRoleError}
			causes = append(causes, cause)
		}
	}

	return causes
}
