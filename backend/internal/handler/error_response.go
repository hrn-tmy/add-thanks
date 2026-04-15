package handler

import "add-thanks/internal/gen"

func ErrorResponses(code, message string) gen.ErrorResponse {
	return gen.ErrorResponse{
		Errors: gen.ErrorResponseData{
			Code:    code,
			Message: message,
		},
	}
}
