package config

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func LoadValidator() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}
