package validator

import (
	"fmt"

	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	ru_translations "github.com/go-playground/validator/v10/translations/ru"
)

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	XValidator struct {
		validator  *validator.Validate
		translator ut.Translator
	}
)

func (v *XValidator) Struct(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := v.validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			err.Translate(v.translator)
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func New() *XValidator {
	r := ru.New()
	uni = ut.New(r, r)

	trans, _ := uni.GetTranslator("ru")

	validate = validator.New()

	if err := ru_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		fmt.Println("Initialize validator error ", err)
	}

	return &XValidator{validate, trans}
}
