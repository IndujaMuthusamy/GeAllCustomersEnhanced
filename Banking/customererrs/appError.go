package customererrs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func NotFoundError(message string) *AppError {
	return &AppError{Code: http.StatusNotFound, Message: message}
}

func InternalServerError(message string) *AppError {
	return &AppError{Code: http.StatusInternalServerError, Message: message}
}

func (e AppError) AsMessage() *AppError {
	return &AppError{Message: e.Message}
}

func BadRequest(message string) *AppError {
	return &AppError{Code: http.StatusBadRequest, Message: message}
}
