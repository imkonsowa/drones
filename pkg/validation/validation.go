package validation

import (
	"github.com/go-playground/validator/v10"
)

// Validation is a wrapper for validator.Validate to provide the app dependencies
type Validation struct {
	Validator *validator.Validate
}

var v *Validation

func GetValidator() *Validation {
	if v == nil {
		v = newValidator()
	}

	return v
}

func newValidator() *Validation {
	validate := validator.New()

	return &Validation{
		Validator: validate,
	}
}
