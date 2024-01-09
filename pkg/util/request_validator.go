package util

import (
	"github.com/go-playground/validator"
)

type RequestValidator struct {
	validator *validator.Validate
}

func NewRequestValidation(val *validator.Validate) *RequestValidator {
	return &RequestValidator{
		validator: val,
	}
}

func (r *RequestValidator) Validate(i interface{}) error {
	if err := r.validator.Struct(i); err != nil {
		return err
	}
	return nil
}
