package result

import "net/http"

type Result struct {
	HttpStatus int         `json:"-"`
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func Ok(data interface{}) Result {
	var result Result
	result.HttpStatus = http.StatusOK
	result.Code = 0
	result.Message = "ok"
	result.Data = data
	return result
}

func Error(data interface{}) Result {
	var result Result
	result.HttpStatus = http.StatusOK
	result.Code = -1
	result.Message = "error"
	result.Data = data
	return result
}
