package common

import (
	"github.com/xhminc/xhm-framework/component/result"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type RequestDTO struct {
}

func (r *RequestDTO) GetError(err validator.ValidationErrors) result.Result {

	var message string
	if len(err) > 0 {
		message = err[0].Param()
	} else {
		message = "request parameters error"
	}

	return result.Result{
		HttpStatus: http.StatusOK,
		Code:       -1,
		Message:    message,
	}
}
