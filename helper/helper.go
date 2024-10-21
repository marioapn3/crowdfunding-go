package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta `json:"meta"`
	// data nantinya akan menampung data yang akan kita kirimkan ke client. jadi bisa array bisa object
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ApiResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	res := Response{
		Meta: meta,
		Data: data,
	}

	return res
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
