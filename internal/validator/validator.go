package validator

import "github.com/go-playground/validator/v10"

type Validator interface {
	Validate(s any) error
}

type validatorImpl struct {
	validate *validator.Validate
}

func NewValidator() Validator {
	return &validatorImpl{
		validate: validator.New(),
	}
}

func (v *validatorImpl) Validate(s any) error {
	return v.validate.Struct(s)
}
