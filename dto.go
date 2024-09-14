package dto_object

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Validate(dto interface{}) error {
	err := validate.Struct(dto)

	if err != nil {
		var validationErrors []string

		for _, err := range err.(validator.ValidationErrors) {
			tag := err.Tag()
			field := err.Field()

			if tag == "required" {
				validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", field, tag))
			} else if tag == "email" {
				validationErrors = append(validationErrors, fmt.Sprintf("%s got an invalid format", field))
			} else if tag == "oneof" {
				validationErrors = append(validationErrors, fmt.Sprintf("%s is not one of the list", field))
			} else if tag == "iscolor" {
				validationErrors = append(validationErrors, fmt.Sprintf("%s got an invalid color format", field))
			} else if tag == "lte" {
				validationErrors = append(validationErrors, fmt.Sprintf("%s must be less than or equal to %v", field, err.Param()))
			} else if tag == "gte" {
				validationErrors = append(validationErrors, fmt.Sprintf("%s must be greater than or equal to %v", field, err.Param()))
			} else {
				validationErrors = append(validationErrors, err.Error())
			}
		}

		return errors.New(strings.Join(validationErrors, "&"))
	}

	return nil
}

func Divide(err error) []string {
	errStr := err.Error()
	errs := strings.Split(errStr, "&")

	return errs
}
