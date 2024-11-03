package validator

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field   string
	Message string
}

func ParseError(err error) []ApiError {
	resp := make([]ApiError, 0, 2)
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, e := range ve {
			ae := ApiError{
				Field:   e.StructField(),
				Message: e.Tag(),
			}
			resp = append(resp, ae)
		}
	}

	return resp
}
