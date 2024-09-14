package utils

import "imzakir.dev/e-commerce/app/domains/types"

func SetSuccessDeleteReponse(code int, message string) types.SuccessResponse {
	return types.SuccessResponse{
		StatusCode: code,
		Message:    message,
	}
}


func SetSuccessReponse(code int, message string, data any) types.SuccessResponse {
	return types.SuccessResponse{
		StatusCode: code,
		Message:    message,
		Data:       data,
	}
}

func SetErrorResponse(code int, message string, err error) types.ErrorResponse {
	return types.ErrorResponse{
		StatusCode: code,
		Message:    message,
		Error:      err.Error(),
	}
}
