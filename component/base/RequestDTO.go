package base

import (
	"github.com/go-playground/validator/v10"
	"net/http"
)

type RequestDTO struct {
	PageNo   int `json:"pageNo,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
}

func (r *RequestDTO) GetPageNo() int {
	if r.PageNo <= 0 {
		return 1
	} else {
		return r.PageNo
	}
}

func (r *RequestDTO) GetPageSize() int {
	if r.PageSize <= 0 {
		return 50
	} else {
		return r.PageSize
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
