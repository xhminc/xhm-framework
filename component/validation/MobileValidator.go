package validation

import (
	"github.com/xhminc/xhm-framework/config"
	"gopkg.in/go-playground/validator.v9"
)

func Mobile(fieldLevel validator.FieldLevel) bool {
	if mobile, ok := fieldLevel.Field().Interface().(string); ok {
		return config.Mobile.MatchString(mobile)
	}
	return false
}
