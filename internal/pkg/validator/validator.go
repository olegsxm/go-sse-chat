package validator

import (
	v "github.com/go-playground/validator/v10"
)

type Validator struct {
	Validator *v.Validate
}

func (v *Validator) Validate(i any) error {
	if err := v.Validator.Struct(i); err != nil {
		return err
	}

	return nil
}

func New() *v.Validate {
	return v.New()
}
