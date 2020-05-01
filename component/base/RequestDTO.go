package base

import (
	"github.com/go-playground/validator/v10"
	"net/http"
)

type RequestDTO struct {
	pageNo   int `json:"pageNo"`
	pageSize int `json:"pageSize"`
}

func (r *RequestDTO) GetPageNo() int {
	if r.pageNo <= 0 {
		return 1
	} else {
		return r.pageNo
	}
}

func (r *RequestDTO) GetPageSize() int {
	if r.pageSize <= 0 {
		return 50
	} else {
		return r.pageSize
	}
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
