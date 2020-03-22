package validation

import (
	"github.com/xhminc/xhm-framework/config"
	"gopkg.in/go-playground/validator.v9"
)

func Email(fieldLevel validator.FieldLevel) bool {
	if email, ok := fieldLevel.Field().Interface().(string); ok {
		return config.Email.MatchString(email)
	}
	return false
}
