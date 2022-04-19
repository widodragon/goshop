package validation

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func TitleValidation(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "Wido")
}
