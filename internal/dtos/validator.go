package dtos

import (
	"fmt"

	"github.com/go-playground/validator"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func (v *Validator) ValidateStruct(s interface{}) error {
	return v.validator.Struct(s)
}

func (v *Validator) FormatValidationErrors(err error) map[string]string {
	formattedErrors := make(map[string]string)

	validationErrors, ok := err.(validator.ValidationErrors)

	if ok {
		for _, fieldErr := range validationErrors {
			formattedErrors[fieldErr.Field()] = fmt.Sprintf("Validation failed on '%s' rule", fieldErr.Tag())
		}
	}
	return formattedErrors
}
