package config

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()

	_ = Validate.RegisterValidation("lettersonly", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		matched, _ := regexp.MatchString(`^[A-Za-záéíóúÑñ]+$`, value)

		return matched
	})
}
