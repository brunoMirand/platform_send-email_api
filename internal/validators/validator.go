package validators

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(object interface{}) error {
	validate := validator.New()
	err := validate.Struct(object)
	if err == nil {
		return nil
	}

	validatorErrors := err.(validator.ValidationErrors)[0]
	field := strings.ToLower(validatorErrors.StructField())
	switch validatorErrors.Tag() {
	case "required":
		return errors.New(field + " is required")
	case "max":
		return errors.New(field + " is required with max " + validatorErrors.Param())
	case "min":
		return errors.New(field + " is required with min " + validatorErrors.Param())
	case "email":
		return errors.New(field + " is invalid")
	}
	return nil
}
