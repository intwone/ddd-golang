package validations

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/intwone/ddd-golang/internal/constants"
	er "github.com/intwone/ddd-golang/internal/presentation/errors"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	val, ok := binding.Validator.Engine().(*validator.Validate)

	if ok {
		en := en.New()
		unicode := ut.New(en, en)
		transl, _ = unicode.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ErrorValidation(validationError error) *er.RestError {
	var jsonError *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationError, &jsonError) {
		causes := []er.Cause{
			{Field: "", Message: constants.UnMarshalJSONError},
		}

		return er.NewBadRequestError(constants.InvalidFieldTypeError, causes)
	}

	if errors.As(validationError, &jsonValidationError) {
		causes := []er.Cause{}

		for _, e := range validationError.(validator.ValidationErrors) {
			cause := er.Cause{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			causes = append(causes, cause)
		}

		return er.NewBadRequestError(constants.InvalidFieldsError, causes)
	}

	causes := []er.Cause{
		{Field: "", Message: constants.ConvertFieldsError},
	}

	return er.NewBadRequestError(constants.InvalidFieldsError, causes)
}
