package domain

import (
	"net/http"
)

type GenericResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type SuccessResponse struct {
	Data interface{}
}

func (s SuccessResponse) ToGenericResponse() GenericResponse {
	return GenericResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    s.Data,
	}
}
