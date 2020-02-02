package validation

import (
	"github.com/xhminc/xhm-framework/config"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
)

func Username(
	v *validator.Validate,
	topStruct reflect.Value,
	currentStructOrField reflect.Value,
	field reflect.Value,
	fieldType reflect.Type,
	fieldKind reflect.Kind,
	param string,
) bool {
	if username, ok := field.Interface().(string); ok {
		return config.Mobile.MatchString(username) || config.Email.MatchString(username)
	}
	return false
}
