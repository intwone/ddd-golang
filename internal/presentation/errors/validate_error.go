package errors

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
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

func ValidateError(validationError error) *RestError {
	var jsonError *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationError, &jsonError) {
		return NewBadRequestError("invalid field type")
	}

	if errors.As(validationError, &jsonValidationError) {
		errorsCauses := []Cause{}

		for _, e := range validationError.(validator.ValidationErrors) {
			cause := Cause{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return NewBadRequestValidationError("some field are invalids", errorsCauses)
	}

	return NewBadRequestError("error trying to convert fields")
}
