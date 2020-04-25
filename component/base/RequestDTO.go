package base

import (
	"github.com/go-playground/validator/v10"
	"net/http"
)

type RequestDTO struct {
}

func (r *RequestDTO) GetError(err validator.ValidationErrors) Result {

	var message string
	if len(err) > 0 {
		message = err[0].Param()
	} else {
		message = "request parameters error"
	}

	return Result{
		HttpStatus: http.StatusOK,
		Code:       -1,
		Message:    message,
	}
}
